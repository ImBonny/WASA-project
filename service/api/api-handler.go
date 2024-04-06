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

	rt.router.POST("/session", rt.doLogin) // DONE
	// API routes
	rt.router.GET("/users", rt.searchUser)                   // DONE
	rt.router.GET("/users/:username/stream", rt.getMyStream) // DONE
	rt.router.GET("/users/:username/profiles", rt.getUserProfile)
	rt.router.PUT("/users/:username", rt.setMyUserName) // DONE
	rt.router.PUT("/users/:username/profile", rt.followUser)
	rt.router.DELETE("/users/:username/profile", rt.unfollowUser)
	rt.router.PUT("/users/:username/banned", rt.banUser)                                      // DONE
	rt.router.DELETE("/users/:username/banned/:bannedUser", rt.unbanUser)                     // DONE
	rt.router.POST("/users/:username/posts", rt.uploadPhoto)                                  // DONE
	rt.router.DELETE("/users/:username/posts/:postId", rt.deletePhoto)                        // DONE
	rt.router.POST("/users/:username/posts/:postId/comments", rt.commentPhoto)                // DONE
	rt.router.GET("/users/:username/posts/:postId/comments", rt.getComments)                  // DONE
	rt.router.POST("/users/:username/posts/:postId/likes", rt.likePhoto)                      // DONE
	rt.router.DELETE("/users/:username/posts/:postId/comments/:commentId", rt.uncommentPhoto) // DONE
	rt.router.DELETE("/users/:username/posts/:postId/likes", rt.unlikePhoto)                  // DONE
	rt.router.GET("/users/:username/posts/:postId/likes", rt.checkUserLike)                   // DONE
	rt.router.GET("/users/:username/followers", rt.getFollowers)                              // DONE
	rt.router.GET("/users/:username/following", rt.getFollowing)                              // DONE
	rt.router.GET("/users/:username/followers/:usernameFollowing", rt.getFollows)             // DONE
	rt.router.GET("/utils/usernames", rt.searchUserById)                                      // DONE
	rt.router.GET("/utils/banned", rt.getBanned)                                              // DONE 	// DONE

	return rt.router
}
