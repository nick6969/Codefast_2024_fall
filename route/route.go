package route

import (
	"codefast_2024/app"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func NewServer(app *app.App) *Server {
	engine := gin.Default()
	server := &http.Server{
		Addr:    ":5001",
		Handler: engine,
	}

	registerHandlers(engine, app)

	engine.NoRoute(NoRouter())

	return &Server{
		server: server,
		engine: engine,
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown() {
	log.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Server exiting")
}

func NoRouter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
