// Package controllers contains Gin context helpers
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"io"
	"strings"
)

// CopyRequestBody will return a copy of the request body that can be used without altering the request body in the
// gin context. Helpful if you want to use the request body more than once
func CopyRequestBody(c *gin.Context) ([]byte, error) {
	byteBody, err := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))
	return byteBody, err
}

// DoubleBindTest will perform a double-bind of the request body to: first deserialize the structure attributes of a
// test, then the second bind will deserialize any dynamic attributes into a models.Test's Doc. Will return a pointer
// to the fully-bound models.Test object and any potential errors that occurred during the process. Note that it will
// not be validated or cleaned.
func DoubleBindTest(c *gin.Context) (*Test, error) {
	test := &Test{}

	// We must copy our request body for the second unmarshal because the bind operation will consume it
	byteBody, err := CopyRequestBody(c)
	if err != nil {
		return nil, err
	}

	// First bind binds the test information
	if err := c.BindJSON(test); err != nil {
		return nil, err
	}

	// Check if Doc is manually defined, it should not be. If it is, it causes all sorts of conflicts
	if test.Doc != nil {
		return nil, fmt.Errorf("'Doc' is reserved! Cannot use that key")
	}

	// Second bind will move all dynamic fields to test.Doc
	if err := json.Unmarshal(byteBody, &test.Doc); err != nil {
		return nil, err
	}

	// Removes the keys that are from the first binding
	for key := range test.Doc {
		if slices.Contains([]string{"summary", "id", "outcome", "analysis", "resolution"}, strings.ToLower(key)) {
			delete(test.Doc, key)
		}
	}

	return test, nil
}
