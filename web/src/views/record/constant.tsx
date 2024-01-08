// import { Tag } from 'ant-design-vue';
import { ColumnProps } from 'ant-design-vue/es/table';

export const columns: ColumnProps[] = [
  {
    title: '通道名称',
    dataIndex: 'channelName',
    width: 80,
  },
  {
    title: '通道编号',
    dataIndex: 'channelId',
    width: 120,
  },
  {
    title: '设备名称',
    dataIndex: 'deviceName',
    width: 100,
  },
  {
    title: '录制中',
    dataIndex: 'isRecording',
    key: 'status',
    width: 80,
  },
  {
    title: '最近录制时间',
    dataIndex: 'lastRecordTime',
    width: 120,
  },
  {
    title: '总时长',
    dataIndex: 'duration',
    width: 80,
  },
  {
    title: '操作',
    key: 'action',
    width: 100,
  },
];

export const RecordStatusList = [
  {
    status: 1,
    label: '否',
    color: 'error',
  },
  {
    status: 2,
    label: '是',
    color: 'success',
  },
];
