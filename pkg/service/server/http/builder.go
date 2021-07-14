package http

import (
	"github.com/fasthttp/router"

	"github.com/geoirb/event/pkg/service"
)

const (
	version = "/v1"

	startURI  = version + "/start"
	finishURI = version + "/finish"
)

func Route(router *router.Router, svc service.Event) {

}
