package route

import (
	"codefast_2024/app"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
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
		value, _ := httputil.DumpRequest(c.Request, true)
		fmt.Println("---- NO Route LOG Start ----")
		fmt.Println("Request Time: " + time.Now().Format("2006/01/02 - 15:04:05"))
		fmt.Println("Request IP: " + c.ClientIP())
		fmt.Print(string(value))
		fmt.Println("---- NO Route LOG End ----")
		c.Status(http.StatusBadGateway)
	}
}
