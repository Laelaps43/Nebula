import { get } from '/@/utils/http';
import { OverViewResult, SystemInfoResult } from '/@/api/home/model';

enum URL {
  overView = '/v1/home/overview',
  systemInfo = '/v1/system/info',
}

const overView = async () => get<OverViewResult>({ url: URL.overView });

const systemInfo = async () => get<SystemInfoResult>({ url: URL.systemInfo });

export default { overView, systemInfo };
