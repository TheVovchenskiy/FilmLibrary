package middleware

import (
	"context"
	"encoding/json"
	"filmLibrary/pkg/logging"
	"filmLibrary/pkg/responseTemplate"
	"filmLibrary/pkg/utils"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	status        int
	body          string
	isHeaderWrote bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.isHeaderWrote {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.isHeaderWrote = true
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if rw.isHeaderWrote && rw.status >= 400 {
		message := responseTemplate.ErrToSend{}
		err := json.Unmarshal(data, &message)
		if err == nil {
			rw.body = message.Message
		}
	}

	return rw.ResponseWriter.Write(data)
}

func newResponseWriterWrapper(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func AccessLogMiddleware( /* logger *logrus.Logger,  */ next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := newResponseWriterWrapper(w)

		// requestID := r.Context().Value(REQUEST_ID_KEY)
		requestID := utils.GetRequestIDFromCtx(r.Context())
		logging.Logger = logging.Logger.WithField("request_id", requestID).Logger

		// next.ServeHTTP(wrappedWriter, r)

		contextLogger := logging.Logger.WithFields(logrus.Fields{
			"method":     r.Method,
			"URL":        r.URL.Path,
			"request_id": requestID,
		})

		ctx := context.WithValue(r.Context(), utils.LOGGER_KEY, contextLogger)

		next.ServeHTTP(wrappedWriter, r.WithContext(ctx))

		toLog := contextLogger.WithFields(logrus.Fields{
			"status":         wrappedWriter.status,
			"execution_time": time.Since(start).String(),
		})

		if wrappedWriter.body != "" {
			toLog.Error(wrappedWriter.body)
		} else {
			toLog.Info("HTTP Request")
		}
	})
}
