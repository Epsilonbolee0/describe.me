const path = require('path');

const HtmlWebpackPlugin = require("html-webpack-plugin");
const ReactRefreshWebpackPlugin = 
  require("@pmmmwh/react-refresh-webpack-plugin");
const ForkTsCheckerPlugin = 
  require("fork-ts-checker-webpack-plugin");

const buildPath = path.resolve(__dirname, "dist");
const publicPath = path.resolve(__dirname, "public");
const srcPath = path.resolve(__dirname, "src");

const isProd = process.env.NODE_ENV === "production";


const plugins = [
  new HtmlWebpackPlugin({
    template: path.join(publicPath, "index.html")
  }),

  !isProd && new ReactRefreshWebpackPlugin(),
  
  new ForkTsCheckerPlugin()
].filter(Boolean);


module.exports = {
  entry: path.resolve(srcPath, "index.tsx"),
  target: process.env.NODE_ENV === "development" ? 
  "web" : "browserslist",
  output: {
    path: buildPath,
    filename: "bundle.js"
  },
  plugins,
  devServer: {
    static: publicPath,
    host: "127.0.0.1",
    port: 1488,
    hot: true
  },
  module: {
    rules: [
      {
        test: /\.([jt])sx?$/,
        use: "babel-loader"
      },
      {
        test: /\.css/,
        use: ["style-loader", "css-loader"]
      }
    ]
  }
}