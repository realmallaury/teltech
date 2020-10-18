package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/realmallaury/teltech/internal/arithmetic"
	"github.com/stretchr/testify/assert"
)

func getTestResources() (*gin.Engine, ArithmeticHandler) {
	arithmeticHandler := ArithmeticHandler{
		Logger: log.New(os.Stdout, "Test : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
	}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r, arithmeticHandler
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	r, arithmeticHandler := getTestResources()

	// Route for /add endpoint
	r.GET(AddEndpoint, arithmeticHandler.Add)

	// Test that GET to /add returns addition of two numbers
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, createQueryURL(AddEndpoint, "1", "1"), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code, "Response status should be OK")

	// Unmarshal response body
	var result arithmetic.Result
	_ = json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal("2", result.Answer, "Result should be the same")

	// Test that bad request GET to /add returns error message
	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, createQueryURL(AddEndpoint, "1--", "1.."), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusBadRequest, w.Code, "Response status should be Bad Request")
	assert.Equal(
		`{"error":"x value: 1-- not valid number, y value: 1.. not valid number"}`,
		w.Body.String(),
		"Response should contain error message",
	)
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)
	r, arithmeticHandler := getTestResources()

	// Route for /subtract endpoint
	r.GET(SubtractEndpoint, arithmeticHandler.Subtract)

	// Test that GET to /subtract returns addition of two numbers
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, createQueryURL(SubtractEndpoint, "1", "1"), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code, "Response status should be OK")

	// Unmarshal response body
	var result arithmetic.Result
	_ = json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal("0", result.Answer, "Result should be the same")

	// Test that bad request GET to /subtract returns error message
	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, createQueryURL(SubtractEndpoint, "1--", "1.."), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusBadRequest, w.Code, "Response status should be Bad Request")
	assert.Equal(
		`{"error":"x value: 1-- not valid number, y value: 1.. not valid number"}`,
		w.Body.String(),
		"Response should contain error message",
	)
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)
	r, arithmeticHandler := getTestResources()

	// Route for /multiply endpoint
	r.GET(MultiplyEndpoint, arithmeticHandler.Multiply)

	// Test that GET to /multiply returns addition of two numbers
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, createQueryURL(MultiplyEndpoint, "2", "2"), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code, "Response status should be OK")

	// Unmarshal response body
	var result arithmetic.Result
	_ = json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal("4", result.Answer, "Result should be the same")

	// Test that bad request GET to /multiply returns error message
	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, createQueryURL(MultiplyEndpoint, "1--", "1.."), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusBadRequest, w.Code, "Response status should be Bad Request")
	assert.Equal(
		`{"error":"x value: 1-- not valid number, y value: 1.. not valid number"}`,
		w.Body.String(),
		"Response should contain error message",
	)
}

func TestDivide(t *testing.T) {
	assert := assert.New(t)
	r, arithmeticHandler := getTestResources()

	// Route for /divide endpoint
	r.GET(DivideEndpoint, arithmeticHandler.Divide)

	// Test that GET to /divide returns addition of two numbers
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, createQueryURL(DivideEndpoint, "2", "2"), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code, "Response status should be OK")

	// Unmarshal response body
	var result arithmetic.Result
	_ = json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal("1", result.Answer, "Result should be the same")

	// Test that bad request GET to /divide returns error message
	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, createQueryURL(DivideEndpoint, "1--", "1.."), nil)

	assert.NoError(err, "Error should be nil")

	r.ServeHTTP(w, req)
	assert.Equal(http.StatusBadRequest, w.Code, "Response status should be Bad Request")
	assert.Equal(
		`{"error":"x value: 1-- not valid number, y value: 1.. not valid number"}`,
		w.Body.String(),
		"Response should contain error message",
	)
}

func createQueryURL(endpoint, x, y string) string {
	params := url.Values{}
	params.Add("x", x)
	params.Add("y", y)

	return fmt.Sprintf("%s?%s", endpoint, params.Encode())
}
