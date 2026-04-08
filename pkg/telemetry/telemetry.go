package telemetry

import (
	"github.com/gin-gonic/gin"
)

type TelemetryData struct {
	Route      string `json:"route"`
	APIVersion string `json:"apiVersion"`
}

type telemetryService struct{}

func (t *telemetryService) TelemetryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		route := c.FullPath()
		go SendTelemetry(route)
		c.Next()
	}
}

type TelemetryService interface {
	TelemetryMiddleware() gin.HandlerFunc
}

func SendTelemetry(route string) {
	return
}

func NewTelemetryService() TelemetryService {
	return &telemetryService{}
}
