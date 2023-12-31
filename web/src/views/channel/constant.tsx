import { ColumnProps } from 'ant-design-vue/es/table';

export const columns: ColumnProps[] = [
  {
    title: '设备名称',
    dataIndex: 'name',
    width: 80,
  },
  {
    title: '通道编号',
    dataIndex: 'channelId',
    width: 200,
    // customRender: ({ text }) => <Tag>{text}</Tag>,
  },
  {
    title: '地址',
    dataIndex: 'address',
    width: 100,
  },
  {
    title: '厂家',
    dataIndex: 'manufacturer',
    width: 80,
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    width: 50,
  },
  {
    title: '传输模式',
    dataIndex: 'transport',
    width: 50,
  },
  {
    title: '操作',
    key: 'action',
    width: 180,
    // slots: { customRender: 'action' }, // 该用法已废弃
  },
];

export const ChannelStatusList = [
  {
    status: 'OFF',
    label: '离线',
    color: 'error',
  },
  {
    status: 'ON',
    label: '在线',
    color: 'success',
  },
];
