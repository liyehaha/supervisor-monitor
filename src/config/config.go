package config

import (
	"strings"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	NotifyType		string			`yaml:"type"`	
	Mail			MailConfig		`yaml:"mail"`
}

type MailConfig struct {
	User     	string		`yaml:"user"`
	Password 	string		`yaml:"password"`
	Host     	string		`yaml:"smtp"`
	Port     	int			`yaml:"port"`
	Repcients 	string		`yaml:"repcients"`
}

var config Config

func SplitRepcients(repcients string) []string {
	if strings.Contains(repcients, ",") {
		return strings.Split(repcients, ",")
	}
	return append(make([]string, 1), repcients)
}

func Load(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal(f, &config)
		return &config, nil
	}else{
		return nil, err
	}
}