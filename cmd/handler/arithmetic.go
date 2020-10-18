package handler

import (
	"log"
	"net/http"

	"github.com/realmallaury/teltech/internal/arithmetic"
	"github.com/realmallaury/teltech/internal/utils"

	"github.com/gin-gonic/gin"
)

// ArithmeticHandler holds data for handling basic math related requests.
type ArithmeticHandler struct {
	Logger *log.Logger
}

// Add resource accepts two numbers and returns result in JSON response.
func (ah *ArithmeticHandler) Add(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")

	if ok, err := utils.IsXYValid(x, y); !ok {
		ah.Logger.Printf("Add method validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := arithmetic.Add(x, y)
	if err != nil {
		ah.Logger.Printf("Add method error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Subtract resource accepts two numbers and returns result in JSON response.
func (ah *ArithmeticHandler) Subtract(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")

	if ok, err := utils.IsXYValid(x, y); !ok {
		ah.Logger.Printf("Add method validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := arithmetic.Subtract(x, y)
	if err != nil {
		ah.Logger.Printf("subtract method error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Multiply resource accepts two numbers and returns result in JSON response.
func (ah *ArithmeticHandler) Multiply(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")

	if ok, err := utils.IsXYValid(x, y); !ok {
		ah.Logger.Printf("Add method validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := arithmetic.Multiply(x, y)
	if err != nil {
		ah.Logger.Printf("multiply method error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Divide resource accepts two numbers and returns result in JSON response.
func (ah *ArithmeticHandler) Divide(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")

	if ok, err := utils.IsXYValid(x, y); !ok {
		ah.Logger.Printf("Add method validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := arithmetic.Divide(x, y)
	if err != nil {
		ah.Logger.Printf("divide method error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
