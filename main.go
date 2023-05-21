package main

import (
	"feather/credmanage"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()
	e.POST("/signup",credmanage.SignUp)
	e.POST("/login",credmanage.Authenticate)
	e.Logger.Fatal(e.Start(":5000"))
}

