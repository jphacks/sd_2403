const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  pages: {
    index: {
      entry: 'src/main.js',
      // template title tag needs to be <title><%= htmlWebpackPlugin.options.title %></title>
      title: 'フードチェーン ~みんなで食安全を守ろう~',
    },
  },
  publicPath: process.env.NODE_ENV === "production" 
    ? "/ProductTraceability-TMA/" 
    : "/",
  
})
