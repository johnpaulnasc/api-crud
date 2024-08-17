package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"api-crud/config"
	"api-crud/internal/domain"
	"api-crud/internal/repository"
	"api-crud/internal/service"
	"api-crud/pkg/db"

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
		priceStr := r.FormValue("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			http.Error(w, "Invalid price format", http.StatusBadRequest)
			return
		}
		idStr := r.FormValue("id")
		if idStr != "" { // Se o ID existir, é uma atualização
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			item := domain.Item{
				ID:    uint(id),
				Name:  name,
				Price: price,
			}
			err = itemService.UpdateItem(&item)
			if err != nil {
				http.Error(w, "Failed to update item", http.StatusInternalServerError)
				return
			}
		} else { // Se não, é uma criação
			item := domain.Item{
				Name:  name,
				Price: price,
			}
			err = itemService.CreateItem(&item)
			if err != nil {
				http.Error(w, "Failed to create item", http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "DELETE":
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		err = itemService.DeleteItem(uint(id))
		if err != nil {
			http.Error(w, "Failed to delete item", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
