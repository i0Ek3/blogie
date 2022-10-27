package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/routers"
	"github.com/i0Ek3/blogie/pkg/shutdown"
)

func Boot() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: global.ServerSetting.HeaderBytes,
	}

	go func() {
		err := ser.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("main::ser.ListenAndServe err: %v", err)
		}
	}()

	shutdown.Quit(ser)
}
