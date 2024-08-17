package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"api-crud/config"
	"api-crud/pkg/db"
	"api-crud/internal/domain"
	"api-crud/internal/repository"
	"api-crud/internal/service"

	"gorm.io/gorm"
)

var database *gorm.DB

func main() {
	cfg := config.LoadConfig()

	var err error
	database, err = db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	database.AutoMigrate(&domain.Item{})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

	http.HandleFunc("/", renderIndexPage)

	http.HandleFunc("/items", handleItems)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderIndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("frontend/templates", "index.html"))
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	itemRepo := repository.NewItemRepository(database)
	itemService := service.NewItemService(itemRepo)

	switch r.Method {
	case "GET":
		items, err := itemService.GetAllItems()
		if err != nil {
			http.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
			return
		}
		tmpl, err := template.ParseFiles(filepath.Join("frontend/templates", "index.html"))
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, items)
	case "POST":
		name := r.FormValue("name")
		price := r.FormValue("price")
		item := domain.Item{
			Name:  name,
			Price: price,
		}
		err := itemService.CreateItem(&item)
		if err != nil {
			http.Error(w, "Failed to create item", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "PUT":
	case "DELETE":
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}