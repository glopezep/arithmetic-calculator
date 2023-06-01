import { userService } from "../services/user-service";

export default defineEventHandler(async (event) => {
  try {
    const rawBody = await readRawBody(event);
    const body = new URLSearchParams(rawBody);
    const res = await userService.authenticate(
      body.get("email")!,
      body.get("password")!
    );

    event.context.session.user = res;

    await event.context.session.save();

    await sendRedirect(event, "/");
  } catch (error) {
    const params = new URLSearchParams((error as any).data);
    await sendRedirect(event, `/signin?${params}`);
  }
});
