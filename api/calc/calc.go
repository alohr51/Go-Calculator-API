package calc

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	firstOperandKey  = "firstOperand"
	secondOperandKey = "secondOperand"
	sum              = "sum"
	subtract         = "subtract"
	multiply         = "multiply"
	divide           = "divide"
)

// Routes - Takes in the gin router and adds routes that perform operations on 2 numbers
func Routes(router *gin.Engine) {
	calcInts := router.Group("/api/calc")
	{
		calcInts.GET("/sum/:firstOperand/:secondOperand", handleOp(sum))
		calcInts.GET("/subtract/:firstOperand/:secondOperand", handleOp(subtract))
		calcInts.GET("/multiply/:firstOperand/:secondOperand", handleOp(multiply))
		calcInts.GET("/divide/:firstOperand/:secondOperand", handleOp(divide))
	}
}

func handleOp(op string) gin.HandlerFunc {

	return func(c *gin.Context) {
		firstOperand, err := strconv.ParseFloat(c.Param(firstOperandKey), 64)

		if err != nil {
			setError(c, err.Error())
			return
		}

		secondOperand, err := strconv.ParseFloat(c.Param(secondOperandKey), 64)

		if err != nil {
			setError(c, err.Error())
			return
		}

		if op == "divide" && secondOperand == 0 {
			setError(c, "Cannot Divide by 0")
			return
		}

		c.JSON(200, gin.H{
			"result": getResult(firstOperand, secondOperand, op),
		})
	}
}

func getResult(first, second float64, op string) float64 {
	switch op {
	case sum:
		return first + second
	case subtract:
		return first - second
	case multiply:
		return first * second
	case divide:
		return first / second
	default:
		log.Println("Unsupported operator", op)
		return 0
	}
}

func setError(c *gin.Context, err string) {
	c.JSON(400, gin.H{
		"msg": err,
	})
}
