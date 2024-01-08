export interface Record {
  channelName: String;
  channelId: String;
  deviceName: String;
  isRecording: Boolean;
  lastRecordTime: String;
  duration: Number;
}

export interface PageParams {
  limit: number;
  page: number;
}

export interface RangeParams {
  start: String;
  end: String;
  stream: String;
}

export interface SelectRecord {
  label: String;
  value: Number;
}
