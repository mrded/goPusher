package cfg

import "gopkg.in/ini.v1"
import log "gopkg.in/inconshreveable/log15.v2"

type Options struct {
  Port    string
  Token   string
}

var options Options

func GetOptions() Options {
  return options
}

func init() {
  cfg, err := ini.Load("./config.ini")

  if err != nil {
    log.Error("Cannot read ./config.ini;", "err", err)
  }

  options.Port = cfg.Section("").Key("port").String()
  options.Token = cfg.Section("").Key("token").String()
}
