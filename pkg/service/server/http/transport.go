package http

import (
	"encoding/json"

	"github.com/geoirb/event/pkg/service"
	"github.com/valyala/fasthttp"
)

type startEventTransport struct{}

func (t *startEventTransport) DecodeRequest(req *fasthttp.Request) (info service.StartEvent, err error) {
	var request eventStartRequest
	if err = json.Unmarshal(req.Body(), &request); err != nil {
		return
	}
	info = service.StartEvent{
		Type: request.Type,
	}
	return
}

func (t *startEventTransport) EncodeResponse(res *fasthttp.Response, err error) {
	code := getStatusCode(err)
	res.SetStatusCode(code)
}

func NewStartTransport() *startEventTransport {
	return &startEventTransport{}
}

type finishEventTransport struct{}

func (t *finishEventTransport) DecodeRequest(req *fasthttp.Request) (info service.FinishEvent, err error) {
	var request eventFinishRequest
	if err = json.Unmarshal(req.Body(), &request); err != nil {
		return
	}
	info = service.FinishEvent{
		Type: request.Type,
	}
	return
}

func (t *finishEventTransport) EncodeResponse(r *fasthttp.Response, err error) {
	code := getStatusCode(err)
	r.SetStatusCode(code)
}

func NewFinishTransport() *startEventTransport {
	return &startEventTransport{}
}
