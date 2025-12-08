package main

import (
    "log"
    "net/http"
    "pz10-auth/internal/platform/config"
    "pz10-auth/internal/http/router"
)

func main() {
    cfg := config.Load()
    mux := router.Build(cfg)
    
    log.Printf("Server starting on %s", cfg.Port)
    log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
