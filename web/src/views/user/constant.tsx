// import { Tag } from 'ant-design-vue';
import { ColumnProps } from 'ant-design-vue/es/table';

export const columns: ColumnProps[] = [
  {
    title: '用户名',
    dataIndex: 'username',
    width: 80,
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    width: 150,
  },
  {
    title: '角色',
    dataIndex: 'role',
    key: 'role',
    width: 150,
  },
  {
    title: '是否被冻结',
    dataIndex: 'enable',
    key: 'status',
    width: 80,
  },
  {
    title: '登录地址',
    dataIndex: 'lastLoginAddress',
    width: 100,
  },
  {
    title: '操作',
    key: 'action',
    width: 180,
  },
];

export const enableStatusList = [
  {
    status: '0',
    label: '冻结',
    color: 'error',
  },
  {
    status: '1',
    label: '正常',
    color: 'success',
  },
];
