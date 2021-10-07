package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	HttpConfig  struct{
		Port int `yaml:"port"`
	} 	`yaml:"http"`
	LogConfig struct{
		Output string`yaml:"output"`
	} `yaml:"log"`
}

func NewConfig()*Config  {
	data,err:=ioutil.ReadFile("./config.yaml")
	if err!=nil{
		panic(err)
	}
	c:=&Config{}
	err=yaml.Unmarshal(data,c)
	if err!=nil{
		panic(err)
	}
	return c
}
