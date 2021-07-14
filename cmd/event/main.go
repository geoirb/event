package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fasthttp/router"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"

	"github.com/geoirb/event/pkg/builder"
	"github.com/geoirb/event/pkg/mongo"
	"github.com/geoirb/event/pkg/mongo/query"
	"github.com/geoirb/event/pkg/service"
	"github.com/geoirb/event/pkg/service/server/http"
	"github.com/geoirb/event/pkg/verificator"
)

type configuration struct {
	HttpPort string `envconfig:"HTTP_PORT" default:"8080"`

	TypeLayout string `envconfig:"TYPE_LAYOUT" default:"^[0-9a-z]*$"`

	StorageConnect    string `envconfig:"STORAGE_CONNECT" default:"mongodb://event:event@127.0.0.1:27017"`
	StorageDatabase   string `envconfig:"STORAGE_DATABASE" default:"event"`
	StorageCollection string `envconfig:"STORAGE_COLLECTION" default:"event"`
}

const (
	prefixCfg   = ""
	serviceName = "event"
)

func main() {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.WithPrefix(logger, "service", serviceName)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var cfg configuration
	if err := envconfig.Process(prefixCfg, &cfg); err != nil {
		level.Error(logger).Log("msg", "configuration", "err", err)
		os.Exit(1)
	}

	verificator, err := verificator.NewEvent(
		cfg.TypeLayout,
	)
	if err != nil {
		level.Error(logger).Log("msg", "init verificator", "err", err)
		os.Exit(1)
	}

	eventBuilder := builder.NewEvent(
		time.Now().Unix,
		time.Now().Unix,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryBuilder := query.NewBuilder()
	storage, err := mongo.NewStorage(
		ctx,
		cfg.StorageConnect,
		cfg.StorageDatabase,
		cfg.StorageCollection,
		queryBuilder,
	)
	if err != nil {
		level.Error(logger).Log("msg", "init mongo", "err", err)
		os.Exit(1)
	}

	svc := service.NewEvent(
		verificator,
		eventBuilder,
		storage,
		logger,
	)

	router := router.New()
	http.Routing(router, svc)

	httpServer := &fasthttp.Server{
		Handler:          router.Handler,
		DisableKeepalive: true,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	level.Info(logger).Log("msg", "received signal", "signal", <-c)

	if err := httpServer.Shutdown(); err != nil {
		level.Info(logger).Log("msg", "http server shoutdown", "err", err)
	}
}
