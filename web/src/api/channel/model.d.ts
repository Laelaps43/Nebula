export interface ReqParams {
  limit: number;
  page: number;
  deviceId: string;
}

export interface CreateChannelParams {
  name: string;
  channelId: string;
  deviceId: string;
}

export interface UpdateChannelParams {
  name: string;
  channelId: string;
}

export interface VideoRequestPayload {
  channelId: string;
  deviceId: string;
}

export interface VideoResponsePayload {
  HTTP: string;
  RTSP: string;
  RTMP: string;
  WSFLV: string;
}

export interface ResResult {
  list: {
    channelId: string;
    name: string;
    manufacturer: string;
    transport: number;
    address: string;
    status: string;
  }[];
  total: number;
}
