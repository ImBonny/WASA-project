<script>
import {RouterLink} from "vue-router";

export default {
	components: {RouterLink},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			id: localStorage.getItem("id") || "",
			stream: [],
			comment: "",
			usernames: {},
			target: ""
		}
	},
	methods: {
		async getMyStream() {
			try {
				let response = await this.$axios.get(`/users/${this.username}/stream`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params: {
						Username: this.username
					}
				});
				console.log("Stream loaded");
				this.stream = response.data.Posts;
			} catch (error) {
				this.errormsg = error.response.data.message;
			}
		},
		async loadImage() {
			await this.getMyStream();
			for (let i = 0; i < this.stream.length; i++) {
				console.log("Loading image: " + this.stream[i]);
				try {
					let response = await this.$axios.get(`/images/${this.stream[i].PostId}`, {});
					let post = response.data.image;
					post.isLiked = await this.checkLike(post);
					this.stream[i] = post; // Update the existing post in the stream array
					console.log("Image loaded: " + this.stream[i].Description);
					console.log("Loading comments for post: " + this.stream[i]);
					this.stream[i].Comments = await this.getComments(this.stream[i]);
				} catch (error) {
					console.error("Error loading image:", error);
				}
			}
		},
		async getComments(post) {
			try {
				await this.getUsername(post.PostOwner);
				let response = await this.$axios.get(`users/${this.target}/posts/${post.PostId}/comments`, {
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
					await this.getUsername(post.PostOwner);
					let response = await this.$axios.post(`users/${this.target}/posts/${post.PostId}/comments`, {
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
					post.Comments = await this.getComments(post);
					return response.data.CommentId;
				} catch (error) {
					console.error("Error adding comment:", error);
				}
			}
		},
		async HandleLike(post) {
			try {
				await this.getUsername(post.PostOwner);
				let response = await this.$axios.post(`users/${this.target}/posts/${post.PostId}/likes`, {
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
				await this.getUsername(post.PostOwner);
				let response = await this.$axios.delete(`users/${this.target}/posts/${post.PostId}/likes`, {
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
				await this.getUsername(post.PostOwner);
				let response = await this.$axios.delete(`users/${this.target}/posts/${post.PostId}/comments/${comment.CommentId}`, {
					data: {
						CommentId: comment.CommentId
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Comment removed successfully");
				post.Comments = await this.getComments(post);
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
				this.target = response.data.Username;
				return response.data.Username;
			} catch (error) {
				console.error("Error getting username:", error);

			}
		},
		formatDate(date) {
			return new Date(date).toLocaleString();
		},
		async checkLike(post) {
			try {
				console.log("Checking if user liked post: " + post.PostId);
				await this.getUsername(post.PostOwner);
				let response = await this.$axios.get(`users/${this.username}/posts/${post.PostId}/likes`, {
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
		async home() {
			this.$router.push(`/users/${this.username}/stream`);
		},
		async search() {
			this.$router.push(`/users`);
		},
		async myProfile() {
			this.$router.push(`/users/${this.username}/profile`);
		},
		async logout() {
			localStorage.clear();
			this.$router.push(`/`);
		}

	},
	mounted() {
		this.getMyStream();
		this.loadImage();
	}

}
</script>


<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>{{username}}'s Stream</h1>
	</div>
	<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		<div class="position-sticky pt-3 sidebar-sticky">
			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>General</span>
			</h6>
			<ul class="nav flex-column">
				<li class="nav-item">
					<router-link to=""  class="nav-link" @click="home">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
						Home
					</router-link>
				</li>
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="search">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
						Search
					</RouterLink>
				</li>
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="myProfile">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
						My Profile
					</RouterLink>
				</li>
			</ul>

			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>Secondary menu</span>
			</h6>
			<ul class="nav flex-column">
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="logout">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
						Item 1
					</RouterLink>
				</li>
			</ul>
		</div>
	</nav>
	<!-- stream -->
	<div v-for="(post, index) in stream" :key="index" class="postBox">
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
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
