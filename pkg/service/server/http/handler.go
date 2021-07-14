package http

import (
	"github.com/valyala/fasthttp"

	"github.com/geoirb/event/pkg/service"
)

type startServe struct {
	svc       service.Event
	transport *startEventTransport
}

func (s *startServe) Handler(ctx *fasthttp.RequestCtx) {
	info, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.Start(ctx, info)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newStartHandler(svc service.Event, transport *startEventTransport) fasthttp.RequestHandler {
	s := startServe{
		svc:       svc,
		transport: transport,
	}
	return s.Handler
}

type finishServe struct {
	svc       service.Event
	transport *finishEventTransport
}

func (s *finishServe) Handler(ctx *fasthttp.RequestCtx) {
	info, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.Finish(ctx, info)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newFinishHandler(svc service.Event, transport *finishEventTransport) fasthttp.RequestHandler {
	s := finishServe{
		svc:       svc,
		transport: transport,
	}
	return s.Handler
}
