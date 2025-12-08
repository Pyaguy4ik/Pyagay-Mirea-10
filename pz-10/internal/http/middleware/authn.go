package middleware

import (
    "context"
    "net/http"
    "strings"
    "pz10-auth/internal/platform/jwt"
    "pz10-auth/internal/core"
)

func AuthN(v jwt.Validator) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
                http.Error(w, `{"error":"authorization header required"}`, http.StatusUnauthorized)
                return
            }
            
            token := strings.TrimPrefix(authHeader, "Bearer ")
            claims, err := v.Parse(token)
            if err != nil {
                http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
                return
            }
            
            ctx := context.WithValue(r.Context(), core.ClaimsKey, map[string]any(claims))
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
