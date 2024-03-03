package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Upload a photo to the current user's profile
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	_ = ps
	token := getToken(r.Header.Get("Authorization"))
	auth, err := rt.db.CheckAuthorization(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	uploadedImage := r.URL.Query().Get("image")
	caption := r.URL.Query().Get("caption")
	photoId, err1 := rt.db.UploadPhoto(token, uploadedImage, caption)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	err2 := json.NewEncoder(w).Encode(photoId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
