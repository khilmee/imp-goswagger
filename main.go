//	 Product Api:
//	  version: 0.1
//	  title: Product Api
//	 Schemes: http, https
//	 Host:
//	 BasePath: /api/v1
//		Consumes:
//		- application/json
//	 Produces:
//	 - application/json
//	 SecurityDefinitions:
//	  Bearer:
//	   type: apiKey
//	   name: Authorization
//	   in: header
//	 swagger:meta
package main

import (
	"fmt"
	"imp-goswagger/app/api/initialization"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)

		logger = level.NewFilter(logger, level.AllowAll())
		logger = level.NewInjector(logger, level.InfoValue())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	db, err := initialization.DbInit()
	if err != nil {
		_ = logger.Log("Err Db connection :", err.Error())
		panic(err.Error())
	}
	_ = logger.Log("message", "Connection Db Success")

	mux := initialization.InitRouting(db, logger)
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		_ = logger.Log("transport", "HTTP", "addr", 5600)
		errs <- http.ListenAndServe(fmt.Sprintf(":%d", 5600), nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	_ = logger.Log("exit", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
