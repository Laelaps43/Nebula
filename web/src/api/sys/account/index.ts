/**
 * @name account
 * @description 系统设置 - 账户模块
 */

import { ReqAccount, ResAccount } from './model';
import { get, post } from '/@/utils/http';

enum URL {
  update = '/v1/user/my/edit',
  account = '/v1/user/my/info',
}

const account = async () => get<ResAccount>({ url: URL.account });

const update = async (data: ReqAccount) => post<any>({ url: URL.update, data });

export default { account, update };
