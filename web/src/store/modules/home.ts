import { defineStore } from 'pinia';
import { store } from '/@/store';
import fetchApi from '/@/api/home';
import { OverViewResult, ServerInfoResult } from '/@/api/home/model';

interface HomeState {
  overView: Nullable<OverViewResult>;
  systemInfo: Nullable<any>;
  serverInfo: Nullable<ServerInfoResult>;
}

export const useHomeStore = defineStore({
  id: 'app-home',
  state: (): HomeState => ({
    overView: null,
    systemInfo: null,
    serverInfo: null,
  }),
  getters: {
    getOverView(): Nullable<OverViewResult> {
      return this.overView || null;
    },
    getSystemInfo(): any {
      return this.systemInfo || null;
    },
    getServerInfo(): Nullable<ServerInfoResult> {
      return this.serverInfo;
    },
  },
  actions: {
    setOverView(overview: Nullable<OverViewResult>) {
      this.overView = overview;
    },
    setSystemInfo(systemInfo: any) {
      this.systemInfo = systemInfo;
    },
    setServerInfo(serverInfo: ServerInfoResult) {
      this.serverInfo = serverInfo;
    },
    resetState() {
      this.overView = null;
      this.systemInfo = null;
      this.serverInfo = null;
    },
    /**
     * @description: login
     */
    async fetchInfo() {
      const res = await fetchApi.overView();
      if (res) {
        // save token
        this.setOverView(res);
      }
      return res;
    },
    async fetchSystemInfo() {
      // 最近一个小时的时间
      const res = await fetchApi.systemInfo();
      if (res) {
        // save token
        this.setSystemInfo(res);
      }
      return res;
    },
    async fetchServerInfo() {
      // 最近一个小时的时间
      const res = await fetchApi.serverInfo();
      if (res) {
        // save token
        this.setServerInfo(res);
      }
      return res;
    },
  },
});

// Need to be used outside the setup
export function useHomeStoreWithOut() {
  return useHomeStore(store);
}
