package main

import (
	"os"
	"patient-chatbot/internal/client/llm"
	"patient-chatbot/internal/client/stock"
	"patient-chatbot/internal/config"
	"patient-chatbot/internal/handler"
	logger "patient-chatbot/internal/log"
	"patient-chatbot/internal/middleware"
	"patient-chatbot/internal/service"
	"patient-chatbot/internal/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.LocaleMiddleware(utils.Bundle))
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestID())

	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.Use(logger.Init())

	stockClient := stock.NewStockClient(cfg)
	llmClient := llm.NewLLMClient(cfg, stockClient)
	chatService := service.NewService(cfg, llmClient, stockClient)
	h := handler.NewHandler(chatService)

	RegisterRoutes(r, h)

	return &Server{router: r}
}

func (s *Server) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return s.router.Run(":" + port)
}

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1")
	{
		api.GET("/health", h.HandleGetHealth)
		api.POST("/chat", h.HandleChat)
		api.GET("/dashboard", h.HandleGetDashboard)
		api.GET("/dashboard/chart", h.HandleGetCompanyChart)
	}
}
