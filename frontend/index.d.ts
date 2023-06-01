import type { IronSession } from "iron-session";

export interface RecordItem {
  id: string;
  amount: number;
  userBalance: number;
  operationResponse: string;
  createdAt: string;
}

declare module "h3" {
  interface H3EventContext {
    session: IronSession;
  }
}

declare module "iron-session" {
  interface IronSessionData {
    user?: {
      token: string;
    };
  }
}

declare module "nuxt/schema" {
  interface RuntimeConfig {
    apiSecret: string;
    public: {
      apiBase: string;
    };
  }
}

export {};
