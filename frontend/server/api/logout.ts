export default defineEventHandler(async (event) => {
  await event.context.session.destroy();

  return sendRedirect(event, "/signin");
});
