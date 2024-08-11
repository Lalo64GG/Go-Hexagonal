package main

import (
    "log"
    "my-hexagonal-api/cmd/api/bootstrap"
)

func main() {
    if err := bootstrap.Run(); err != nil {
        log.Fatal(err)
    }
}
