interface AuthenticateResponse {
  token: string;
}

interface GetUserInfoResponse {
  id: string;
  email: string;
  status: string;
  balance: number;
}

class UserService {
  baseUrl: string;

  constructor() {
    this.baseUrl = "http://localhost:3000";
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
    const res = await $fetch("/user_info", {
      baseURL: this.baseUrl,
      method: "GET",
      headers: {
        authorization: context.authorization,
      },
    });

    return res as GetUserInfoResponse;
  }
}

export const userService = new UserService();
