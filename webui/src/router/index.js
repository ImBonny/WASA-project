import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SessionView from "../views/SessionView.vue";
import changeUsernameView from "../views/ChangeUsernameView.vue";
import searchView from "../views/SearchView.vue";
import profileView from "../views/ProfileView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/session', component: SessionView},
		{path: `/users/:username`, component: changeUsernameView},
		{path: '/some/:id/link', component: HomeView},
		{path: '/users', component: searchView},
		{path: `/users/:username/profile`, component: profileView}
	]
})

export default router
