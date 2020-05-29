package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	movie "github.com/madgam/sukinema/pkg"
)

func main() {
	e := echo.New()
	e.GET("/api/v1/movies", getAllMovie)
	e.GET("/api/v1/movies/pref/:prefID", getSelectedPrefMovie)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func getAllMovie(c echo.Context) error {

	ctx := context.Background()
	movies, err := movie.GetAllMovie(ctx)
	if err != nil {
		log.Fatalf("Failed to get all movie: %v", err)
	}

	return c.JSON(http.StatusOK, movies)
}

func getSelectedPrefMovie(c echo.Context) error {

	ctx := context.Background()
	prefID := c.Param("prefID")
	movies, err := movie.GetPrefMovie(ctx, prefID)
	if err != nil {
		log.Fatalf("Failed to get pref movie: %v", err)
	}

	return c.JSON(http.StatusOK, movies)
}
