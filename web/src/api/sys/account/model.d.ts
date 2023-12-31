export interface ReqAccount {
  id: number;
  username?: string;
  password?: string;
}

export interface ResAccount {
  last_login: string;
  email: string;
  role: string;
  headerImg: string;
  id: number;
}

export type ResPermission = { auths: Array<string> };
