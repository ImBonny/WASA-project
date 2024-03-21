package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type getCommentsRequest struct {
	PostId uint64 `json:"postId"`
}

type getCommentsResponse struct {
	Comments []database.Database_comment `json:"Comments"`
}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Get the post id
	var request getCommentsRequest
	var err error
	request.PostId, err = strconv.ParseUint(p.ByName("postId"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid post id", http.StatusBadRequest)
		return
	}
	// Get the comments
	comments, err := rt.db.GetComments(request.PostId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var response getCommentsResponse
	response.Comments = *comments
	// Send the comments
	w.Header().Set("Content-Type", "application/json")
	err1 := json.NewEncoder(w).Encode(response)
	if err1 != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
