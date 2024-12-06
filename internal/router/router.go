package router

import (
	"product/internal/config"
	"product/internal/handlers"
	"product/internal/repository"
	"net/http"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(cfg *config.Config, client *mongo.Client) http.Handler{

	repo := repository.NewMongoRepository(client, cfg.Database,cfg.Collection)
	handler := handler.NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/products", handler.GetAllProducts)
	r.Post("/products/create", handler.CreateProduct)
	r.Delete("/products/delete/{id}", handler.DeleteProduct)
	r.Put("/products/update/{id}", handler.UpdateProduct)
	return r
}