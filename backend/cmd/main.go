package main

import (
	"edu-web-backend/config"
	"edu-web-backend/internal/handlers"
	"edu-web-backend/internal/middleware"
	"edu-web-backend/internal/repository"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	db, err := repository.NewDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer db.Close()

	if err := db.Migrate(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Database migrated successfully")

	if err := db.MigrateAuth(); err != nil {
		log.Fatalf("Auth migration error: %v", err)
	}
	log.Println("Auth tables migrated successfully")

	if err := db.SeedData(); err != nil {
		log.Printf("Seed warning: %v", err)
	} else {
		log.Println("Data seeded successfully")
	}

	if err := db.SeedScenarios(); err != nil {
		log.Printf("Scenario seed warning: %v", err)
	} else {
		log.Println("Psychological scenarios seeded successfully")
	}

	h := handlers.NewHandler(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL, "http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api/v1")
	{
		api.GET("/health", h.HealthCheck)
		api.GET("/videos", h.GetVideos)
		api.GET("/audios", h.GetAudios)
		api.GET("/qrcodes", h.GetQRCodes)
		api.POST("/qrcodes/generate", h.GenerateQR)
		api.GET("/chat/:session_id", h.GetChatHistory)
		api.POST("/chat", h.SendChat)

		auth := api.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
			auth.GET("/me", middleware.AuthRequired(), h.GetMe)
		}

		protected := api.Group("")
		protected.Use(middleware.AuthRequired())
		{
			protected.POST("/messages", h.SendDirectMessage)
			protected.GET("/messages/:other_user_id", h.GetDirectMessages)
			protected.GET("/users", h.GetUsers)
		}
	}

	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
