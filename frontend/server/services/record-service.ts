interface Record {
  id: string;
  amount: number;
  userBalance: number;
  operationResponse: string;
}

interface ListOperationResponse {
  items: Record[];
}

class RecordService {
  baseUrl: string;

  constructor() {
    this.baseUrl = "http://localhost:3000";
  }

  async listRecords(context: {
    authorization: string;
  }): Promise<ListOperationResponse> {
    const res = await $fetch("/records", {
      baseURL: this.baseUrl,
      method: "GET",
      headers: {
        authorization: context.authorization,
      },
    });

    return res as ListOperationResponse;
  }
}

export const recordService = new RecordService();
