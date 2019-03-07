package service

import (
	"net/http"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

func useLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("request-id")

		if reqID == "" {
			reqID = uuid.NewV4().String()
		}

		logger := log.With().Str("request-id", reqID).Logger()
		ctx := logger.WithContext(r.Context())

		log.Ctx(ctx).Debug().Msgf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
