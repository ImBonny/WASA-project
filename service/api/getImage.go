package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getImageRequest = struct {
	ImageId string `json:"imageId"`
}

type getImageResponse = struct {
	Image database.Database_photo `json:"image"`
}

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the image id from the request
	var request getImageRequest
	request.ImageId = ps.ByName("imageId")

	// Get the image from the database
	image, err := rt.db.GetImage(request.ImageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response getImageResponse
	response.Image = image

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err1 := json.NewEncoder(w).Encode(response)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
}
