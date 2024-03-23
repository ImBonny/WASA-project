
<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: "",
			id: 0,
		}
	},
	methods: {
		async doLogin() {
			if (this.username == "") {
				this.errormsg = "Username cannot be empty.";
			} else {
				try{
					let response = await this.$axios.post("/session", {
						username: this.username})
					this.id = response.data.Identifier;
					localStorage.setItem("username", this.username);
					localStorage.setItem("id", this.id);
					this.$router.push(`/users/${this.username}/stream`);

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
		<input type="text" id="username" v-model="username" class="form-control"
			   placeholder="Insert a username to log in WASAPhoto." aria-label="Recipient's username"
			   aria-describedby="basic-addon2">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="doLogin">Login</button>
		</div>
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>

</style>
