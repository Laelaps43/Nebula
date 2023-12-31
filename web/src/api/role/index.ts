import { get, post } from '/@/utils/http';
import { PageParam, PageResult, RoleCreate, RoleUpdate } from "/@/api/role/model";

enum URL {
  create_role = '/v1/role/create',
  list_role = '/v1/role/list',
  update_role = '/v1/role/update',
  delete_role = '/v1/role/delete',
}

const CreateRole = async (data: RoleCreate) => post<String>({ url: URL.create_role, data });

const ListRole = async (data: PageParam) => post<PageResult>({ url: URL.list_role, data });

const UpdateRole = async (data: RoleUpdate) => post<String>({ url: URL.update_role, data });

const DeleteRole = async (roleId: string) => get<String>({ url: URL.delete_role + `/${roleId} ` });

export default { CreateRole, ListRole, UpdateRole, DeleteRole };
