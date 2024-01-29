package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Upload a photo to the current user's profile
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	token := getToken(r.Header.Get("Authorization"))

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	uploadedImage := r.URL.Query().Get("image")
	caption := r.URL.Query().Get("caption")
	photoId, err := rt.db.UploadPhoto(token, uploadedImage, caption)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(photoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
