import { ColumnProps } from 'ant-design-vue/es/table';

export const columns: ColumnProps[] = [
  {
    title: '角色编号',
    key: 'slug',
    width: '20%',
  },
  {
    title: '角色名称',
    key: 'name',
    width: '30%',
  },
  {
    title: '操作',
    key: 'action',
    width: '50%',
  },
];
