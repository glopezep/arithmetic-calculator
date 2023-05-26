interface AuthenticateResponse {
  token: string;
}

interface CreateUserRequest {
  email: string;
  password: string;
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
}

export const userService = new UserService();
