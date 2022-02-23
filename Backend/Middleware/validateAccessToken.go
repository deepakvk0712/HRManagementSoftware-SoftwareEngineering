package middleware

import (
	"context"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
	"strings"
)

func ValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		accessToken := req.Header.Get("Authorization")
		if accessToken == "" {
			errorResponses.SendUnauthorizedResponse(w, "No Authorization header")

			return
		}

		extractedToken := strings.Split(accessToken, "Bearer ")

		if len(extractedToken) == 2 {
			accessToken = strings.TrimSpace(extractedToken[1])
		} else {
			errorResponses.SendUnauthorizedResponse(w, "Incorrect format of authorization")

			return
		}

		claims, err := utils.ValidateToken(accessToken)
		if err != nil {
			errorResponses.SendUnauthorizedResponse(w, "Unauthorized, please login.")

			return
		}

		ctx := context.WithValue(req.Context(), "email", claims.Email)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
