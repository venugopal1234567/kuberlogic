package app

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
	apiService "github.com/kuberlogic/operator/modules/apiserver/internal/generated/restapi/operations/service"
	"github.com/kuberlogic/operator/modules/apiserver/internal/security"
	"github.com/kuberlogic/operator/modules/apiserver/util"
)

func (srv *Service) LogsGetHandler(params apiService.LogsGetParams, principal *models.Principal) middleware.Responder {
	ns, name, err := util.SplitID(params.ServiceID)
	if err != nil {
		return util.BadRequestFromError(err)
	}

	if authorized, err := srv.authProvider.Authorize(principal.Token, security.LogsGetSecGrant, params.ServiceID); err != nil {
		srv.log.Errorw("error checking authorization", "error", err)
		return apiService.NewLogsGetForbidden()
	} else if !authorized {
		return apiService.NewLogsGetForbidden()
	}

	m := srv.serviceStore.NewServiceObject(name, ns)
	logs, errLogs := srv.serviceStore.GetServiceLogs(m, params.ServiceInstance, *params.Tail, params.HTTPRequest.Context())
	if errLogs != nil {
		srv.log.Errorw("error getting service logs", "error", errLogs.Err)
		return apiService.NewLogsGetServiceUnavailable().WithPayload(&models.Error{Message: errLogs.ClientMsg})
	}

	return apiService.NewLogsGetOK().WithPayload(&models.Log{
		Body:  logs,
		Lines: *params.Tail,
	})
}