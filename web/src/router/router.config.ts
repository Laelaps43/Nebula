import BasicLayout from '/@/layouts/BasicLayout/index.vue';
import BlankLayout from '/@/layouts/BlankLayout.vue';
import type { RouteRecordRaw } from 'vue-router';

export const accessRoutes: RouteRecordRaw[] = [
  {
    path: '/app',
    name: 'app',
    component: BasicLayout,
    redirect: '/app/home',
    meta: { title: '管理平台' },
    children: [
      {
        path: '/app/home',
        component: () => import('/@/views/home/index.vue'),
        name: 'home',
        meta: {
          title: '数据',
          icon: 'liulanqi',
          auth: ['home'],
        },
      },
      {
        path: '/app/devices',
        name: 'website',
        component: () => import('/@/views/devices/index.vue'),
        meta: {
          title: '设备列表',
          keepAlive: true,
          icon: 'jiedianguanli',
          auth: ['website'],
        },
      },
      {
        path: '/app/table-demo',
        name: 'table-demo',
        component: () => import('/@/views/table-demo/index.vue'),
        meta: {
          title: '远端录像',
          keepAlive: true,
          icon: 'rili',
        },
      },
      {
        path: '/app/record',
        name: 'table-recor',
        component: () => import('/@/views/table-demo/index.vue'),
        meta: {
          title: '用户管理',
          keepAlive: true,
          icon: 'rili',
        },
      },
      {
        path: '/app/others',
        name: 'others',
        component: BlankLayout,
        redirect: '/app/others/about',
        meta: {
          title: '系统设置',
          icon: 'shurumimadenglu',
          auth: ['others'],
        },
        children: [
          {
            path: '/app/others/about',
            name: 'about',
            component: () => import('/@/views/others/about/index.vue'),
            meta: { title: '信令服务', keepAlive: true, hiddenWrap: true },
          },
          {
            path: '/app/others/antdv',
            name: 'antdv',
            component: () => import('/@/views/others/antdv/index.vue'),
            meta: { title: '媒体服务', keepAlive: true, breadcrumb: true },
          },
          {
            path: '/app/others/login',
            name: 'login',
            component: () => import('/@/views/others/antdv/index.vue'),
            meta: { title: '系统日志', keepAlive: true, breadcrumb: true },
          },
        ],
      },
      {
        path: '/sys/account',
        name: 'account',
        component: () => import('/@/views/account/index.vue'),
        meta: { title: '用户管理', keepAlive: true, breadcrumb: true },
      },
    ],
  },
];

const constantRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    component: () => import('/@/views/login/index.vue'),
    name: 'login',
    meta: { title: '登录' },
  },
  {
    path: '/',
    name: 'Root',
    redirect: '/app',
    meta: {
      title: 'Root',
    },
  },
  // ...accessRoutes,
];

export const publicRoutes = [
  {
    path: '/redirect',
    component: BlankLayout,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('/@/views/redirect/index'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)',
    redirect: '/404',
  },
  {
    path: '/404',
    component: () => import('/@/views/404.vue'),
  },
];

// /**
//  * 基础路由
//  * @type { *[] }
//  */
// export const constantRouterMap = [];

export default constantRoutes;
