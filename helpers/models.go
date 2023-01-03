package helpers

import "github.com/gin-gonic/gin"

func ConvertErrToGinH(err error) gin.H {
	return gin.H{"error": err.Error()}
}
