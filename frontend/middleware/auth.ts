export default defineNuxtRouteMiddleware((to, from) => {
  if (process.server) {
    const { ssrContext } = useNuxtApp();
    const session = ssrContext?.event?.context?.session;

    if (!session?.user) {
      return navigateTo("/signin");
    }
  }
});
