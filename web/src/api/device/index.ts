import { CreateDeviceParams, ReqParams, ResResult, UpdateParams } from './model';
import { post, get } from '/@/utils/http';

enum URL {
  device_page = '/v1/device/list',
  device_update = '/v1/device/update',
  device_generate = `/v1/device/create/generate`,
  device_create = `/v1/device/create/create`,
  device_delete = `/v1/device/delete`,
}

const device_page = async (data: ReqParams) => post<ResResult>({ url: URL.device_page, data });

const device_update = async (data: UpdateParams) => post<any>({ url: URL.device_update, data });

const device_generate = async () => get<string>({ url: URL.device_generate });

const device_create = async (data: CreateDeviceParams) => post({ url: URL.device_create, data });

const device_delete = async (deviceId: string) =>
  get<string>({ url: URL.device_delete + `/${deviceId}` });

export default { device_page, device_update, device_generate, device_create, device_delete };
