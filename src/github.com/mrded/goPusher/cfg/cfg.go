package cfg

import (
  "gopkg.in/ini.v1"
  "log"
)

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
    log.Fatal("Cannot read ./config.ini; %s", err)
  }

  options.Port = cfg.Section("").Key("port").String()
  options.Token = cfg.Section("").Key("token").String()
}
