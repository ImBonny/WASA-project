<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			profile: localStorage.getItem("profile") || "",
			id: localStorage.getItem("id") || "",
			isFollowing: false,
			isBanned: false,
			posts: [],
			comment: "",
			usernames: {},
			userProfile: {},
			ismyProfile: false,
			selectedFile: null,
			caption: "",
			followers: [],
			following: [],
		}
	},
	mounted() {
		if (!this.ismyProfile) {
			this.checkFollowing();
			this.checkBanned();
			console.log("Is following: " + this.isFollowing);
		}
		this.getFollowers()
		this.getFollowing()
		this.loadImage()
	},
	created() {
		this.checkIsMyProfile();
	},
	methods: {
		formatDate(value) {
			let date = new Date(value);
			return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
		},
		async checkIsMyProfile() {
			if (JSON.parse(this.profile).Username === this.username) {
				this.ismyProfile = true;
			}
		},
		async checkBanned() {
			try {
				console.log("Checking if user is banned: " + JSON.parse(this.profile).Username);
				let response = await this.$axios.get(`/utils/banned`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params: {
						UsernameToCheck: JSON.parse(this.profile).Username,
						UsernameBanning: this.username
					}
				},);
				console.log("IS banned:" + response.data.Banned);
				this.isBanned = response.data.Banned;
			} catch (error) {
				if (error.response) {
					this.errormsg = error.response.data;
				} else {
					console.error("Errore durante la richiesta:", error.message);
				}
			}
		},
		async checkFollowing() {
			try {
				console.log("Checking if user is following: " + JSON.parse(this.profile).Username);
				let response = await this.$axios.get(`/users/${JSON.parse(this.profile).Username}/followers/${this.username}`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params: {
						username1: this.username,
						username2: JSON.parse(this.profile).Username
					}
				},);
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
					params: {
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
				await this.UnfollowUser();
			} else {
				await this.FollowUser();
			}
		},
		async FollowUser() {
			try {
				await this.$axios.put(`/users/${JSON.parse(this.profile).Username}/profile`, {
					username: JSON.parse(this.profile).Username
				}, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.isFollowing = true;
				console.log("User followed: " + JSON.parse(this.profile).Username);
				location.reload()
			} catch (error) {
				this.errormsg = error.response.data;
			}
		},
		async UnfollowUser() {
			try {
				await this.$axios.delete(`/users/${JSON.parse(this.profile).Username}/profile`, {
					params: {
						username: JSON.parse(this.profile).Username,
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.isFollowing = false;
				console.log("User unfollowed: " + JSON.parse(this.profile).Username);
				location.reload()
			} catch (error) {
				this.errormsg = error.response.data;
			}
		},
		async loadImage() {
			for (let i = 0; i < JSON.parse(this.profile).Posts.length; i++) {
				console.log("Loading image: " + JSON.parse(this.profile).Posts[i]);
				try {
					let post = JSON.parse(this.profile).Posts[i];
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
			} catch (error) {
				console.error("Error removing like:", error);
			}

		},
		async uncommentPost(post, comment) {
			try {
				await this.$axios.delete(`users/${JSON.parse(this.profile).Username}/posts/${post.PostId}/comments/${comment.CommentId}`, {
					data: {
						CommentId: comment.CommentId
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Comment removed successfully");
				post.Comments = await this.getComments(post.PostId);
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
		async HandleBan() {
			if (this.isBanned) {
				await this.UnbanUser();
			} else {
				await this.BanUser();
			}
		},
		async BanUser() {
			try {
				await this.$axios.put(`/users/${this.username}/banned`, {
					BannedUser: JSON.parse(this.profile).Username
				}, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.isBanned = true;
				console.log("User banned: " + JSON.parse(this.profile).Username);
			} catch (error) {
				this.errormsg = error.response.data;
			}
		},
		async UnbanUser() {
			try {
				await this.$axios.delete(`/users/${this.username}/banned/${JSON.parse(this.profile).Username}`, {
					params: {
						BannedUser: JSON.parse(this.profile).Username,
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.isBanned = false;
				console.log("User unbanned: " + JSON.parse(this.profile).Username);
			} catch (error) {
				console.log("Error unbanning user: " + JSON.parse(this.profile).Username);
				this.errormsg = error.response.data;
			}
		},
		async deletePost(post) {
			try {
				await this.$axios.delete(`users/${this.username}/posts/${post.PostId}`, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("Post deleted: " + post.PostId);
				this.posts = this.posts.filter(p => p.PostId !== post.PostId);
			} catch (error) {
				console.error("Error deleting post:", error);
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
					params: {
						username: this.username
					},
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
				this.$router.push(`/users/${this.username}/profile`);
				//refresh the page
				location.reload();
			} catch (error) {
				this.errormsg = error.response.data;
			}
		},
		async logout() {
			localStorage.clear();
			this.$router.push(`/session`);
		},
		async UploadPhoto() {
			// Crea un elemento input di tipo file
			let input = document.createElement("input");
			input.type = "file";
			input.accept = "image/*";

			// Aspetta che venga selezionato un file prima di andare avanti
			let selectedFilePromise = new Promise((resolve) => {
				input.onchange = () => resolve(input.files[0]);
			});

			// Simula il clic sull'elemento input per aprire la finestra di dialogo di selezione del file
			input.click();

			// Attendi che l'utente selezioni un file
			this.selectedFile = await selectedFilePromise;

			// Leggi il contenuto del file come array di byte
			let reader = new FileReader();
			reader.readAsArrayBuffer(this.selectedFile);

			// Attendi che il reader completi la lettura del file
			let fileDataPromise = new Promise((resolve) => {
				reader.onload = () => resolve(reader.result);
			});

			// Ottieni i dati binari del file
			let fileData = await fileDataPromise;

			//Aggiungi la caption
			this.caption = prompt("Inserisci una caption per la foto");
			if (this.caption == null) {
				this.caption = "";
			}

			// Invia i dati binari del file al backend
			try {
				console.log("Uploading file: " + this.selectedFile.name);
				let response = await this.$axios.post(`/users/${this.username}/posts`, {
					image: Array.from(new Uint8Array(fileData)),  // Converti i dati del file in un array di byte
					caption: this.caption
				}, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("File uploaded: " + response.data);
				await this.myProfile();
			} catch (error) {
				console.log("Error uploading file: " + error.response.data);
				this.errormsg = error.response.data;
			}


		},
		async changeUsername() {
			this.newUsername = prompt("Enter new username");
			if (this.newUsername === null || this.newUsername === "") {
				this.errormsg = "Username cannot be empty.";
				return;
			}
			try {
				let response = await this.$axios.put(`/users/${this.username}`, {

					newUsername: this.newUsername
				}, {
					headers:
						{
							Authorization: "Bearer " + this.id
						}
				})
				this.username = response.data.newUsername;
				console.log("Username changed to:" + this.username)
				localStorage.setItem("username", this.username);
				await this.myProfile();

			} catch (error) {
				this.errormsg = error.response.data;
			}

		},
		async getFollowers() {
			try {
				let response = await this.$axios.get(`/users/${JSON.parse(this.profile).Username}/followers`, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				if (response.data.Followers === null) {
					this.followers = [];
				}
				else {
					this.followers = response.data.Followers;
				}
			} catch (error) {
				console.error("Error getting followers:", error);
			}
		},
		async getFollowing() {
			try {
				let response = await this.$axios.get(`/users/${JSON.parse(this.profile).Username}/following`, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				if (response.data.Followings === null || response.data.Followings === undefined) {
					this.following = [];
				}
				else {
					this.following = response.data.Followings;
				}
			} catch (error) {
				console.error("Error getting following:", error);
			}
		}
	}
}
</script>


<template>
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
	<div class="profile-header">
		<h1>{{JSON.parse(profile).Username}}'s Profile</h1>
		<div class="counters"> Followers: {{followers.length>0 ? followers.length : 0}}<br>
		Followings: {{following.length>0 ? following.length : 0}}</div>
		<div class="input-group-append" style="margin-right: 10px" v-if="ismyProfile">
			<button class="changeUsername-button" type="button" @click="changeUsername" >Change Username</button>
		</div>
	</div>
	<div class="input-group mb-3">
		<div class="input-group-append" style="margin-right: 10px" v-if="!ismyProfile">
			<button class="follow-button" type="button" @click="HandleFollow">{{ isFollowing ? 'Unfollow' : 'Follow' }}</button>
		</div>
		<div class="input-group-append" v-if="!ismyProfile">
			<button class="ban-button" type="button" @click="HandleBan">{{ isBanned ? 'UnBan' : 'Ban' }}</button>
		</div>
	</div>
	<div>
		<h2>Number of posts: {{posts.length}}</h2>
		<div class="input-group-append" v-if="ismyProfile">
			<button class="upload-button" type="button" @click="UploadPhoto">Upload a Post</button>
		</div>
		<div v-for="(post, index) in posts" :key="index" class="postBox">
			<div>
				<div class="input-group-append" v-if="ismyProfile">
					<button class="deletePost-button" type="button" @click="deletePost(post)">
						<svg class="feather trash"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
					</button>
				</div>
				<img :src="`data:image/*;base64,${post.Image}`" alt="photo" class="post">
				<p class="description">{{ post.Description }}</p>
				<p class="dateTime">{{formatDate(post.CreationTime)}}</p>
				<!-- bottone per aggiungere un commento -->
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
					<!-- like counter -->
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
			<!-- commenti -->
			<div class="comment-box">
				<div class="col-md-4" v-for="(comment, index) in post.Comments" :key="index">
					<div class="comment">
						<div style="display: flex;flex-direction: row">
							<p class="comment-user">{{usernames[comment.CommentOwner]}}: </p>
							<p class="comment-text">{{ comment.CommentText }}</p>
						</div>
						<p class="dateTime">{{formatDate(comment.CreationTime)}}</p>
						<button class="delete-button" type="button" @click="uncommentPost(post,comment)" v-if="comment.CommentOwner==this.id">
							<svg class="feather trash"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
						</button>
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
</style>
