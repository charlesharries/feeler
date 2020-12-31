const { resolve } = require('path');

const js = {
  test: /\.js$/,
  exclude: /node_modules/,
  use: {
      loader: 'babel-loader',
      options: {
          presets: [
            ['@babel/preset-env', { targets: { node: 10 }}],
            '@babel/preset-react',
          ],
      }
  }
}

module.exports = {
  mode: process.env.NODE_ENV === 'production' ? 'production' : 'development',
  entry: './assets/js/index.js',
  output: {
    path: resolve(__dirname, 'static'),
    filename: 'app.js',
  },
  module: { rules: [js] }
}