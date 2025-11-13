import {createRouter, createWebHistory} from 'vue-router'
import Settings from '../views/Settings.vue'
import {renderIcon} from "../utils/common";
import SettingsTwotone from "../assets/icons/SettingsTwotone.svg";

export const routes = [
    {
        path: '/settings', name: 'Settings', component: Settings, meta: {
            label: '设置',
            icon: renderIcon(SettingsTwotone),
        }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router