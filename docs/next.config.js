const withNextra = require('nextra')({
  theme: 'nextra-theme-docs',
  themeConfig: 'theme.config.jsx',
})

if (process.env.TARGET === 'static') {
  module.exports = withNextra({
    basePath: process.env.BASE_PATH,

    output: 'export',
    distDir: 'dist',

    images: {
      unoptimized: true,
    },
  })
  return
}

module.exports = withNextra({})
