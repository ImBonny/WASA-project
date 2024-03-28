package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type checkLikeRequest struct {
	TargetPost uint64 `json:"TargetPost"`
	LikeOwner  uint64 `json:"LikeOwner"`
}

type checkLikeResponse struct {
	Like bool `json:"Like"`
}

// Handler for checking if a post is liked
func (rt *_router) checkUserLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new checkLike request
	var checkLikeReq checkLikeRequest
	var err error
	token := getToken(r.Header.Get("Authorization"))
	//decode the body
	checkLikeReq.TargetPost, err = strconv.ParseUint(ps.ByName("postId"), 10, 64)
	checkLikeReq.LikeOwner = token
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth, err1 := rt.db.CheckAuthorization(token)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	like, err2 := rt.db.CheckUserLike(checkLikeReq.LikeOwner, checkLikeReq.TargetPost)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
	checkLikeResponse := checkLikeResponse{Like: like}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err3 := json.NewEncoder(w).Encode(checkLikeResponse)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
}
