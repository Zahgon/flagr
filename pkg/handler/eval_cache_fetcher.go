package handler

import (
	"github.com/openflagr/flagr/pkg/entity"
	"gorm.io/gorm"
)

// EvalCacheJSON is the JSON serialization format of EvalCache's flags
type EvalCacheJSON struct {
	Flags []entity.Flag
}

func (ec *EvalCache) export() EvalCacheJSON { _ = "STUB: not implemented"; return *new(EvalCacheJSON) }

func (ec *EvalCache) fetchAllFlags() (idCache map[string]*entity.Flag, keyCache map[string]*entity.Flag, tagCache map[string]map[uint]*entity.Flag, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

type evalCacheFetcher interface {
	fetch() ([]entity.Flag, error)
}

func newFetcher() (evalCacheFetcher, error) {
	_ = "STUB: not implemented"
	return *new(evalCacheFetcher), nil
}

var fetchAllFlags = func() ([]entity.Flag, error) {
	fetcher, err := newFetcher()
	if err != nil {
		return nil, err
	}
	return fetcher.fetch()
}

type jsonFileFetcher struct {
	filePath string
}

func (ff *jsonFileFetcher) fetch() ([]entity.Flag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type jsonHTTPFetcher struct {
	url string
}

func (hf *jsonHTTPFetcher) fetch() ([]entity.Flag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type dbFetcher struct {
	db *gorm.DB
}

func (df *dbFetcher) fetch() ([]entity.Flag, error) {
	_ = "STUB: not implemented"
	// Use eager loading to avoid N+1 problem
	// doc: http://jinzhu.me/gorm/crud.html#preloading-eager-loading
	return nil, nil
}
