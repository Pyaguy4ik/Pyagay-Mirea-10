package router

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "pz10-auth/internal/http/middleware"
    "pz10-auth/internal/platform/config"
    "pz10-auth/internal/platform/jwt"
    "pz10-auth/internal/repo"
    "pz10-auth/internal/core"
)

func Build(cfg config.Config) http.Handler {
    r := chi.NewRouter()
    
    // Dependency Injection
    userRepo := repo.NewUserMem()
    jwtv := jwt.NewHS256(cfg.JWTSecret, cfg.JWTTTL)
    svc := core.NewService(userRepo, jwtv)
    
    // Public routes
    r.Post("/api/v1/login", svc.LoginHandler)
    
    // Protected routes for users and admins
    r.Group(func(priv chi.Router) {
        priv.Use(middleware.AuthN(jwtv))
        priv.Use(middleware.AuthZRoles("admin", "user"))
        priv.Get("/api/v1/me", svc.MeHandler)
    })
    
    // Admin-only routes
    r.Group(func(admin chi.Router) {
        admin.Use(middleware.AuthN(jwtv))
        admin.Use(middleware.AuthZRoles("admin"))
        admin.Get("/api/v1/admin/stats", svc.AdminStats)
    })
    
    return r
}
