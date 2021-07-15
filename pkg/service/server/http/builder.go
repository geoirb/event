package http

import (
	"net/http"

	"github.com/fasthttp/router"

	"github.com/geoirb/event/pkg/service"
)

const (
	version = "/v1"

	startURI  = version + "/start"
	finishURI = version + "/finish"
)

func Routing(router *router.Router, svc service.Event) {
	router.Handle(http.MethodPost, startURI, newStartHandler(svc, NewStartTransport()))
	router.Handle(http.MethodPost, finishURI, newFinishHandler(svc, NewFinishTransport()))
}
