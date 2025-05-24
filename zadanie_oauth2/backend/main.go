package main

import (
	"backend/db"
	"backend/routes"
	"net/http"
	"github.com/joho/godotenv"
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Nie udało się załadować pliku .env")
	}



	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	db.ConnectDB()
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
