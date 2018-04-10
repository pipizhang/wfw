package main

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	ErrorHttpCommonMessage = echo.NewHTTPError(http.StatusInternalServerError, "somthing went wrong, please try again later")
	ErrorWrongCredentials  = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invaid")
	JWTSecret              = "mysecret"
	Version                = "1.0"
)

type (
	LoginRequest struct {
		Username string `json:"uerrname"`
		Password string `json:"password"`
	}
)

func LoginHandler(c echo.Context) error {
	reqData := LoginRequest{}
	err := json.Decoder(c.Request().Body).Decode(&reqData)
	if err != nil {
		log.Printf("invalid request, could not read credentials")
		return ErrorHttpCommonMessage
	}

	ctx := c.Request().Context()

	return c.String(http.StatusOK, "login")
}

func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func VersionHandler(c echo.Context) error {
	return c.String(http.StatusOK, Version)
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/vesion", VersionHandler)
	e.GET("/ping", PingHandler)
	e.POST("/login", LoginHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
