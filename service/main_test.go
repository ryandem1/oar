package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	if config.PG == nil {
		t.Error("PG config was null")
	}
}

func TestGetRouter(t *testing.T) {
	router := GetRouter()
	routes := router.Routes()

	expectedRoutes := []gin.RouteInfo{
		{
			Method:      http.MethodGet,
			Path:        "/tests",
			Handler:     "github.com/ryandem1/oar.(*TestController).GetTests-fm",
			HandlerFunc: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/health",
			Handler:     "github.com/ryandem1/oar.GetRouter.func2",
			HandlerFunc: nil,
		},
		{
			Method:      http.MethodPatch,
			Path:        "/tests",
			Handler:     "github.com/ryandem1/oar.(*TestController).PatchTests-fm",
			HandlerFunc: nil,
		},
		{
			Method:      http.MethodDelete,
			Path:        "/tests",
			Handler:     "github.com/ryandem1/oar.(*TestController).DeleteTests-fm",
			HandlerFunc: nil,
		},
		{
			Method:      http.MethodPost,
			Path:        "/test",
			Handler:     "github.com/ryandem1/oar.(*TestController).CreateTest-fm",
			HandlerFunc: nil,
		},
		{
			Method:      http.MethodPost,
			Path:        "/query",
			Handler:     "github.com/ryandem1/oar.EncodeSearchQuery",
			HandlerFunc: nil,
		},
	}
	for i, expectedRoute := range expectedRoutes {
		assert.Equal(t, routes[i].Method, expectedRoute.Method)
		assert.Equal(t, routes[i].Path, expectedRoute.Path)
		assert.Equal(t, routes[i].Handler, expectedRoute.Handler)
	}
}
