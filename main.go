package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/mdmims/go-echo-crud/config"
	"github.com/mdmims/go-echo-crud/db"
	"github.com/mdmims/go-echo-crud/handlers"
	"github.com/mdmims/go-echo-crud/repository"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"

	_ "github.com/mdmims/go-echo-crud/docs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

// @title Go Sample API
// @version 1.0
// @description Sample Go Rest API using Echo Framework and sqlite
// @host localhost:1323
// @BasePath /v1
func main() {
	// instantiate Echo
	e := echo.New()

	// setup rotating log file
	e.Logger.SetOutput(&lumberjack.Logger{
		Filename:   "./server.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     30, //days
	})

	// define log level
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)

	// middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// version endpoints
	v1 := e.Group("/v1")

	// initialize DB
	d, err := db.NewDB(config.New())
	if err != nil {
		panic(err)
	}
	defer d.Close()

	// setup TTL cache
	var cache = ttlcache.NewCache()
	c := repository.NewCacheStore(cache)
	c.SetTTL(600 * time.Second)
	c.SetCacheSizeLimit(500)

	// instantiate handlers
	items := repository.NewItemStore(d)

	// map handlers
	h := handlers.NewHandler(items, c)

	// register endpoints
	h.Register(v1)

	// start server and handle graceful shutdown with max timeout
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
