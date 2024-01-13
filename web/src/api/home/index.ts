import { get } from '/@/utils/http';
import { OverViewResult, ServerInfoResult, SystemInfoResult } from '/@/api/home/model';

enum URL {
  overView = '/v1/home/overview',
  systemInfo = '/v1/home/system/info',
  serverInfo = `/v1/home/server/info`,
}

const overView = async () => get<OverViewResult>({ url: URL.overView });

const systemInfo = async () => get<SystemInfoResult>({ url: URL.systemInfo });

const serverInfo = async () => get<ServerInfoResult>({ url: URL.serverInfo });

export default { overView, systemInfo, serverInfo };
