<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			profile: localStorage.getItem("profile"),
			id : localStorage.getItem("id"),
			isFollowing: false,
			posts: [],
			comment: "",
		}
	},
	created() {
		this.checkFollowing();
		console.log("Is following: " + this.isFollowing);
		console.log("posts: " + JSON.parse(this.profile).Posts);
		this.loadImage()

	},
	methods: {
		async checkFollowing(){
			try{
				console.log("Checking if user is following: " + JSON.parse(this.profile).Username);
				let response = await this.$axios.get(`/utils/follows`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params:{
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
		async HandleFollow(){
			if (this.isFollowing) {
				this.UnfollowUser();
			} else {
				this.FollowUser();
			}
		},
		async FollowUser() {
			try {
				 let response = await this.$axios.put(`/users/${JSON.parse(this.profile).Username}/profile`, {
					 username: JSON.parse(this.profile).Username}, {
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
		async UnfollowUser(){
			try {
				let response = await this.$axios.delete(`/users/${JSON.parse(this.profile).Username}/profile`, {
					params:{
					username: JSON.parse(this.profile).Username,},
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
		async loadImage(){
			for (let i = 0; i < JSON.parse(this.profile).Posts.length; i++) {
				console.log("Loading image: " + JSON.parse(this.profile).Posts[i]);
				try {
					let response = await this.$axios.get(`/images/${JSON.parse(this.profile).Posts[i]}`, {
					});
					this.posts = [...this.posts, response.data.image];
					console.log("Image loaded: " + this.posts[i].Description);
					this.posts[i].Comments = await this.getComments(this.posts[i].PostId);
				} catch (error) {
					console.error("Error loading image:", error);
				}
			}
		},
		async getComments(postId){
			try {
				let response = await this.$axios.get(`users/${JSON.parse(this.profile).Username}/posts/${postId}/comments`, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Comments loaded");
				return response.data.Comments
			} catch (error) {
				console.error("Error loading comments:", error);
			}
		},
		async HandleComment(post){
			if (this.comment == "") {
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
					return response.data.CommentId;
				} catch (error) {
					console.error("Error adding comment:", error);
				}
			}
		}
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
			<div class="col-md-4" v-for="(post, index) in posts" :key="index">
				<img :src="`data:image/*;base64,${post.Image}`" alt="photo">
				<p>{{ post.Description }}</p>
				<!-- bottone per aggiungere un commento -->
				<div class="input-group mb-3">
					<div class="input-group-append" style="margin-right: 10px">
						<button class="btn btn-success" type="button" @click="HandleLike(post)">Like</button>
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
					<p>{{ comment.CommentText }}</p>

				</div>

			</div>
		</div>
	</div>

	<!-- Display error message if any -->
	<ErrorMsg v-if="errormsg" :msg="errormsg" />

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
