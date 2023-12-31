import {
  CreateUser,
  EnableUser,
  PageResult,
  ReqAuth,
  ReqParams,
  ResResult,
  UpdateUser,
} from './model';
import { get, post } from '/@/utils/http';

enum URL {
  login = '/v1/user/login',
  permission = '/v1/user/permission',
  page = '/v1/user/page',
  create = '/v1/user/create',
  enable = '/v1/user/enable',
  update = `/v1/user/update`,
  delete = '/v1/user/delete',
}

const login = async (data: ReqParams) => post<ResResult>({ url: URL.login, data });

const permission = async () => get<ReqAuth>({ url: URL.permission });

const user_page = async (data: PageResult) => post<PageResult>({ url: URL.page, data });

const user_create = async (data: CreateUser) => post<string>({ url: URL.create, data });

const user_enable = async (data: EnableUser) => post<string>({ url: URL.enable, data });

const user_update = async (data: UpdateUser) => post<string>({ url: URL.update, data });

const user_delete = async (userId: string) => get<string>({ url: URL.delete + `/${userId}` });

export default { login, permission, user_page, user_create, user_enable, user_update, user_delete };
