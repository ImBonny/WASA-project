<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			profile: localStorage.getItem("profile"),
			id: localStorage.getItem("id"),
			isFollowing: false,
			posts: [],
			comment: "",
			usernames: {},
		}
	},
	created() {
		this.checkFollowing();
		console.log("Is following: " + this.isFollowing);
		console.log("posts: " + JSON.parse(this.profile).Posts);
		this.loadImage()

	},
	methods: {
		formatDate(value) {
			let date = new Date(value);
			return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
		},
		async checkFollowing() {
			try {
				console.log("Checking if user is following: " + JSON.parse(this.profile).Username);
				let response = await this.$axios.get(`/utils/follows`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params: {
						username1: this.username,
						username2: JSON.parse(this.profile).Username
					}
				},);
				console.log(response.data.result);
				this.isFollowing = response.data.result;
			} catch (error) {
				if (error.response) {
					this.errormsg = error.response.data;
				} else {
					console.error("Errore durante la richiesta:", error.message);
				}
			}
		},
		async checkLike(post) {
			try {
				console.log("Checking if user liked post: " + post.PostId);
				let response = await this.$axios.get(`users/${JSON.parse(this.profile).Username}/posts/${post.PostId}/likes`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params:{
						TargetPost: post.PostId,
						LikeOwner: JSON.parse(this.id)
					}
				});
				console.log("response: " + response.data.Like);
				return response.data.Like;
			} catch (error) {
				if (error.response) {
					this.errormsg = error.response.data;
				} else {
					console.error("Errore durante la richiesta:", error.message);
				}
			}
		},
		async HandleFollow() {
			if (this.isFollowing) {
				this.UnfollowUser();
			} else {
				this.FollowUser();
			}
		},
		async FollowUser() {
			try {
				let response = await this.$axios.put(`/users/${JSON.parse(this.profile).Username}/profile`, {
					username: JSON.parse(this.profile).Username
				}, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.isFollowing = true;
				console.log("User followed: " + JSON.parse(this.profile).Username);
			} catch (error) {
				console.log("Error following user: " + JSON.parse(this.profile).Username);
				this.errormsg = error.response.data;
			}
		},
		async UnfollowUser() {
			try {
				let response = await this.$axios.delete(`/users/${JSON.parse(this.profile).Username}/profile`, {
					params: {
						username: JSON.parse(this.profile).Username,
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.isFollowing = false;
				console.log("User unfollowed: " + JSON.parse(this.profile).Username);
			} catch (error) {
				console.log("Error unfollowing user: " + JSON.parse(this.profile).Username);
				this.errormsg = error.response.data;
			}
		},
		async loadImage() {
			for (let i = 0; i < JSON.parse(this.profile).Posts.length; i++) {
				console.log("Loading image: " + JSON.parse(this.profile).Posts[i]);
				try {
					let response = await this.$axios.get(`/images/${JSON.parse(this.profile).Posts[i]}`, {});
					let post = response.data.image;
					post.isLiked = await this.checkLike(post);
					this.posts = [...this.posts, post];
					console.log("Image loaded: " + this.posts[i].Description);
					this.posts[i].Comments = await this.getComments(this.posts[i].PostId);
				} catch (error) {
					console.error("Error loading image:", error);
				}
			}
		},
		async getComments(postId) {
			try {
				let response = await this.$axios.get(`users/${JSON.parse(this.profile).Username}/posts/${postId}/comments`, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Comments loaded");
				// Load usernames for each comment
				for (let comment of response.data.Comments) {
					console.log("Loading username for comment: " + comment.CommentOwner);
					this.usernames[comment.CommentOwner] = await this.getUsername(comment.CommentOwner);
				}
				return response.data.Comments
			} catch (error) {
				console.error("Error loading comments:", error);
			}
		},
		async HandleComment(post) {
			if (this.comment === "") {
				this.errormsg = "Comment cannot be empty.";
			} else {
				try {
					let response = await this.$axios.post(`users/${JSON.parse(this.profile).Username}/posts/${post.PostId}/comments`, {
							CommentText: this.comment
						},
						{
							headers: {
								Authorization: "Bearer " + this.id
							}
						}
					);
					console.log("Comment added: " + response.data.CommentId);
					//reload the comments
					post.Comments = await this.getComments(post.PostId);
					return response.data.CommentId;
				} catch (error) {
					console.error("Error adding comment:", error);
				}
			}
		},
		async HandleLike(post) {
			try {
				let response = await this.$axios.post(`users/${JSON.parse(this.profile).Username}/posts/${post.PostId}/likes`, {
						TargetPost: post.PostId,
						LikeOwner: JSON.parse(this.id)
					},
					{
						headers: {
							Authorization: "Bearer " + this.id
						}
					}
				);
				console.log("Like added: " + response.data.Like.LikeOwner);
				post.NLikes = post.NLikes + 1;
				post.isLiked = true;
				return response.data.Like;
			} catch (error) {
				console.error("Error adding like:", error);
			}
		},
		async unlikePost(post) {
			try {
				let response = await this.$axios.delete(`users/${JSON.parse(this.profile).Username}/posts/${post.PostId}/likes`, {
					data: {
						TargetPost: post.PostId,
						LikeOwner: JSON.parse(this.id)
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Like removed successfully");
				post.NLikes = post.NLikes - 1;
				post.isLiked = false;
				return response.data;
			} catch (error) {
				console.error("Error removing like:", error);
			}

		},
		async uncommentPost(post, comment) {
			try {
				let response = await this.$axios.delete(`users/${JSON.parse(this.profile).Username}/posts/${post.PostId}/comments/${comment.CommentId}`, {
					data: {
						CommentId: comment.CommentId
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Comment removed successfully");
				post.Comments = await this.getComments(post.PostId);
				return response.data;
			} catch (error) {
				console.error("Error removing comment:", error);
			}
		},
		 async getUsername(userId) {
			 try {
				 let response = await this.$axios.get(`/utils/usernames`, {
					 headers: {
						 Authorization: "Bearer " + this.id
					 },
					 params: {
						 UserId: userId
					 }
				 });
				 console.log("Username loaded: " + response.data.Username)
				 return response.data.Username;
			 } catch (error) {
				 console.error("Error getting username:", error);

			 }
		 },
}
}
</script>


<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>{{JSON.parse(profile).Username}}'s Profile</h1>
	</div>
	<div>
		<h3>Please select an option from the buttons below:</h3>
	</div>
	<div class="input-group mb-3">
		<div class="input-group-append" style="margin-right: 10px">
			<button class="btn btn-success" type="button" @click="HandleFollow">{{ isFollowing ? 'Unfollow' : 'Follow' }}</button>
		</div>
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="Ban">Ban</button>
		</div>
	</div>
	<div>
		<h2>Profile Posts</h2>
		<div>
			<div v-for="(post, index) in posts" :key="index" class="postBox">
				<img :src="`data:image/*;base64,${post.Image}`" alt="photo" class="post">
				<p>{{ post.Description }}</p>
				<!-- bottone per aggiungere un commento -->
				<div class="input-group mb-3">
					<div class="input-group-append" style="margin-right: 10px">
						<button class="btn btn-success" type="button" @click="HandleLike(post)" v-if="!post.isLiked">Like</button>
					</div>
					<div class="input-group-append" style="margin-right: 10px">
						<button class="btn btn-success" type="button" @click="unlikePost(post)" v-if="post.isLiked">UnLike</button>
					</div>
					<!-- like counter -->
					<div class="input-group-append">
						<p>{{ post.NLikes }} likes</p>
					</div>
				</div>
				<div class="input-group mb-3">
					<div class="input-group-append" style="margin-right: 10px">
						<input type="text" class="form-control" placeholder="Comment" v-model="comment">
					</div>
					<div class="input-group-append">
						<button class="btn btn-success" type="button" @click="HandleComment(post)">Add Comment</button>
					</div>
				</div>
				<div>
					<h3>Comments</h3>
				</div>
				<!-- chiamare la funzione getComments(postId) -->
				<div class="col-md-4" v-for="(comment, index) in post.Comments" :key="index">
					<p class="comment">{{usernames[comment.CommentOwner]}}: {{ comment.CommentText }} <br>
						<small>{{formatDate(comment.CreationTime)}}</small>
					</p>

					<div class="input-group-append" v-if="comment.CommentOwner == id">
						<button class="btn btn-success" type="button" @click="uncommentPost(post,comment)">Delete Comment</button>
					</div>
					</div>

			</div>
		</div>
	</div>

	<!-- Display error message if any -->
	<ErrorMsg v-if="errormsg" :msg="errormsg" />

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>
	.postBox {
		background-color: #f0f0f0;
		padding: 10px;
		margin: 10px;
	}
	.post {
		max-width: 200px;
		max-height: 200px;
	}
	.comment {
		background-color: #f0f0f0;
		padding: 10px;
		margin: 10px;
	}

</style>
