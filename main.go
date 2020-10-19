package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	myMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	//INIT ECHO
	e := echo.New()

	//Custom Middleware
	e.Use(CustomMiddleWareForServerHeader)

	adminGroup := e.Group("/admin")

	//ROUTING
	e.GET("/", welcome)
}
