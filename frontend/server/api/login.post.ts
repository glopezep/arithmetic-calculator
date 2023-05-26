import { userService } from "../services/user-service";

export default defineEventHandler(async (event) => {
  const rawBody = await readRawBody(event);
  const body = new URLSearchParams(rawBody);
  const res = await userService.authenticate(
    body.get("email")!,
    body.get("password")!
  );

  event.context.session.user = res;

  await event.context.session.save();

  await sendRedirect(event, "/");
});
