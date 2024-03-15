<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
		}
	},
	methods: {
		async changeUsername(){
			let username = this.username
			this.$router.push(`/users/${username}`);
		},
		async searchUser(){
			let username = this.username
			this.$router.push('/users');
		}
	}
}
</script>


<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>Welcome {{username}}!</h1>
	</div>
	<div>
		<h3>Please select an option from the buttons below:</h3>
	</div>
	<div class="input-group mb-3">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="changeUsername">Change Username</button>
		</div>
	</div>
	<div class="input-group mb-3">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="searchUser">Search</button>
		</div>
	</div>
	<div class="input-group mb-3">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="changeUsername">Change Username</button>
		</div>
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
