import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import changeUsernameView from "../views/ChangeUsernameView.vue";
import searchView from "../views/SearchView.vue";
import profileView from "../views/ProfileView.vue";
import myStreamView from "../views/MyStreamView.vue";
import uploadView from "../views/UploadView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: `/users/:username`, component: changeUsernameView},
		{path: '/some/:id/link', component: HomeView},
		{path: '/users', component: searchView},
		{path: `/users/:username/profile`, component: profileView},
		{path: `/users/:username/stream`, component: myStreamView},
		{path: `/users/:username/posts`, component: uploadView},
	]
})

export default router
