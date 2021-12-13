/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */

package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/ngalert/metrics"
)

type PrometheusApiForkingService interface {
	RouteGetAlertStatuses(*models.ReqContext) response.Response
	RouteGetRuleStatuses(*models.ReqContext) response.Response
}

type PrometheusApiService interface {
	RouteGetAlertStatuses(*models.ReqContext) response.Response
	RouteGetRuleStatuses(*models.ReqContext) response.Response
}

func (f *ForkedPrometheusApi) RouteGetAlertStatuses(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetAlertStatuses(ctx)
}

func (f *ForkedPrometheusApi) RouteGetRuleStatuses(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetRuleStatuses(ctx)
}

func (api *API) RegisterPrometheusApiEndpoints(srv PrometheusApiForkingService, m *metrics.API) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Get(
			toMacaronPath("/api/prometheus/{Recipient}/api/v1/alerts"),
			metrics.Instrument(
				http.MethodGet,
				"/api/prometheus/{Recipient}/api/v1/alerts",
				srv.RouteGetAlertStatuses,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/prometheus/{Recipient}/api/v1/rules"),
			metrics.Instrument(
				http.MethodGet,
				"/api/prometheus/{Recipient}/api/v1/rules",
				srv.RouteGetRuleStatuses,
				m,
			),
		)
	}, middleware.ReqSignedIn)
}
