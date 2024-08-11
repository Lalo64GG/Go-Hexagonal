package bootstrap

import (
    "my-hexagonal-api/internal/platform/database"
    "my-hexagonal-api/internal/platform/server"
    "os"
)

func Run() error {
    // Obtener la cadena de conexión de la base de datos desde variables de entorno
    connStr := os.Getenv("DB_CONN")

    // Crear la conexión a la base de datos
    db, err := database.NewPostgresDB(connStr)
    if err != nil {
        return err
    }

    // Iniciar el servidor
    srv := server.New("localhost", "8080", db )
    return srv.Run()
}
