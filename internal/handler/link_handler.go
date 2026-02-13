package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fririz/URLShortener/internal/dto"
)

type LinkService interface {
	CreateLink(linkDto dto.LinkDto) (string, error)
	GetLinkBySlug(slug string) (string, error)
}

type LinkHandler struct {
	ls LinkService
}

func NewLinkHandler(service LinkService) *LinkHandler {
	return &LinkHandler{ls: service}
}

func (lh *LinkHandler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var linkDto dto.LinkDto
	if err := json.NewDecoder(r.Body).Decode(&linkDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortLink, err := lh.ls.CreateLink(linkDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortLink,
	})
}

func (lh *LinkHandler) GetFullUrl(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("link")

	fullUrl, err := lh.ls.GetLinkBySlug(slug)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, fullUrl, http.StatusFound)
}
