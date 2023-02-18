package models

import "github.com/gin-gonic/gin"

// ConvertErrToGinH will convert any go error into a gin.H response body to return back to the caller
func ConvertErrToGinH(err error) gin.H {
	return gin.H{"error": err.Error()}
}
