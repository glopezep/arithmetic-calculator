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
  runtimeConfig: {
    baseApiHost: process.env.API_HOST || "http://localhost:3000",
  },
});

// "https://l5b2uptpi5.execute-api.us-east-1.amazonaws.com/v1",
