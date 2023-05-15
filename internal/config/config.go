package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/ysmood/gop"
	"gopkg.in/yaml.v2"
)

var GlobalCfg = new(AppConfig)

type AppConfig struct {
	ResourceName string            `yaml:"resourceName"`
	ApiVersion   string            `yaml:"apiVersion"`
	ApiKey       string            `yaml:"apiKey"`
	Mapper       map[string]string `yaml:"mapper"`
}

func LoadConfig(filename string) error {
	// 读取 config.yaml 文件内容
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, GlobalCfg)
	if err != nil {
		logrus.Fatalf("Error unmarshaling config file: %v", err)
	}
	// logrus.Infof("cfg:%v", *GlobalCfg)
	gop.P(*GlobalCfg)
	return nil
}
