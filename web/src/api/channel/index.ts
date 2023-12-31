import {
  CreateChannelParams,
  ReqParams,
  ResResult,
  UpdateChannelParams,
  VideoRequestPayload, VideoResponsePayload
} from "./model";
import { get, post } from '/@/utils/http';

enum URL {
  channel_page = '/v1/channel/list',
  channel_generate = `/v1/channel/generate`,
  channel_create = `/v1/channel/create`,
  channel_update = `/v1/channel/update`,
  channel_delete = `/v1/channel/delete`,
  channel_play = `/v1/video/play`,
}

const channel_page = async (data: ReqParams) => post<ResResult>({ url: URL.channel_page, data });

const channel_generate = async () => get<string>({ url: URL.channel_generate });

const channel_create = async (data: CreateChannelParams) => post({ url: URL.channel_create, data });

const channel_update = async (data: UpdateChannelParams) => post({ url: URL.channel_update, data });

const channel_delete = async (channelId: string) =>
  get<string>({ url: URL.channel_delete + `/${channelId}` });

const video_play = async (data: VideoRequestPayload) =>
  post<VideoResponsePayload>({ url: URL.channel_play, data });

export default {
  channel_page,
  channel_generate,
  channel_create,
  channel_update,
  channel_delete,
  video_play,
};
