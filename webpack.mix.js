let path = require("path")
let mix = require("laravel-mix");


mix.alias({
  "@": path.join(__dirname, "./client/src/")
});

mix.setPublicPath("dist")
  .js("client/src/main.jsx", "dist")
  .extract()
  .version()
  .sourceMaps()
  .react()
  .postCss('client/src/css/style.css', 'css', [
    require('@tailwindcss/postcss')
  ]);
