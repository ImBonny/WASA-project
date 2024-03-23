<script>
import {RouterLink} from "vue-router";

export default {
	components: {RouterLink},
	data: function () {
		return {
			errormsg: null,
			username: localStorage.getItem("username") || "",
			id: localStorage.getItem("id"),
			selectedFile: null,
			caption: ""
		}
	},
	methods: {
		async selectFile() {
			// Crea un elemento input di tipo file
			let input = document.createElement("input");
			input.type = "file";
			input.accept = "image/*";

			// Aspetta che venga selezionato un file prima di andare avanti
			let selectedFilePromise = new Promise((resolve) => {
				input.onchange = () => resolve(input.files[0]);
			});

			// Simula il clic sull'elemento input per aprire la finestra di dialogo di selezione del file
			input.click();

			// Attendi che l'utente selezioni un file
			this.selectedFile = await selectedFilePromise;

			// Leggi il contenuto del file come array di byte
			let reader = new FileReader();
			reader.readAsArrayBuffer(this.selectedFile);

			// Attendi che il reader completi la lettura del file
			let fileDataPromise = new Promise((resolve) => {
				reader.onload = () => resolve(reader.result);
			});

			// Ottieni i dati binari del file
			let fileData = await fileDataPromise;

			//Aggiungi la caption
			this.caption = prompt("Inserisci una caption per la foto");
			if (this.caption == null) {
				this.caption = "";
			}

			// Invia i dati binari del file al backend
			try {
				console.log("Uploading file: " + this.selectedFile.name);
				let response = await this.$axios.post(`/users/${this.username}/posts`, {
					image: Array.from(new Uint8Array(fileData)),  // Converti i dati del file in un array di byte
					caption: this.caption
				}, {
					headers: {
						Authorization: "Bearer " + this.id
					}
				});
				console.log("File uploaded: " + response.data);
			} catch (error) {
				console.log("Error uploading file: " + error.response.data);
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
		<h1>{{username}}, let's upload a new Photo!</h1>
	</div>

	<div class="input-group mb-3">
		<div class="input-group-append">
			<button class="btn btn-success" type="button" @click="selectFile">Select a File</button>
		</div>
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
					<RouterLink :to="'/'" class="nav-link">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
						Logout
					</RouterLink>
				</li>
			</ul>
		</div>
	</nav>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style scoped>

</style>

<style scoped>

</style>
