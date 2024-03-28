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
			target: "",
			userProfile: {},

		}
	},
	methods: {
		async getMyStream() {
			try {
				let response = await this.$axios.get(`/users/${this.username}/stream`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
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
				try {
					//let response = await this.$axios.get(`/images/${this.stream[i].PostId}`, {});
					let post = this.stream[i];
					post.isLiked = await this.checkLike(post);
					this.stream[i] = post; // Update the existing post in the stream array
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
						LikeOwner: JSON.parse(this.id)
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Like removed successfully");
				post.NLikes = post.NLikes - 1;
				post.isLiked = false;
			} catch (error) {
				console.error("Error removing like:", error);
			}

		},
		async uncommentPost(post, comment) {
			try {
				await this.getUsername(post.PostOwner);
				let response = await this.$axios.delete(`users/${this.target}/posts/${post.PostId}/comments/${comment.CommentId}`, {
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
					}
				});
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
			try {
				console.log("Searching for profile");
				let response = await this.$axios.get(`/users/${this.username}/profiles`, {
					headers:
						{
							Authorization: "Bearer " + this.id
						}
				});
				this.userProfile = response.data.profile;
				console.log("Profile found: " + this.userProfile.Username);
				localStorage.setItem("profile", JSON.stringify(this.userProfile));
				this.$router.push(`/users/${this.username}/profile`);
			} catch (error) {
				this.errormsg = error.response.data;
			}
		},
		async logout() {
			localStorage.clear();
			this.$router.push(`/session`);
		},
		async loadProfile(username){
			try {
				console.log("Searching for profile");
				let response = await this.$axios.get(`/users/${username}/profiles`, {
					headers:
						{
							Authorization: "Bearer " + this.id
						}
				});
				this.userProfile = response.data.profile;
				console.log("Profile found: " + this.userProfile.Username);

				this.posts = response.data.profile.Posts;

				console.log("Posts found: " + response.data.profile.Posts);
				localStorage.setItem("profile", JSON.stringify(this.userProfile));
				console.log(localStorage.getItem("profile"));
				this.$router.push(`/users/${username}/profile`);
			} catch (error) {
				this.errormsg = error.response.data;
			}
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
				<span>Navigation</span>
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
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
						Search
					</RouterLink>
				</li>
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="myProfile">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
						My Profile
					</RouterLink>
				</li>
			</ul>

			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>Settings</span>
			</h6>
			<ul class="nav flex-column">
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="logout">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
						Logout
					</RouterLink>
				</li>
			</ul>
		</div>
	</nav>

	<div v-for="(post, index) in stream" :key="index" class="postBox">
		<div>
			<h2 style="margin-left: 10px;cursor: pointer" @click="loadProfile(usernames[post.PostOwner])">{{usernames[post.PostOwner]}}</h2>
			<img :src="`data:image/*;base64,${post.Image}`" alt="photo" class="post">
			<p class="description">{{ post.Description }}</p>
			<p class="dateTime">{{formatDate(post.CreationTime)}}</p>

			<div class="like-bar">
				<div class="input-group-append">
					<button class="like-button" @click="HandleLike(post)" v-if="!post.isLiked">
						<svg class="feather like"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
					</button>
				</div>
				<div class="input-group-append">
					<button class="like-button" @click="unlikePost(post)" v-if="post.isLiked">
						<svg class="feather like"><use href="/feather-sprite-v4.29.0.svg#heart" style="fill: red;"/></svg>
					</button>
				</div>

				<div class="input-group-append" style="margin-left: 10px">
					<p class=" like-counter">{{ post.NLikes }} likes</p>
				</div>
			</div>
			<div class="comment-bar">
				<div class="comment-input">
					<input type="text" class="form-control" placeholder="Add a comment..." v-model="comment">
				</div>
				<div class="input-group-append">
					<button class="comment-button" type="button" @click="HandleComment(post)">
						<svg class="feather addcomment"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg>
					</button>
				</div>
			</div>
		</div>

		<div class="comment-box">
		<div class="col-md-4" v-for="(comment, index) in post.Comments" :key="index">
			<div class="comment">
				<div style="display: flex;flex-direction: row">
					<p class="comment-user">{{usernames[comment.CommentOwner]}}: </p>
					<p class="comment-text">{{ comment.CommentText }}</p>
				</div>
				<p class="dateTime">{{formatDate(comment.CreationTime)}}</p>
				<button class="delete-button" type="button" @click="uncommentPost(post,comment)" v-if="comment.CommentOwner === JSON.parse(this.id)">
					<svg class="feather trash"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
				</button>
			</div>
		</div>
		</div>

	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>


</style>
