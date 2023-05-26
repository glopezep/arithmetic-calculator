import type { IronSession } from "iron-session";

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
