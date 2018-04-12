package main

import (
	"encoding/json"
	_ "fmt"
	"net/http"
	"os"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	ErrorHttpCommonMessage = echo.NewHTTPError(http.StatusInternalServerError, "somthing went wrong, please try again later")
	ErrorWrongCredentials  = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invaid")
	App                    = NewApp()
)

type (
	Application struct {
		Port           string
		JWTSecret      string
		Version        string
		UserAPIAddress string
		UserService    UserService
	}

	LoginRequest struct {
		Username string `json:"uerrname"`
		Password string `json:"password"`
	}
)

func NewApp() (app *Application) {
	app = &Application{
		Port:      ":8081",
		JWTSecret: "mysecret",
	}

	godotenv.Load()

	if len(os.Getenv("AUTH_API_PORT")) != 0 {
		app.Port = ":" + os.Getenv("AUTH_API_PORT")
	}
	app.UserAPIAddress = os.Getenv("USER_API_ADDRESS")
	if len(os.Getenv("JWT_SECRET")) != 0 {
		app.JWTSecret = os.Getenv("JWT_SECRET")
	}

	return
}

func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func LoginHandler(c echo.Context) error {
	reqData := LoginRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&reqData)
	if err != nil {
		log.Printf("invalid request, could not read credentials")
		return ErrorHttpCommonMessage
	}

	//ctx := c.Request().Context()

	return c.String(http.StatusOK, "login")
}

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Meddlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/ping", PingHandler)
	e.POST("/login", LoginHandler)

	e.Logger.Fatal(e.Start(App.Port))
}
