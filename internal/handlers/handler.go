package handler

import (
	"encoding/json"
	"net/http"
	"product/internal/models"
	"product/internal/repository"
	"github.com/go-chi/chi/v5"
)

type Handler struct{
	repo repository.Repository
}

func NewHandler(repo repository.Repository)*Handler{
	return &Handler {repo: repo}
}

func (h *Handler)GetAllProducts(w http.ResponseWriter, r *http.Request){
	products, err := h.repo.GetAll()
	if err !=nil{
		http.Error(w,"Invalid request",http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *Handler)CreateProduct(w http.ResponseWriter, r *http.Request){

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err!=nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
    
	if err := h.repo.Create(product); err!=nil{
		http.Error(w, "Product not created", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message" : "Item created successfully",
		"product": product,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler)DeleteProduct(w http.ResponseWriter,r *http.Request){
	id := chi.URLParam(r, "id")
	err := h.repo.Delete(id)
	if err != nil{
		http.Error(w,"invalid request",http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"message" : "Item Deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func(h *Handler)UpdateProduct(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product) 
	if err!=nil{
		http.Error(w,"Invalid request",http.StatusInternalServerError)
		return
	}
    _,err = h.repo.Update(id,product)
 
   if err !=nil{
	http.Error(w, "productnot updated",http.StatusInternalServerError)
	return
   }

   w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Item updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}