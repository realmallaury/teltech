package handler

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/realmallaury/teltech/internal/cache"
)

// URL endpoint constants.
const (
	AddEndpoint      string = "/add"
	SubtractEndpoint string = "/subtract"
	MultiplyEndpoint string = "/multiply"
	DivideEndpoint   string = "/divide"
)

// Router initializes handler and middleware for API routes.
func Router(ctx context.Context, logger *log.Logger, store cache.Store) *gin.Engine {
	router := gin.New()

	// Middleware will write the logs to specified writer
	router.Use(gin.Logger())

	// Middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Custom middleware for caching result.
	middlewareHandler := Middleware{
		store: store,
	}

	arithmeticHandler := ArithmeticHandler{
		Logger: logger,
	}

	router.GET(AddEndpoint, middlewareHandler.CacheResult, arithmeticHandler.Add)
	router.GET(SubtractEndpoint, middlewareHandler.CacheResult, arithmeticHandler.Subtract)
	router.GET(MultiplyEndpoint, middlewareHandler.CacheResult, arithmeticHandler.Multiply)
	router.GET(DivideEndpoint, middlewareHandler.CacheResult, arithmeticHandler.Divide)

	return router
}
