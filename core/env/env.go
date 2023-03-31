package env

import (
	"flag"
	"os"
)

const (
	dev  = "dev"
	test = "test"
	prod = "prod"
)

var (
	// environment 当前系统环境
	environment string

	pid         = os.Getpid()
	hostname, _ = os.Hostname()
)

func Init() (err error) {
	flag.CommandLine.StringVar(&environment, "e", "", "set env, e.g dev test prod")
	err = flag.CommandLine.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	// 从环境变量中获取
	if len(environment) == 0 {
		environment = os.Getenv("ENVIRONMENT")
	}

	// 默认 dev
	if len(environment) == 0 {
		environment = dev
	}

	return nil
}

func GetEnv() string {
	return environment
}

func GetPid() int {
	return pid
}

func GetHostname() string {
	return hostname
}
