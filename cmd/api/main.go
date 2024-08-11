package main

import (
    "log"
    "my-hexagonal-api/cmd/api/bootstrap"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno desde el archivo .env
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Ejecutar el bootstrap de la aplicación
    if err := bootstrap.Run(); err != nil {
        log.Fatalf("Error running application: %v", err)
    }
}
