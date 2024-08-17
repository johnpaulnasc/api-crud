package v1

import (
    "net/http"
    "github.com/gorilla/mux"
    "api-crud/api/v1/controller"
    "api-crud/pkg/middleware"
)

func NewRouter(itemController *controller.ItemController) *mux.Router {
    router := mux.NewRouter()

    router.Use(middleware.LoggingMiddleware)

    router.HandleFunc("/api/v1/items", itemController.CreateItem).Methods("POST")
    // Outras rotas...

    return router
}