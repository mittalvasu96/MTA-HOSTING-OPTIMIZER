package main

import (
	"fmt"
	"log"
	"mta-hosting/env"
	getip "mta-hosting/getIP"
	"mta-hosting/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func init() {
	// base.InitKeyDB()
	getip.InitDataStore()
}

func main() {
	var g errgroup.Group

	g.Go(func() error {
		return startAPIServer()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func startAPIServer() (err error) {
	apiHost := env.Env["API_HOST"]
	apiPort := env.Env["API_PORT"]
	apiPath := env.Env["API_PATH"]

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()
	router.Handle("GET", "/healthCheckup", service.HealthCheck)
	router.Handle("GET", "/mta-hosting-optimizer/", service.GetMTAHost)
	// api_group := router.Group(apiPath)
	// {
	// 	api_group.GET("/", helper.GetMTAHost)
	// }
	apiServer := &http.Server{
		Addr:         apiHost + ":" + apiPort,
		Handler:      router,
		ReadTimeout:  time.Second * 50,
		WriteTimeout: time.Second * 50,
	}
	fmt.Println("Starting API server on " + apiHost + ":" + apiPort + apiPath)
	err = apiServer.ListenAndServe()
	return
}
