export interface OverViewResult {
  onlineDevice: string;
  offlineDevice: string;
  channel: string;
  video: string;
}

export interface SystemInfoResult {
  CPUList: [];
  MemList: [];
  DiskList: [];
  UpList: [];
  DownList: [];
  TimeList: [];
}
