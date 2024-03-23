
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
				return;
			}
			else {

				try {
					console.log("Searching for profile");
					let response = await this.$axios.get(`/users/${this.usernameToSearch}/profiles`, {
						params: {
							username: this.usernameToSearch
						},
						headers:
							{
								Authorization: "Bearer " + this.id
							}
					});
					this.profile = response.data.profile;
					console.log("Profile found: " + this.profile.Username);
					console.log("Number of photos: " + this.profile.NumberOfPhotos);

					this.posts = response.data.profile.Posts;
					console.log("Posts found: " + response.data.profile.Posts);
					localStorage.setItem("profile", JSON.stringify(this.profile));
					console.log(localStorage.getItem("profile"));
					this.$router.push(`/users/${this.usernameToSearch}/profile`);
				} catch (error) {
					this.errormsg = error.response.data;
				}
			}

		}
	}

}
</script>

<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Welcome to WASAPhoto</h1>
	</div>
	<div class="input-group mb-3">
		<input type="text" id="username" v-model="usernameToSearch" class="form-control"
			   placeholder="Insert the username to search" aria-label="Recipient's username"
			   aria-describedby="basic-addon2">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="searchUser">Search</button>
		</div>
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>

</style>
