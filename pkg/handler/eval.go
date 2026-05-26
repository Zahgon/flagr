package handler

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"encoding/json"

	"github.com/openflagr/flagr/pkg/config"
	"github.com/openflagr/flagr/pkg/entity"
	"github.com/openflagr/flagr/pkg/util"
	"github.com/openflagr/flagr/swagger_gen/models"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/evaluation"
	"gorm.io/gorm"

	"github.com/bsm/ratelimit"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime/middleware"
	"github.com/zhouzhuojie/conditions"
)

// Eval is the Eval interface
type Eval interface {
	PostEvaluation(evaluation.PostEvaluationParams) middleware.Responder
	PostEvaluationBatch(evaluation.PostEvaluationBatchParams) middleware.Responder
}

// NewEval creates a new Eval instance
func NewEval() Eval { _ = "STUB: not implemented"; return *new(Eval) }

type eval struct{}

func (e *eval) PostEvaluation(params evaluation.PostEvaluationParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (e *eval) PostEvaluationBatch(params evaluation.PostEvaluationBatchParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

// Deduplicate flagKeys to prevent DoS via repeated keys

// Deduplicate flagIDs to prevent DoS via repeated IDs

// Validate batch size to prevent DoS attacks via resource exhaustion (if enabled)

// Calculate total evaluations: entities * (flagIDs + flagKeys + flagTags)
// For flagTags, we count each tag as potentially matching one flag (conservative estimate)

// flagTags is evaluated once per entity regardless of count

// TODO make it concurrent

// BlankResult creates a blank result
func BlankResult(f *entity.Flag, evalContext models.EvalContext, msg string) *models.EvalResult {
	_ = "STUB: not implemented"
	return nil
}

var LookupFlag = func(evalContext models.EvalContext) *entity.Flag {
	cache := GetEvalCache()
	flagID := util.SafeUint(evalContext.FlagID)
	flagKey := util.SafeString(evalContext.FlagKey)
	f := cache.GetByFlagKeyOrID(flagID)
	if f == nil {
		f = cache.GetByFlagKeyOrID(flagKey)
	}
	return f
}

var EvalFlagsByTags = func(evalContext models.EvalContext) []*models.EvalResult {
	cache := GetEvalCache()
	fs := cache.GetByTags(evalContext.FlagTags, evalContext.FlagTagsOperator)
	results := []*models.EvalResult{}
	for _, f := range fs {
		results = append(results, EvalFlagWithContext(f, evalContext))
	}
	return results
}

var EvalFlag = func(evalContext models.EvalContext) *models.EvalResult {
	flag := LookupFlag(evalContext)
	return EvalFlagWithContext(flag, evalContext)
}

var EvalFlagWithContext = func(flag *entity.Flag, evalContext models.EvalContext) *models.EvalResult {
	flagID := util.SafeUint(evalContext.FlagID)
	flagKey := util.SafeString(evalContext.FlagKey)

	if flag == nil {
		emptyFlag := &entity.Flag{Model: gorm.Model{ID: flagID}, Key: flagKey}
		return BlankResult(emptyFlag, evalContext, fmt.Sprintf("flagID %v not found or deleted", flagID))
	}

	if !flag.Enabled {
		return BlankResult(flag, evalContext, fmt.Sprintf("flagID %v is not enabled", flag.ID))
	}

	if len(flag.Segments) == 0 {
		return BlankResult(flag, evalContext, fmt.Sprintf("flagID %v has no segments", flag.ID))
	}

	if evalContext.EntityID == "" {
		evalContext.EntityID = fmt.Sprintf("randomly_generated_%d", rand.Int31())
	}

	if flag.EntityType != "" {
		evalContext.EntityType = flag.EntityType
	}

	logs := []*models.SegmentDebugLog{}
	var vID int64
	var sID int64

	for _, segment := range flag.Segments {
		sID = int64(segment.ID)
		variantID, log, evalNextSegment := evalSegment(flag.ID, evalContext, segment)
		if config.Config.EvalDebugEnabled && evalContext.EnableDebug {
			logs = append(logs, log)
		}
		if variantID != nil {
			vID = int64(*variantID)
		}
		if !evalNextSegment {
			break
		}
	}
	evalResult := BlankResult(flag, evalContext, "")
	evalResult.EvalDebugLog.SegmentDebugLogs = logs
	evalResult.SegmentID = sID
	evalResult.VariantID = vID
	v := flag.FlagEvaluation.VariantsMap[util.SafeUint(vID)]
	if v != nil {
		evalResult.VariantAttachment = v.Attachment
		evalResult.VariantKey = v.Key
	}

	logEvalResult(evalResult, flag.DataRecordsEnabled)
	return evalResult
}

var logEvalResult = func(r *models.EvalResult, dataRecordsEnabled bool) {
	if r == nil {
		// this is just a safety check, r is from BlankResult,
		// and usually it cannot be nil
		return
	}

	if config.Config.EvalLoggingEnabled {
		rateLimitPerFlagConsoleLogging(r)
	}

	logEvalResultToDatadog(r)
	logEvalResultToPrometheus(r)

	if !config.Config.RecorderEnabled || !dataRecordsEnabled {
		return
	}
	rec := GetDataRecorder()
	rec.AsyncRecord(*r)
}

var logEvalResultToDatadog = func(r *models.EvalResult) {
	if config.Global.StatsdClient == nil {
		return
	}

	config.Global.StatsdClient.Incr(
		"evaluation",
		[]string{
			fmt.Sprintf("FlagID:%d", util.SafeUint(r.FlagID)),
			fmt.Sprintf("VariantID:%d", util.SafeUint(r.VariantID)),
			fmt.Sprintf("VariantKey:%s", util.SafeStringWithDefault(r.VariantKey, "null")),
		},
		float64(1),
	)
}

var logEvalResultToPrometheus = func(r *models.EvalResult) {
	if config.Global.Prometheus.EvalCounter == nil {
		return
	}
	config.Global.Prometheus.EvalCounter.WithLabelValues(
		util.SafeStringWithDefault(r.EvalContext.EntityType, "null"),
		util.SafeStringWithDefault(r.FlagID, "null"),
		util.SafeStringWithDefault(r.FlagKey, "null"),
		util.SafeStringWithDefault(r.VariantID, "null"),
		util.SafeStringWithDefault(r.VariantKey, "null"),
	).Inc()

}

var evalSegment = func(
	flagID uint,
	evalContext models.EvalContext,
	segment entity.Segment,
) (
	vID *uint, // returns VariantID
	log *models.SegmentDebugLog,
	evalNextSegment bool,
) {
	if len(segment.Constraints) != 0 {
		m, ok := evalContext.EntityContext.(map[string]any)
		if !ok {
			log = &models.SegmentDebugLog{
				Msg:       fmt.Sprintf("constraints are present in the segment_id %v, but got invalid entity_context: %s.", segment.ID, spew.Sdump(evalContext.EntityContext)),
				SegmentID: int64(segment.ID),
			}
			return nil, log, true
		}

		expr := segment.SegmentEvaluation.ConditionsExpr
		match, err := conditions.Evaluate(expr, m)
		if err != nil {
			if config.Config.EvalDebugEnabled && evalContext.EnableDebug {
				log = &models.SegmentDebugLog{
					Msg:       err.Error(),
					SegmentID: int64(segment.ID),
				}
			}
			return nil, log, true
		}
		if !match {
			if config.Config.EvalDebugEnabled && evalContext.EnableDebug {
				log = &models.SegmentDebugLog{
					Msg:       debugConstraintMsg(evalContext.EnableDebug, expr, m),
					SegmentID: int64(segment.ID),
				}
			}
			return nil, log, true
		}
	}

	vID, debugMsg := segment.SegmentEvaluation.DistributionArray.Rollout(
		evalContext.EntityID,
		segment.SegmentEvaluation.FlagIDStr, // pre-formatted during PrepareEvaluation
		segment.RolloutPercent,
	)

	log = &models.SegmentDebugLog{
		Msg:       "matched all constraints. " + debugMsg,
		SegmentID: int64(segment.ID),
	}

	// at this point, all constraints are matched, so we shouldn't go to next segment
	// thus setting evalNextSegment = false
	return vID, log, false
}

func debugConstraintMsg(enableDebug bool, expr conditions.Expr, m map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

var rateLimitMap = sync.Map{}

var rateLimitPerFlagConsoleLogging = func(r *models.EvalResult) {
	flagID := util.SafeUint(r.FlagID)
	rl, _ := rateLimitMap.LoadOrStore(flagID, ratelimit.New(
		config.Config.RateLimiterPerFlagPerSecondConsoleLogging,
		time.Second,
	))
	if !rl.(*ratelimit.RateLimiter).Limit() {
		jsonStr, _ := json.Marshal(struct{ FlagEvalResult *models.EvalResult }{FlagEvalResult: r})
		fmt.Println(string(jsonStr))
	}
}
