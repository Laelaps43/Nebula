export interface ReqParams {
  limit: number;
  page: number;
  deviceId: string;
}

export interface CreateDeviceParams {
  name: string;
  deviceId: string;
  port: number;
}

export interface UpdateParams {
  name: string;
  deviceId: string;
}

export interface ResResult {
  list: {
    deviceId: number;
    realm: string;
    name: string;
    manufacturer: string;
    transport: number;
    registerAt: string;
    keepLiveAt: string;
    channelCount: string;
    status: string;
  }[];
  total: number;
}
