package main

import (
	"fmt"
	"reflect"
	"time"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}

func main() {

	//i := 0
	//getTableName(i)
	//return

	// 测试1，解析到struct
	cfg := config{}
	if err := Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}

func Test() {

	// 判断是否是指针类型
	t := "a"

	fmt.Println(reflect.ValueOf(t).Kind() != reflect.Ptr)

}
