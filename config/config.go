package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/beslow/rocketmq_consumer/logger"
	"gopkg.in/yaml.v2"
)

type Config struct {
	NameServerAddr string `yaml:"namesrv_addr"`
	Port           string
}

var c Config

func init() {
	dir, _ := os.Getwd()

	configPath := "config.yml"

	data, err := os.ReadFile(filepath.Join(dir, configPath))
	if err != nil {
		logger.Errorf("Read config.yml fail, err: %#v", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		logger.Errorf("unmarshal config.yml fail, err: %#v", err)
	}
}

func NameServerAddr() string {
	var addr strings.Builder
	if s := os.Getenv("NAME_SERVER_ADDR"); s != "" {
		addr.Write([]byte(s))
	} else {
		addr.Write([]byte(c.NameServerAddr))
	}

	addr.Write([]byte(":"))

	port := c.Port
	if port == "" {
		port = "9876"
	}
	addr.Write([]byte(port))

	return addr.String()
}
