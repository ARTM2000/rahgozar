package http

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log/slog"
	"time"
)

type Controller interface {
	InitRoutes(app *fiber.App)
}

type Server interface {
	Start()
	Stop(force bool) error
	RegisterController(controllers ...Controller)
}

type Config struct {
	Host string
	Port string
}

type httpServer struct {
	config Config
	app    *fiber.App
}

func NewHTTPServer(config Config) Server {
	return &httpServer{
		config: config,
	}
}

func (h *httpServer) Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// check that if error was a fiber NewError and got status code,
			// specify that in error handler
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
			slog.Error("unhandled error occurred", slog.Any("error", err))

			return c.Status(code).JSON(FormatResponse(c, ResponseData{
				Message: err.Error(),
				IsError: true,
			}))
		},
	})

	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(requestid.New(requestid.Config{
		Next: func(c *fiber.Ctx) bool {
			trackId := c.Get(fiber.HeaderXRequestID)
			if trackId != "" {
				c.Set(fiber.HeaderXRequestID, trackId)
				return true
			}
			return false
		},
	}))
	app.Use(compress.New(compress.Config{Level: compress.LevelBestCompression}))
	app.Hooks().OnListen(func(ld fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}
		scheme := "http"
		if ld.TLS {
			scheme = "https"
		}
		slog.Info(fmt.Sprintf("server start listening on '%s'", scheme+"://"+ld.Host+":"+ld.Port))
		return nil
	})

	go func() {
		if err := app.Listen(fmt.Sprintf("%s:%s", h.config.Host, h.config.Port)); err != nil {
			slog.Error(err.Error())
		}
	}()

	h.app = app
}

func (h *httpServer) Stop(force bool) error {
	if h.app == nil {
		panic("http app server not defined. first start the http server")
	}
	if force {
		return h.app.Shutdown()
	}
	return h.app.ShutdownWithTimeout(time.Second * 30)
}

func (h *httpServer) RegisterController(controllers ...Controller) {
	if h.app == nil {
		panic("http app server not defined. first start the http server")
	}
	for _, c := range controllers {
		c.InitRoutes(h.app)
	}
}
