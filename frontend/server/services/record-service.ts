import { RecordItem } from "../..";
import { Context } from "./context";

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

export class RecordService {
  baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  async listRecords(
    request: ListRecordsRequest,
    context: Context
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

  async deleteRecord(id: string, context: Context) {
    return await $fetch(`/records/${id}`, {
      baseURL: this.baseUrl,
      method: "DELETE",
      headers: {
        authorization: context.authorization,
      },
    });
  }
}

export const recordService = new RecordService(useRuntimeConfig().baseApiHost);
