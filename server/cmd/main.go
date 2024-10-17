package main

import (
	goutils "github.com/ARTM2000/goutils"
	"github.com/ARTM2000/rahgozar/internal/controller/http"
	"github.com/ARTM2000/rahgozar/internal/core/common"
	"github.com/ARTM2000/rahgozar/internal/core/service"

	"log/slog"
)

func main() {
	slog.SetDefault(common.NewLogger(slog.LevelDebug))

	h := http.NewHTTPServer(http.Config{Host: "localhost", Port: "5500"})
	h.Start()
	h.RegisterController(
		http.NewMapLayerController(
			service.NewMapLayersService(),
		),
	)

	goutils.OnInterrupt(func() {
		if err := h.Stop(false); err != nil {
			slog.Error("fail to stop the http server", slog.Any("error", err))
		}
	})
}
