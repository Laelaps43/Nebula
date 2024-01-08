import { post, get } from '/@/utils/http';
import { PageParams, RangeParams, Record, SelectRecord } from '/@/api/record/model';

enum URL {
  record_page = '/v1/record/page',
  record_range = '/v1/record/range',
  record_select = '/v1/record/select',
  record_play = '/v1/record/play',
}

const recordPage = async (data: PageParams) => post<Record>({ url: URL.record_page, data });

const recordRange = async (data: RangeParams) => post<String[]>({ url: URL.record_range, data });

const recordSelect = async (data: string, stream: string) =>
  get<SelectRecord[]>({ url: URL.record_select + `/${stream}/${data}` });

const recordPlay = async (id: number, stream: string) =>
  get<string>({ url: URL.record_play + `/${stream}/${id}` });

export default { recordPage, recordRange, recordSelect, recordPlay };
