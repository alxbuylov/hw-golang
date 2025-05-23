package main

import (
	"flag"
	"fmt"

	"github.com/alxbuylov/hw-golang/hw12_13_14_15_calendar/pkg/config"
)

type Config struct {
	App struct {
		Storage string `mapstructure:"storage"`
		Server  string `mapstructure:"server"`
	} `mapstructure:"app"`

	Logger struct {
		Level     string `mapstructure:"level"`
		Handler   string `mapstructure:"handler"`
		Filename  string `mapstructure:"filename"`
		AddSource bool   `mapstructure:"add_source"`
	} `mapstructure:"logger"`

	Server struct {
		HTTPAddr string `mapstructure:"http_addr"`
		GRPCAddr string `mapstructure:"grpc_addr"`
	} `mapstructure:"server"`

	Database struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"database"`

	DatabaseDNS string
}

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
}

func NewConfig() (*Config, error) {
	instance := &Config{}
	if err := config.LoadConfig(instance, configFile); err != nil {
		return nil, err
	}

	instance.DatabaseDNS = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		instance.Database.Username,
		instance.Database.Password,
		instance.Database.Dbname,
		instance.Database.Host,
		instance.Database.Port,
	)

	return instance, nil
}
