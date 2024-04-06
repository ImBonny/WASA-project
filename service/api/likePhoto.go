package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type LikeRequest struct {
	TargetPost uint64 `json:"TargetPost"`
	LikeOwner  uint64 `json:"LikeOwner"`
}

type LikeResponse struct {
	Like database.Database_like `json:"Like"`
}

// Handler for liking a post
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new like request
	var likeReq LikeRequest

	err0 := json.NewDecoder(r.Body).Decode(&likeReq)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}
	var err error
	likeReq.TargetPost, err = strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token := getToken(r.Header.Get("Authorization"))

	auth, e := rt.db.CheckAuthorization(token)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	err1 := rt.db.LikePhoto(likeReq.TargetPost, token)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	likeResponse := LikeResponse{Like: database.Database_like{PostId: likeReq.TargetPost, LikeOwner: likeReq.LikeOwner}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err2 := json.NewEncoder(w).Encode(likeResponse)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
