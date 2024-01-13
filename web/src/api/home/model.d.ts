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

export interface ServerInfoResult {
  serverDetails: ServerDetails;
  mediaServerDetails: MediaServerDetails;
}

export interface ServerDetails {
  serviceAddress: string;
  sipServerID: string;
  sipServerDomain: string;
  sipPassword: string;
  uptime: string;
}

export interface MediaServerDetails {
  mediaServiceAddress: string;
  mediaUniqueID: string;
  rtpPort: number;
  restfulPort: number;
  rtspPort: number;
  rtmpPort: number;
  tcpSessions: number;
  udpSessions: number;
  lastHeartbeatTime: string;
}
