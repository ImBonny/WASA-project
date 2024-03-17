<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			profile: localStorage.getItem("profile"),
			id : localStorage.getItem("id"),
			isFollowing: false
		}
	},
	created() {
		this.checkFollowing();
		console.log("Is following: " + this.isFollowing);
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

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
