const path = require('path');

let sassImplementation;
try {
  // tslint:disable-next-line:no-implicit-dependencies
  sassImplementation = require('node-sass');
} catch {
  sassImplementation = require('sass');
}

module.exports = {
  module: {
    rules: [
      {
        test: /\.scss$/,
        use: [
          {
            loader: 'postcss-loader',
            options: {
              postcssOptions: {
                ident: 'postcss',
                syntax: 'postcss-scss',
                plugins: ['postcss-import', 'tailwindcss']
              }
            }
          },
          {
            loader: 'sass-loader',
            options: {
              implementation: sassImplementation,
              sourceMap: false,
              sassOptions: {
                precision: 8
              }
            }
          }
        ]
      },
      {
        test: /projects[\\\/]sdk-ui[\\\/]assets[\\\/]styles[\\\/]tailwind\.scss$/,
        use: [
          {
            loader: '@fullhuman/purgecss-loader',
            options: {
              content: [
                path.join(__dirname, 'src/**/*.html'),
                path.join(__dirname, 'src/**/*.ts'),
                path.join(__dirname, 'projects/**/*.html'),
                path.join(__dirname, 'projects/**/*.ts')
              ],
              defaultExtractor: content => content.match(/[\w-/:]+(?<!:)/g) || []
            }
          },
          {
            loader: 'postcss-loader',
            options: {
              postcssOptions: {
                ident: 'postcss',
                syntax: 'postcss-scss',
                plugins: ['postcss-import', 'tailwindcss']
              }
            }
          },
          {
            loader: 'sass-loader',
            options: {
              implementation: sassImplementation,
              sourceMap: false,
              sassOptions: {
                precision: 8
              }
            }
          }
        ]
      }
    ]
  }
};
