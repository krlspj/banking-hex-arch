package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/service"
	"github.com/krlspj/banking-hex-arch/07_security/internal/errs"
)

type AuthMiddleware struct {
	service service.AuthService
}

func (a AuthMiddleware) authroizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouterVars := mux.Vars(r)
			token := r.Header.Get("X-Access-Token")

			if token != "" {
				//token := a.getTokenFromHeader(authHeader)

				isAuthorized := a.service.IsAuthorized(token, currentRoute.GetName(), currentRouterVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}

}
