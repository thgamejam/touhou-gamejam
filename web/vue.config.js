module.exports = {
  pluginOptions: {
    antd: {
      importType: 'full',
      style: 'css'
    }
  },
  css: {
    loaderOptions: {
      less: {
        lessOptions: {
          javascriptEnabled: true
        }
      }
    }
  }
}
