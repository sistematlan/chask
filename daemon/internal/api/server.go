package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"os"
)

// Request represents a JSON request from the panel.
type Request struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

// Response represents a JSON response to the panel.
type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// Server listens on a Unix socket and handles requests from the panel.
type Server struct {
	listener net.Listener
	handlers map[string]HandlerFunc
}

// HandlerFunc processes a request and returns response data or an error.
type HandlerFunc func(data json.RawMessage) (interface{}, error)

// NewServer creates a new Unix socket server.
func NewServer(socketPath string) (*Server, error) {
	// Remove stale socket file
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s: %w", socketPath, err)
	}

	// Set socket permissions: owner + group can read/write
	if err := os.Chmod(socketPath, 0660); err != nil {
		listener.Close()
		return nil, fmt.Errorf("failed to chmod socket: %w", err)
	}

	s := &Server{
		listener: listener,
		handlers: make(map[string]HandlerFunc),
	}
	s.registerHandlers()

	slog.Info("listening", "socket", socketPath)
	return s, nil
}

// Serve accepts connections in a loop.
func (s *Server) Serve() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConnection(conn)
	}
}

// Close shuts down the server.
func (s *Server) Close() {
	s.listener.Close()
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var req Request
	if err := decoder.Decode(&req); err != nil {
		slog.Error("failed to decode request", "error", err)
		encoder.Encode(Response{Status: "error", Error: "invalid request"})
		return
	}

	slog.Info("request", "action", req.Action)

	handler, ok := s.handlers[req.Action]
	if !ok {
		encoder.Encode(Response{Status: "error", Error: fmt.Sprintf("unknown action: %s", req.Action)})
		return
	}

	data, err := handler(req.Data)
	if err != nil {
		slog.Error("handler error", "action", req.Action, "error", err)
		encoder.Encode(Response{Status: "error", Error: err.Error()})
		return
	}

	encoder.Encode(Response{Status: "ok", Data: data})
}

func (s *Server) registerHandlers() {
	// System
	s.handlers["system.health"] = handleHealth
	s.handlers["system.metrics"] = handleMetrics
}
