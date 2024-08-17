package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Config struct {
    Database struct {
        Driver   string `yaml:"driver"`
        Host     string `yaml:"host"`
        Port     int    `yaml:"port"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        Name     string `yaml:"name"`
    } `yaml:"database"`
}

func LoadConfig() (*Config, error) {
    data, err := ioutil.ReadFile("config/config.yaml")
    if err != nil {
        return nil, err
    }

    var cfg Config
    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        return nil, err
    }

    return &cfg, nil
}