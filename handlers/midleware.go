package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spinales/ASIA-backend/util"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

var PasetoKey = "pasetoStruct"

// HTTP middleware setting a value on the request context
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Values(authorizationHeaderKey)
		log.Println(authorizationHeader)

		if len(authorizationHeader) == 0 {
			util.RespondWithError(w, http.StatusUnauthorized, "authorization token is not provided.")
			return
		}

		fields := strings.Fields(authorizationHeader[0])
		if len(fields) < 2 {
			util.RespondWithError(w, http.StatusUnauthorized, "invalid authorization header format")
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			util.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("unsupported authorization type %s", authorizationType))
			return
		}

		accessToken := fields[1]
		payload, err := s.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			util.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), PasetoKey, payload)

		w.Header().Set(authorizationPayloadKey, payload.Expiration.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
