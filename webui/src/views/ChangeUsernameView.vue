
<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			newUsername: "",
			username: localStorage.getItem("username") || "",
			id: localStorage.getItem("id"),
		}
	},
	methods: {
		async changeUsername(){
			try {
				let response = await this.$axios.put(`/users/${this.username}`, {

					newUsername: this.newUsername},{
					headers:
					{
						Authorization: "Bearer " + this.id
					}
			})
				this.username = response.data.newUsername;
				console.log("Username changed to:" + this.username)
				localStorage.setItem("username", this.username);
				this.$router.push("/session");

			} catch (error) {
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
		<input type="text" id="username" v-model="newUsername" class="form-control"
			   placeholder="Insert the new Username" aria-label="Recipient's username"
			   aria-describedby="basic-addon2">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="changeUsername">Change Username</button>
		</div>
	</div>
	<div class = "btn btn-success" @click="back">back</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>

</style>
