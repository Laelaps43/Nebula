import { defineStore } from 'pinia';
import { store } from '/@/store';
import { ReqParams } from '/@/api/user/model';
import fetchApi from '/@/api/user';
// import { encryptByDES } from '/@/utils/crypto';
import { getToken, setToken, removeToken } from '/@/utils/auth';
import { usePermissioStore } from '/@/store/modules/permission';
import { useSysAccountStore } from '/@/store/modules/sysAccount';
import { useHomeStore } from '/@/store/modules/home';
import { router } from '/@/router';

interface UserState {
  token: string;
  auths: string[];
  roleSet: any[];
}

export const useUserStore = defineStore({
  id: 'app-user',
  state: (): UserState => ({
    // token
    token: '',
    // auths
    auths: [],
    roleSet: [],
  }),
  getters: {
    getToken(): string {
      return this.token || getToken();
    },
  },
  actions: {
    setToken(info: string) {
      this.token = info ?? ''; // for null or undefined value
      setToken(info);
    },
    setAuth(auths: string[]) {
      this.auths = auths;
    },
    setRoleSet(roleSet: any[]) {
      this.roleSet = roleSet;
    },
    getRoleSet(): any[] {
      return this.roleSet;
    },
    resetState() {
      this.token = '';
      this.auths = [];
      this.roleSet = [];
    },
    /**
     * @description: login
     */
    async login(params: ReqParams) {
      // 密码加密
      // TODO: 需要对密码加密
      // params.password = encryptByDES(params.password);
      const res = await fetchApi.login(params);
      if (res) {
        // save token
        this.setToken(res.token);
      }
      return res;
    },

    /**
     * @description：获取用户角色
     */
    async fetchUserRole() {
      const res = await fetchApi.user_role();
      if (res) {
        // save token
        this.setRoleSet(res);
      }
      return res;
    },
    async switchRole() {
      this.roleSet = [];
      this.auths = [];
      const permissionStore = usePermissioStore();
      const sysAccountStore = useSysAccountStore();
      const homeStore = useHomeStore();
      permissionStore.resetState();
      sysAccountStore.resetState();
      homeStore.resetState();
      router.go(0);
    },
    /**
     * @description: logout
     */
    async logout() {
      this.resetState();
      removeToken();
      // router.replace('/login');
      // 路由表重置
      location.reload();
    },
  },
});

// Need to be used outside the setup
export function useUserStoreWithOut() {
  return useUserStore(store);
}
