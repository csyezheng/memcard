package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/csyezheng/memcard/pkg/logging"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ip       string
	port     string
	listener net.Listener
}

// NewServer create a new server but not start the server
func NewServer(port string) (*Server, error) {
	addr := fmt.Sprintf(":" + port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", addr, err)
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}

func (s *Server) Addr() string {
	return net.JoinHostPort(s.ip, s.port)
}

func (s *Server) ServeHTTP(ctx context.Context, srv *http.Server) error {
	logger := logging.FromContext(ctx)

	// Spawn a goroutine that listens for context closure. When the context is closed, the server is stopped.
	errCh := make(chan error, 1)
	go func() {
		// Listen for the interrupt signal.
		<-ctx.Done()

		logger.Debug("server.Serve: context closed")
		// The context is used to inform the server it has 5 seconds to finish
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		logger.Debug("shutting down gracefully, press Ctrl+C again to force")
		logger.Debug("server.Serve: shutting down")
		errCh <- srv.Shutdown(ctx)
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	logger.Debug("server.Serve: serving stopped")

	if err := <-errCh; err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           handler,
	})
}
