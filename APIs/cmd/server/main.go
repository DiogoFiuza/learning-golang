package main

import "github.com/DiogoFiuza/learning-golang/APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
