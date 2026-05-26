package config

import (
	"net/http"

	"github.com/DataDog/datadog-go/statsd"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
)

// ServerShutdown is a callback function that will be called when
// we tear down the flagr server
func ServerShutdown() { _ = "STUB: not implemented"; return }

// SetupGlobalMiddleware setup the global middleware
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type recoveryLogger struct{}

func (r *recoveryLogger) Printf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (r *recoveryLogger) Println(v ...any) { _ = "STUB: not implemented"; return }

func setupRecoveryMiddleware() *negroni.Recovery { _ = "STUB: not implemented"; return nil }

/*
setupJWTAuthMiddleware setup an JWTMiddleware from the ENV config
*/
func setupJWTAuthMiddleware() *jwtAuth { _ = "STUB: not implemented"; return nil }

func jwtErrorHandler(w http.ResponseWriter, r *http.Request, err string) {
	_ = "STUB: not implemented"
	return
}

type jwtAuth struct {
	PrefixWhitelistPaths []string
	ExactWhitelistPaths  []string
	JWTMiddleware        *jwtmiddleware.JWTMiddleware
}

func (a *jwtAuth) whitelist(req *http.Request) bool { _ = "STUB: not implemented"; return false }

// If we set to 401 unauthorized, let the client handles the 401 itself

func (a *jwtAuth) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

/*
setupBasicAuthMiddleware setup an BasicMiddleware from the ENV config
*/
func setupBasicAuthMiddleware() *basicAuth { _ = "STUB: not implemented"; return nil }

type basicAuth struct {
	Username             []byte
	Password             []byte
	PrefixWhitelistPaths []string
	ExactWhitelistPaths  []string
}

func (a *basicAuth) whitelist(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func (a *basicAuth) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

type statsdMiddleware struct {
	StatsdClient *statsd.Client
}

func (s *statsdMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

type prometheusMiddleware struct {
	counter   *prometheus.CounterVec
	latencies *prometheus.HistogramVec
}

func (p *prometheusMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}
