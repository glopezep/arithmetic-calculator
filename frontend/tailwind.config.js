const {
  iconsPlugin,
  getIconCollections,
} = require("@egoist/tailwindcss-icons");

module.exports = {
  plugins: [
    require("daisyui"),
    iconsPlugin({
      // Select the icon collections you want to use
      collections: getIconCollections(["mdi", "uil"]),
    }),
  ],
};
