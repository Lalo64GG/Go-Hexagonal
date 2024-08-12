package server

import (
    "fmt"
    "log"
    "my-hexagonal-api/internal/platform/server/handler/health"
    "my-hexagonal-api/internal/platform/repository"
    userApp "my-hexagonal-api/internal/application/user" 
    "my-hexagonal-api/internal/domain/user"
    "github.com/gin-gonic/gin"
    "database/sql"
)

type Server struct {
    engine   *gin.Engine
    httpAddr string
    db       *sql.DB
}

func New(host, port string, db *sql.DB) *Server {
    engine := gin.Default()
    srv := &Server{
        engine:   engine,
        httpAddr: fmt.Sprintf("%s:%s", host, port),
        db:       db,
    }
    srv.registerRoutes()
    return srv
}

func (s *Server) registerRoutes() {
    s.engine.GET("/health", health.CheckHandler())

    // Inicializa las dependencias y crea los controladores
    userRepo := repository.NewPostgresUserRepository(s.db)
    userService := user.NewUserService(userRepo, s.db)
    userController := userApp.NewUserController(userService) 

    s.engine.POST("/users", userController.RegisterUser)
}

func (s *Server) Run() error {
    log.Printf("Starting server at %s", s.httpAddr)
    return s.engine.Run(s.httpAddr)
}
