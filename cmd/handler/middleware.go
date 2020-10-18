package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realmallaury/teltech/internal/arithmetic"
	"github.com/realmallaury/teltech/internal/cache"
)

// Middleware handles caching results.
type Middleware struct {
	store cache.Store
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// CacheResult gets result from cache or stores new result if not present.
func (m *Middleware) CacheResult(c *gin.Context) {
	w := &bodyWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
	c.Writer = w

	value, ok := m.store.GetRecord(c.Request.URL.String())
	if ok {
		result := value.(arithmetic.Result)
		result.Cached = true
		c.AbortWithStatusJSON(http.StatusOK, result)
		return
	}

	c.Next()

	if c.Writer.Status() == http.StatusOK {
		var result arithmetic.Result

		if err := json.Unmarshal(w.body.Bytes(), &result); err == nil {
			m.store.StoreRecord(c.Request.URL.String(), result)
			return
		}
	}
}
