import { userService } from "../services/user-service";

export default defineEventHandler(async (event) => {
  const rawBody = await readRawBody(event);
  const body = new URLSearchParams(rawBody);

  await userService.createUser(body.get("email")!, body.get("password")!);

  await sendRedirect(event, "/signin");
});
