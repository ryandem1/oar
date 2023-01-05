// Package helpers context contains Gin context helpers
package helpers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

// CopyRequestBody will return a copy of the request body that can be used without altering the request body in the
// gin context. Helpful if you want to use the request body more than once
func CopyRequestBody(c *gin.Context) ([]byte, error) {
	byteBody, err := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))
	return byteBody, err
}
