package routes

import (
	"server-service/auth"
	"server-service/middleware"
	"server-service/config"
)

var (
	DB = config.ConnectDB()
	authService = auth.NewAuthService()
	MainMiddleware = middleware.Middleware(authService)
)
