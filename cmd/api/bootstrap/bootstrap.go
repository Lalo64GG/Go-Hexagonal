package bootstrap

import (
    "log"
    "my-hexagonal-api/internal/platform/server"
)

func Run() error {
    srv := server.New("localhost", "8080")
    return srv.Run()
}
