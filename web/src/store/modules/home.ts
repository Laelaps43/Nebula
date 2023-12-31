import { defineStore } from 'pinia';
import { store } from '/@/store';
import fetchApi from '/@/api/home';
import { OverViewResult } from '/@/api/home/model';

interface HomeState {
  overView: Nullable<OverViewResult>;
  systemInfo: Nullable<any>;
}

export const useHomeStore = defineStore({
  id: 'app-home',
  state: (): HomeState => ({
    overView: null,
    systemInfo: null,
  }),
  getters: {
    getOverView(): Nullable<OverViewResult> {
      return this.overView || null;
    },
    getSystemInfo(): any {
      return this.systemInfo || null;
    },
  },
  actions: {
    setOverView(overview: Nullable<OverViewResult>) {
      this.overView = overview;
    },
    setSystemInfo(systemInfo: any) {
      this.systemInfo = systemInfo;
    },
    resetState() {
      this.overView = null;
      this.systemInfo = null;
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
  },
});

// Need to be used outside the setup
export function useHomeStoreWithOut() {
  return useHomeStore(store);
}
