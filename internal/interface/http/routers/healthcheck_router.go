package routers

import (
	"net/http"

	"github.com/andreis3/catalog-write-api/internal/interface/http/controllers"
	"github.com/andreis3/catalog-write-api/internal/interface/http/helpers"
)

type HealthCheckRouter struct{}

func NewHealthCheckRoutes() *HealthCheckRouter {
	return &HealthCheckRouter{}
}
func (r *HealthCheckRouter) HealthCheckRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/healthcheck",
			Controller:  controllers.HealthCheck,
			Description: "Health Check",
			Type:        helpers.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
