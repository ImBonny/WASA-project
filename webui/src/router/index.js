import {createRouter, createWebHashHistory} from 'vue-router'
import searchView from "../views/SearchView.vue";
import profileView from "../views/ProfileView.vue";
import myStreamView from "../views/MyStreamView.vue";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/session'},
		{path: '/session', component: HomeView},
		{path: '/users', component: searchView},
		{path: `/users/:username/profile`, component: profileView},
		{path: `/users/:username/stream`, component: myStreamView},
	]
})

export default router
