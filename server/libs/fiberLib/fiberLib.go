package fiberLib

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type HttpServ struct {
	App    *fiber.App
	logger *zap.Logger
	host   string
	port   int
}

type HttpServProps struct {
	SvcName        string
	IsProd         bool
	Host           string
	Port           int
	Logger         *zap.Logger
	Config         fiber.Config
	DisableLog     bool // 默认打印日志，如果不需要可以设为true，自行定义
	DisableRecover bool // 默认使用recover，如果不需要可以设为true，自行定义
}

func NewHttpServ(props HttpServProps) *HttpServ {
	fiberConfig := props.Config
	if fiberConfig.ErrorHandler == nil {
		fiberConfig.ErrorHandler = ErrorHandler(props.Logger, props.IsProd)
	}
	fiberApp := fiber.New(fiberConfig)
	if props.DisableLog == false {
		fiberApp.Use(logger.New())
	}
	if props.DisableRecover == false {
		fiberApp.Use(Recover(props.Logger, props.IsProd))
	}
	fiberApp.Use(healthcheck.New())
	fiberApp.Use(requestid.New())
	httpServ := &HttpServ{
		App:    fiberApp,
		logger: props.Logger,
		host:   props.Host,
		port:   props.Port,
	}
	return httpServ
}

func (hs *HttpServ) Start() {
	// Listen from a different goroutine
	go func() {
		hs.logger.Info("启动HttpServ")
		addr := hs.host + ":" + strconv.Itoa(hs.port)
		if err := hs.App.Listen(addr); err != nil {
			log.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	_ = <-c                                         // This blocks the main thread until an interrupt is received
	log.Println("HttpServ: Gracefully shutting down...")
	err := hs.App.ShutdownWithTimeout(5 * time.Second)
	if err != nil {
		hs.logger.Error("shutdown", zap.Error(err))
	}
}
