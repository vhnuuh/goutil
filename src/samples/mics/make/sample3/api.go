package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"os"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", helloword.HelloWorld)
	e.start(os.Getenv("ADDR"))
}
