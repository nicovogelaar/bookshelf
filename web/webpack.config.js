const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const extractSass = new ExtractTextPlugin({
    filename: "[name].css",
    disable: process.env.NODE_ENV === "development"
});

module.exports = {
  entry: {
    app: [ './src/main.js' ],
    vendor: [
      'bootstrap-loader',
      'jquery',
      'vue',
      'vue-router',
      'vuex',
      'popper.js',
    ]
  },
  output: {
    path: path.join(__dirname, 'dist'),
    filename: '[name].js',
    publicPath: '/',
  },
  module: {
    loaders: [
      {
        loader: "babel-loader",
        // Skip any files outside of your project's `src` directory
        include: [
          path.resolve(__dirname, 'src'),
        ],
        exclude: [
          path.resolve(__dirname, 'node_modules'),
        ],
        // Only run `.js` and `.jsx` files through Babel
        test: /\.js$/,
        // Options to configure babel with
        query: {
          plugins: ['transform-runtime'],
          presets: ['es2015', 'stage-0'],
        }
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
      },
      {
        test: /\.scss$/,
        use: extractSass.extract({
          fallback: "style-loader",
          use: [
            { loader: "css-loader" },
            { loader: "sass-loader"}
          ],
        })
      },
    ]
  },
  devServer: {
    port: 3000,
    historyApiFallback: true
  },
  plugins: [
    extractSass,
    new HtmlWebpackPlugin({
      template: './src/assets/index.html',
      chunksSortMode: function (chunk1, chunk2) {
        var orders = ['vendor', 'app'];
        var order1 = orders.indexOf(chunk1.names[0]);
        var order2 = orders.indexOf(chunk2.names[0]);

        return order1 - order2;
      }
    }),
    new webpack.ProvidePlugin({
      $: 'jquery',
      jquery: 'jquery',
      'window.jQuery': 'jquery',
      jQuery: 'jquery',
      Popper: ['popper.js', 'default'],
      '_': 'lodash',
      'Tether': 'tether',
      utils: 'utils',
      // Alert: 'exports-loader?Alert!bootstrap/js/dist/alert',
      Button: 'exports-loader?Button!bootstrap/js/dist/button',
      // Carousel: 'exports-loader?Carousel!bootstrap/js/dist/carousel',
      Collapse: 'exports-loader?Collapse!bootstrap/js/dist/collapse',
      // Dropdown: 'exports-loader?Dropdown!bootstrap/js/dist/dropdown',
      // Modal: 'exports-loader?Modal!bootstrap/js/dist/modal',
      // Popover: 'exports-loader?Popover!bootstrap/js/dist/popover',
      // Scrollspy: 'exports-loader?Scrollspy!bootstrap/js/dist/scrollspy',
      // Tab: 'exports-loader?Tab!bootstrap/js/dist/tab',
      Tooltip: "exports-loader?Tooltip!bootstrap/js/dist/tooltip",
      Util: 'exports-loader?Util!bootstrap/js/dist/util'
    }),
  ],
};