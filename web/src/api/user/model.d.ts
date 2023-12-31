export interface ReqParams {
  mobile: string;
  password: string;
}

export interface ReqAuth {
  auths: string[];
  modules: string[];
  is_admin?: 0 | 1;
}

export interface ResResult {
  login_status: number;
  st: string;
  token: string;
}

export interface PageParam {
  limit: number;
  page: number;
}

export interface PageResult {
  username: string;
  email: string;
  enable: number;
}

export interface CreateUser {
  name: string;
  email: string;
  password: string;
  passwordAgain: string;
  roles: number[];
}

export interface UpdateUser {
  id: string;
  name: string;
  email: string;
  roles: number[];
}

export interface EnableUser {
  id: string;
  enable: number;
}
