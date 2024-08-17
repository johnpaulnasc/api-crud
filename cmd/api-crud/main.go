package main

import (
    "log"
    "net/http"
    "api-crud/api/v1"
    "api-crud/config"
    "api-crud/pkg/db"
    "api-crud/internal/repository"
    "api-crud/internal/service"
    "api-crud/api/v1/controller"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    database, err := db.Connect(cfg)
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    itemRepo := repository.NewItemRepository(database)
    itemService := service.NewItemService(itemRepo)
    itemController := controller.NewItemController(itemService)

    router := v1.NewRouter(itemController)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}