import { Tag } from 'ant-design-vue';
import { ColumnProps } from 'ant-design-vue/es/table';

export const columns: ColumnProps[] = [
  {
    title: '设备名称',
    dataIndex: 'name',
    width: 80,
  },
  {
    title: '设备编号',
    dataIndex: 'id',
    width: 200,
    customRender: ({ text }) => <Tag>{text}</Tag>,
  },
  {
    title: '地址',
    dataIndex: 'ip',
    width: 100,
  },
  {
    title: '厂家',
    dataIndex: 'factory',
    width: 80,
  },
  {
    title: '通道',
    dataIndex: 'channel',
    width: 50,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 50,
  },
  {
    title: '传输模式',
    dataIndex: 'protocol',
    width: 50,
    // customRender: ({ text }) => (
    //   <a href={text} target="_blank">
    //     {text}
    //   </a>
    // ),
  },
  {
    title: '最近心跳时间',
    dataIndex: 'heartbeat',
    width: 120,
  },
  {
    title: '最近注册时间',
    dataIndex: 'register',
    width: 120,
  },
  {
    title: '操作',
    key: 'action',
    width: 180,
    // slots: { customRender: 'action' }, // 该用法已废弃
  },
];
