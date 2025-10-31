package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// init functions are executed before main, and is mostly used for package initialization like setting up loggers, initializing configurations, DB connections, etc.

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {

	ctx := context.Background()

	logrus.WithContext(ctx).Info("Starting API Gateway...")

	e := echo.New()

	e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3001",
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
		},
	}))

	e.Use(echomw.RateLimiter(
		echomw.NewRateLimiterMemoryStore(5),
	))

	err := godotenv.Load()

	if err != nil {
		logrus.WithContext(ctx).Fatalf("Error loading .env file: %v", err)
	}

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy")
	})

	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf(":%s", port)
	// fmt.Println("Server is running on port:", port)
	logrus.WithContext(ctx).Infof("API Gateway started on %s", appAddress)
	e.Logger.Fatal(e.Start(appAddress))

}
