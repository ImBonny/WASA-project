
<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			usernameToSearch: "",
			id: localStorage.getItem("id"),
			user: "",
			profile: "",
			userProfile: {},
		}
	},
	created() {
		let username = localStorage.getItem("username");
		let id = localStorage.getItem("id");

		if (!username || !id) {
			// Redirect the user to the login page
			this.$router.push('/');
		} else {
			this.username = username;
			this.id = id;
		}
	},
	methods: {
		async searchUser() {
			try {
				let response = await this.$axios.get(`/users`, {
					params: {
						username: this.usernameToSearch
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.user = response.data;
				console.log("User found: " + this.user.Username);
			} catch (error) {
				this.errormsg = error.response.data;
			}
			if (this.usernameToSearch === this.username) {
				this.errormsg = "You can't search for yourself";
				return;
			}

			try {
				console.log("Checking if user is banned: " + this.username);
				let response = await this.$axios.get(`/utils/banned`, {
					headers: {
						Authorization: "Bearer " + this.id
					},
					params: {
						UsernameToCheck: this.username,
						UsernameBanning: this.usernameToSearch
					}
				},);
				console.log("Is banned:" + response.data.Banned);
				this.isBanned = response.data.Banned;
			} catch (error) {
				if (error.response) {
					this.errormsg = error.response.data;
				} else {
					console.error("Errore durante la richiesta:", error.message);
				}
			}
			if (this.isBanned){
				this.errormsg = "User Not Found";
			}
			else {

				try {
					console.log("Searching for profile");
					let response = await this.$axios.get(`/users/${this.usernameToSearch}/profiles`, {
						headers:
							{
								Authorization: "Bearer " + this.id
							}
					});
					this.profile = response.data.profile;
					console.log("Profile found: " + this.profile.Username);

					this.posts = response.data.profile.Posts;
					localStorage.setItem("profile", JSON.stringify(this.profile));
					console.log(localStorage.getItem("profile"));
					this.$router.push(`/users/${this.usernameToSearch}/profile`);
				} catch (error) {
					this.errormsg = error.response.data;
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
				//refresh the page
			} catch (error) {
				this.errormsg = error.response.data;
			}
		},
		async logout() {
			localStorage.clear();
			this.$router.push(`/session`);
		},
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
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Search a User</h1>
	</div>
	<div class="input-group mb-3">
		<input type="text" id="username" v-model="usernameToSearch" class="form-control search-input"
			   placeholder="Insert the username to search" aria-label="Recipient's username"
			   aria-describedby="basic-addon2">
		<div>
			<button class="search-button" type="button" @click="searchUser">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
				Search
			</button>
		</div>
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>

</style>
