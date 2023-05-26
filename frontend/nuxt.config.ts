export default defineNuxtConfig({
  modules: ["@nuxtjs/tailwindcss", "@nuxtjs/color-mode", "nuxt-iron-session"],
  colorMode: {
    preference: "system", // default theme
    dataValue: "theme", // activate data-theme in <html> tag
    classSuffix: "",
  },
  session: {
    cookieName: "yourapp_cookiename",
    password: "complex_password_at_least_32_characters_long",
    cookieOptions: {
      secure: process.env.NODE_ENV === "production",
    },
  },
});
