package middlewares

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pedidosya/@project_name@/helpers"
	"github.com/pedidosya/@project_name@/models"
	"log"
	"net/http"
)

// headerResponseWriter gives the transaction access to response headers and the
// response code.
type headerResponseWriter struct{ w gin.ResponseWriter }

func (w *headerResponseWriter) Header() http.Header       { return w.w.Header() }
func (w *headerResponseWriter) Write([]byte) (int, error) { return 0, nil }
func (w *headerResponseWriter) WriteHeader(int)           {}

// replacementResponseWriter mimics the behavior of gin.ResponseWriter which
// buffers the response code rather than writing it when
// gin.ResponseWriter.WriteHeader is called.
type replacementResponseWriter struct {
	gin.ResponseWriter
	txn     newrelic.Transaction
	code    int
	written bool
}

func (w *replacementResponseWriter) flushHeader() {
	if !w.written {
		w.txn.WriteHeader(w.code)
		w.written = true
	}
}

func (w *replacementResponseWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *replacementResponseWriter) Write(data []byte) (int, error) {
	w.flushHeader()
	return w.ResponseWriter.Write(data)
}

func (w *replacementResponseWriter) WriteString(s string) (int, error) {
	w.flushHeader()
	return w.ResponseWriter.WriteString(s)
}

func (w *replacementResponseWriter) WriteHeaderNow() {
	w.flushHeader()
	w.ResponseWriter.WriteHeaderNow()
}

// setup new relic configuration application
func SetupNewRelic(c *models.Configuration) newrelic.Application {
	if c.NewRelic.AppName == "" && c.NewRelic.LicenseKey == "" {
		return nil
	}

	cfg := newrelic.NewConfig(c.NewRelic.AppName, c.NewRelic.LicenseKey)
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		log.Printf("failed to make new_relic app: %v", err)
	}

	return app
}

// NewRelicMonitoring is a middleware that starts a newrelic transaction
// stores it in the context
// then calls the next handler
func NewRelicMonitoring(app newrelic.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if app != nil {
			name := ctx.Request.URL.Path
			w := &headerResponseWriter{w: ctx.Writer}
			txn := app.StartTransaction(name, w, ctx.Request)
			defer txn.End()

			repl := &replacementResponseWriter{
				ResponseWriter: ctx.Writer,
				txn:            txn,
				code:           http.StatusOK,
			}
			ctx.Writer = repl
			defer repl.flushHeader()

			ctx.Set(helpers.NewRelicTxnKey, txn)
		}
		ctx.Next()
	}
}
