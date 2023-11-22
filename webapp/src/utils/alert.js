export default {
  alertConfirm(title, msg, callback) {
    window.$messager.confirm({
      title: title,
      msg: msg,
      result: callback
    });
  },

  alertPrompt(title, msg, callback) {
    window.$messager.prompt({
      title: title,
      msg: msg,
      result: callback
    })
  },

  alertSuccess(msg) {
    window.$messager.alert({
      title: "提示",
      msg: msg,
      icon: "success"
    })
  },

  alertError(msg) {
    window.$messager.alert({
      title: "错误",
      msg: msg,
      icon: "error"
    })
  },

  alertInfo(msg) {
    window.$messager.alert({
      title: "提示",
      msg: msg,
      icon: "info"
    })
  },

  alertWarning(msg) {
    window.$messager.alert({
      title: "提示",
      msg: msg,
      icon: "warning"
    })
  }
}
