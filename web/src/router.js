import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login'
import Regist from './views/Regist'
import PageNoteFound from './views/PageNotFound'
import NavMenu from './components/NavMenu'
import Home from './views/Home'
import Read from './views/Read'
import DevTest from './views/DevTest'
import Profile from './views/Profile'
import Write from './views/Write'
import {CheckAuthAPI} from "@/common/api";

Vue.use(Router)

const router = new Router({
    mode: 'history',
    routes: [
        {path: '/devtest', name: 'devtest', component: DevTest},
        {path: '/notfound', name: 'notfound', component: PageNoteFound},

        {path: '/login', name: 'login', component: Login},
        {path: '/regist', name: 'regist', component: Regist},
        {path: '/', redirect: "/home", component: NavMenu, children: [
                {path: '/home', name: 'home', component: Home},
                {path: '/profile', name: 'profile', component: Profile, meta:{requireAuth: true}},
                {path: '/write', name: 'write', component: Write, meta:{requireAuth: true}},
                {path: '/read', name: 'read', component: Read},
            ]
        },
        {path:'*',redirect:'/notfound'},
    ]
})

router.beforeEach((to, from, next) => {
    if (to.meta.requireAuth === true){
        CheckAuthAPI().then(resp=>{
            if (resp.ret === 0) {
                next()
            } else {
                next({ path: '/login' })
            }
        }).catch(() => {
            next({path: '/'})
        })
    } else {
        next()
    }
})

export default router