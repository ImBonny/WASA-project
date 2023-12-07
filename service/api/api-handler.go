package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	// API routes
	rt.router.POST("/session", rt.doLogin)                                                    // DONE
	rt.router.GET("/users/myStream", rt.getMyStream)                                          // DONE
	rt.router.GET("/users/profiles/:username", rt.getUserProfile)                             // DONE
	rt.router.PUT("/users/:username/profiles/:profile", rt.followUser)                        // DONE
	rt.router.DELETE("/users/:username/profiles/:profile", rt.unfollowUser)                   // DONE
	rt.router.PUT("/users/:username/banned", rt.banUser)                                      // DONE
	rt.router.DELETE("/users/:username/banned/:bannedUser", rt.unbanUser)                     // DONE
	rt.router.POST("/users/:username/posts/:postId/likes", rt.likePhoto)                      // DONE
	rt.router.DELETE("/users/:username/posts/:postId/likes/:likeId", rt.unlikePhoto)          // DONE
	rt.router.DELETE("/users/:username/posts/:postId", rt.deletePhoto)                        // DONE
	rt.router.POST("/users/:username/posts/:postId/comments", rt.commentPhoto)                // DONE
	rt.router.DELETE("/users/:username/posts/:postId/comments/:commentId", rt.uncommentPhoto) // DONE
	rt.router.POST("/users/:username/posts/", rt.uploadPhoto)                                 // DONE
	rt.router.PUT("/users/:username", rt.setMyUsername)                                       // DONE
	rt.router.GET("/users", rt.searchUser)                                                    // DONE
	return rt.router
}
