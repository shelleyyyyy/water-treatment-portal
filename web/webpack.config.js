module.exports = {
    module: {
      rules: [
        {
          test: /\.m?(j|t)sx?$/,
          // Excluding node_modules means that core-js will not be compiled
          exclude: /node_modules/,
          use: ['babel-loader']
        }
      ]
    }
  }
  