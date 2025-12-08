package middleware

import (
    "net/http"
    "pz10-auth/internal/core"
)

func AuthZRoles(allowed ...string) func(http.Handler) http.Handler {
    set := make(map[string]struct{})
    for _, a := range allowed {
        set[a] = struct{}{}
    }
    
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            claims, ok := r.Context().Value(core.ClaimsKey).(map[string]any)
            if !ok {
                http.Error(w, `{"error":"claims not found"}`, http.StatusForbidden)
                return
            }
            
            role, ok := claims["role"].(string)
            if !ok {
                http.Error(w, `{"error":"role not found"}`, http.StatusForbidden)
                return
            }
            
            if _, ok := set[role]; !ok {
                http.Error(w, `{"error":"insufficient permissions"}`, http.StatusForbidden)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
