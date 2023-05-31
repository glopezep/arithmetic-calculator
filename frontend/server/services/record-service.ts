import { RecordItem } from "../..";

interface ListRecordsRequest {
  limit?: number;
  offset?: number;
  sort_by?: number;
  order_by?: number;
}

interface ListRecordsResponse {
  items: RecordItem[];
  totalCount: number;
  offset: number;
  limit: number;
  hasNextPage: boolean;
}

class RecordService {
  baseUrl: string;

  constructor() {
    this.baseUrl = "http://localhost:3000";
  }

  async listRecords(
    request: ListRecordsRequest,
    context: {
      authorization: string;
    }
  ): Promise<ListRecordsResponse> {
    const res = await $fetch("/records", {
      baseURL: this.baseUrl,
      query: request,
      method: "GET",
      headers: {
        authorization: context.authorization,
      },
    });

    return res as ListRecordsResponse;
  }
}

export const recordService = new RecordService();
