export interface RoleCreate {
  name: string;
  slug: string;
  desc: string;
  parentId: number;
}

export interface RoleUpdate {
  id: number;
  name: string;
  slug: string;
  desc: string;
  parentId: number;
}

export interface PageParam {
  limit: number;
  page: number;
}

export interface PageResult {
  list: {
    id: string;
    name: string;
    slug: string;
    desc: number;
  }[];
  total: number;
}
