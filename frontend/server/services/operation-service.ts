interface Operation {
  id: string;
  type: string;
  cost: number;
}

interface ExecuteOperationRequest {
  id: string;
  firstValue: number;
  secondValue: number;
}

interface ListOperationResponse {
  items: Operation[];
}

class OperationService {
  baseUrl: string;

  constructor() {
    this.baseUrl = "http://localhost:3000";
  }

  async executeOperation(
    operation: ExecuteOperationRequest,
    context: { authorization: string }
  ): Promise<any> {
    const res = await $fetch("/operations", {
      baseURL: this.baseUrl,
      method: "POST",
      body: operation,
      headers: {
        authorization: context.authorization,
      },
    });

    return res as any;
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

export const operationService = new OperationService();
