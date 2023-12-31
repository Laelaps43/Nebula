import { ReqParams, ResResult } from './model';
import { get, post } from '/@/utils/http';

enum URL {
  page_one_list = '/v1/device/page_one/list',
  list = '/v1/node/nodelist',
  page_channel_list = '/v1/channel/page_one/list',
}

const page_one_list = async (data: ReqParams) => post<ResResult>({ url: URL.page_one_list, data });

const node_list = async (data: ReqParams) => get<ResResult>({ url: URL.list, data });

const channel_page_one_list = async (data: ReqParams) =>
  post<ResResult>({ url: URL.page_channel_list, data });

export default { page_one_list, node_list, channel_page_one_list };
