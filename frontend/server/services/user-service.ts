interface AuthenticateResponse {
  token: string;
}

interface GetUserInfoResponse {
  id: string;
  email: string;
  status: string;
  balance: number;
}

export class UserService {
  baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  async authenticate(
    email: string,
    password: string
  ): Promise<AuthenticateResponse> {
    const res = await $fetch("/login", {
      baseURL: this.baseUrl,
      method: "POST",
      body: {
        email,
        password,
      },
    });

    return res as AuthenticateResponse;
  }

  async createUser(
    email: string,
    password: string
  ): Promise<AuthenticateResponse> {
    const res = await $fetch("/users", {
      baseURL: this.baseUrl,
      method: "POST",
      body: {
        email,
        password,
      },
    });

    return res as AuthenticateResponse;
  }

  async me(context: { authorization: string }): Promise<GetUserInfoResponse> {
    const res = await $fetch("/user-info", {
      baseURL: this.baseUrl,
      method: "GET",
      headers: {
        authorization: context.authorization,
      },
    });

    return res as GetUserInfoResponse;
  }
}

export const userService = new UserService(useRuntimeConfig().baseApiHost);
