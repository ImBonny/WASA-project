
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

		},
		async home() {
			this.$router.push(`/users/${this.username}/stream`);
		},
		async search() {
			this.$router.push(`/users`);
		},
		async myProfile() {
			this.$router.push(`/users/${this.username}/profile`);
		},
		async logout() {
			localStorage.clear();
			this.$router.push(`/`);
		}
	}
}
</script>

<template>
	<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		<div class="position-sticky pt-3 sidebar-sticky">
			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>General</span>
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
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
						Search
					</RouterLink>
				</li>
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="myProfile">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
						My Profile
					</RouterLink>
				</li>
			</ul>

			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>Secondary menu</span>
			</h6>
			<ul class="nav flex-column">
				<li class="nav-item">
					<RouterLink to="" class="nav-link" @click="logout">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
						Item 1
					</RouterLink>
				</li>
			</ul>
		</div>
	</nav>
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
