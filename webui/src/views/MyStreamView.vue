<script>
import {RouterLink} from "vue-router";

export default {
	components: {RouterLink},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
		}
	},
	methods: {
	}
}
</script>


<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>{{username}}'s Stream</h1>
	</div>
	<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		<div class="position-sticky pt-3 sidebar-sticky">
			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>General</span>
			</h6>
			<ul class="nav flex-column">
				<li class="nav-item">
					<router-link to="/users/:username/stream" class="nav-link" @click="home">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
						Home
					</router-link>
				</li>
				<li class="nav-item">
					<RouterLink to="/link1" class="nav-link">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
						Menu item 1
					</RouterLink>
				</li>
				<li class="nav-item">
					<RouterLink to="/link2" class="nav-link">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
						Menu item 2
					</RouterLink>
				</li>
			</ul>

			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
				<span>Secondary menu</span>
			</h6>
			<ul class="nav flex-column">
				<li class="nav-item">
					<RouterLink :to="'/some/' + 'variable_here' + '/path'" class="nav-link">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
						Item 1
					</RouterLink>
				</li>
			</ul>
		</div>
	</nav>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>
