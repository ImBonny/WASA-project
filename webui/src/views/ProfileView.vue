<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			profile: localStorage.getItem("profile"),
			id : localStorage.getItem("id"),
		}
	},
	methods: {
		async Follow() {
			try {
				 let response = await this.$axios.put(`/users/${JSON.parse(this.profile).Username}/profile`, {
					 username: JSON.parse(this.profile).Username}, {
					 headers: {
						 Authorization: "Bearer " + this.id
					 }
				});
				console.log("User followed: " + JSON.parse(this.profile).Username);
			} catch (error) {
				console.log("Error following user: " + JSON.parse(this.profile).Username);
				this.errormsg = error.response.data;
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
			<button class="btn btn-success" type="button" @click="Follow">Follow</button>
		</div>
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="Ban">Ban</button>
		</div>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
