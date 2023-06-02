interface Operation {
  id: string;
  type: string;
  cost: number;
}

interface ListOperationResponse {
  items: Operation[];
}

export class OperationService {
  baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  async listOperations(context: {
    authorization: string;
  }): Promise<ListOperationResponse> {
    const res = await $fetch("/operations", {
      baseURL: this.baseUrl,
      method: "GET",
      headers: {
        authorization: context.authorization,
      },
    });

    return res as ListOperationResponse;
  }
}

export const operationService = new OperationService(
  useRuntimeConfig().baseApiHost
);
