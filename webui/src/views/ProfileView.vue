<script>
export default {
	components: {},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			profile : localStorage.getItem("profile") || "",
		}
	},
	methods: {

	}
}
</script>


<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>Welcome {{profile.username}}!</h1>
	</div>
	<div>
		<h3>Please select an option from the buttons below:</h3>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
