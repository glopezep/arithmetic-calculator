export default defineEventHandler(async (event) => {
  console.log("=====");

  await event.context.session.destroy();

  return sendRedirect(event, "/signin");
});
