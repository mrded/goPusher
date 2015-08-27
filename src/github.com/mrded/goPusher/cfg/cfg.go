package cfg

import (
  "gopkg.in/ini.v1"
  "path/filepath"

  "os"
  "log"
  "fmt"
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
  // Get root dir.
  rootDir, err := filepath.Abs(fmt.Sprintf("%s/../", filepath.Dir(os.Args[0])))

  if err != nil {
    log.Fatal("Cannot get rootDir")
  }

  // Get config path.
  configPath := fmt.Sprintf("%s/config.ini", rootDir)
  cfg, err := ini.Load(configPath)

  if err != nil {
    log.Fatal("Cannot read ", configPath, err)
  }

  options.Port = cfg.Section("").Key("port").String()
  options.Token = cfg.Section("").Key("token").String()
}
