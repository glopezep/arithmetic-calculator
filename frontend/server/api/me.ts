export default defineEventHandler(async (event) => {
  return {
    token: event.context.session.user?.token,
    balance: 89.4,
  };
});
