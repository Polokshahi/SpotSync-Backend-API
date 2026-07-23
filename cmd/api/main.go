package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"spotsync/internal/database"
	"spotsync/internal/handler"
	"spotsync/internal/models"
	"spotsync/internal/repository"
	"spotsync/internal/routes"
	"spotsync/internal/service"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found")
	}

	// Database
	db := database.ConnectDB()

	// Check Database Connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get database instance:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("❌ Database Ping Failed:", err)
	}

	log.Println("✅ Neon Database Connected Successfully")

	// ----------------------------
	// Auto Migration
	// ----------------------------
	err = db.AutoMigrate(
		&models.User{},
		&models.ParkingZone{},
		&models.Reservation{},
	)

	if err != nil {
		log.Fatal("❌ Auto migration failed:", err)
	}

	log.Println("✅ Database migrated successfully")

	// Echo
	e := echo.New()

	// Repository
	userRepo := repository.NewUserRepository(db)
	zoneRepo := repository.NewZoneRepository(db)
	reservationRepo := repository.NewReservationRepository(db)

	// Service
	authService := service.NewAuthService(userRepo)
	zoneService := service.NewZoneService(zoneRepo)
	reservationService := service.NewReservationService(
		reservationRepo,
		db,
	)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	zoneHandler := handler.NewZoneHandler(zoneService)
	reservationHandler := handler.NewReservationHandler(reservationService)

	// Register Routes (Only Once)
	routes.RegisterRoutes(
		e,
		authHandler,
		zoneHandler,
		reservationHandler,
	)

	// Print Registered Routes
	for _, r := range e.Routes() {
		log.Printf("%-6s %s\n", r.Method, r.Path)
	}

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port:", port)

	e.Logger.Fatal(e.Start(":" + port))
}