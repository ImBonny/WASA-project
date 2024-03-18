
<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
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
						username: this.username
					},
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				this.user = response.data.username;
				console.log("User found: " + this.user);
			} catch (error) {
				this.errormsg = error.response.data;
			}

			try {
				console.log("Searching for profile");
				let response = await this.$axios.get(`/users/profiles/${this.username}`, {
					params: {
						username: this.username
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
				this.$router.push(`/users/${this.username}/profile`);
			}
			catch (error) {
				this.errormsg = error.response.data;
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
		<input type="text" id="username" v-model="username" class="form-control"
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
