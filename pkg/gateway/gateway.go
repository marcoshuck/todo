package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/marcoshuck/todo/pkg/conf"
	"github.com/marcoshuck/todo/pkg/telemetry"
	"net/http"
)

type Gateway struct {
	Telemeter telemetry.Telemetry
	mux       *runtime.ServeMux
	handler   http.Handler
	Config    conf.GatewayConfig
}
