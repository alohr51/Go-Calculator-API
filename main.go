package main

import (
	"github.com/calculator/api/calc"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	calc.Routes(router)
	router.Run()
}
