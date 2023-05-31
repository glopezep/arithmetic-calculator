import { userService } from "../services/user-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;

  return userService.me({
    authorization: `Bearer ${token}`,
  });
});
