const path = require('path');

const HtmlWebpackPlugin = require("html-webpack-plugin");
const ReactRefreshWebpackPlugin = 
  require("@pmmmwh/react-refresh-webpack-plugin");
const ForkTsCheckerPlugin = 
  require("fork-ts-checker-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const buildPath = path.resolve(__dirname, "dist");
const publicPath = path.resolve(__dirname, "public");
const srcPath = path.resolve(__dirname, "src");

const isProd = process.env.NODE_ENV === "production";


const plugins = [
  new HtmlWebpackPlugin({
    template: path.join(publicPath, "index.html")
  }),

  !isProd && new ReactRefreshWebpackPlugin(),
  
  new ForkTsCheckerPlugin(),
  new MiniCssExtractPlugin({
    filename: '[name]-[hash].css'
  })
].filter(Boolean);


const getSettingsForStyles = (withModules = false) => {
  return [
    MiniCssExtractPlugin.loader, 
    !withModules ? "css-loader" : {
      loader: "css-loader",
      options: {
        modules: {
          localIdentName: !isProd ? 
            '[path][name]__[local]' : 
            '[hash][base64]'
        }
      },
    },
    {
      loader: "postcss-loader",
      options: {
        postcssOptions: {
          plugins: ["autoprefixer"]
        }
      }
    },
    "sass-loader"
  ];
}


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
  resolve: {
    extensions: [".jsx", ".js", ".tsx", ".ts"],
    alias: {
      components: path.join(srcPath, "components"),
      styles: path.join(srcPath, "styles"),
      pages: path.join(srcPath, "app/pages")
    }
  },
  module: {
    rules: [
      {
        test: /\.([jt])sx?$/,
        use: "babel-loader"
      },
      {
        test: /\.module\.s?css$/,
        use: getSettingsForStyles(true)
      },
      {
        test: /\.s?css$/,
        exclude: /\.module\.s?css$/,
        use: getSettingsForStyles()
      },
      {
        test: /\.(png|svg|jpg)$/,
        type: "asset/resource"
      }
    ]
  }
}