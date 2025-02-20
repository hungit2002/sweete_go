package cmd

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"sweete/user_service/config"
	"sweete/user_service/internal"
	"sweete/user_service/internal/middlewares"
	"sync"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			// init context
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			// init gin
			engine := gin.Default()

			engine.Use(middlewares.Cors)
			internal.Routes(engine)

			var wg sync.WaitGroup
			wg.Add(1)
			go runGin(ctx, engine, &wg)
			wg.Wait()
			log.Info("shutdown")
		},
	}
}

func runGin(ctx context.Context, engine *gin.Engine, wg *sync.WaitGroup) {
	srv := &http.Server{
		Addr:    ":" + config.GetInstance().App.Port,
		Handler: engine,
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("server listen err: %v", err)
		}
	}()
	<-sigint
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	wg.Done()
	log.Info("shutdown http server")
}
