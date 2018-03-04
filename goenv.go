package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	//Some will be loaded from env
	Some SomeConf `env:"SOME_"`

	SubSome SubSomeConf `env:"SUB_"`

	//Other not because it doesnt have the required tag
	Other OtherConf
}

type SomeConf struct {
	// Here we can set the next part of the env name,
	// the final env variable name will be formed by the
	// concatenation of all the env tags found until a single type is found,
	// also we can define a default value in case the env is not set
	VarString string `env:"STRING" envDefault:"ClassicStringIsClassic"`
	VarInt    int    `env:"INT" envDefault:"1313"`
}

type SubSomeConf struct {
	ImBool bool     `env:"BOOL"`
	Some   SomeConf `env:"SOME_"`
}

type OtherConf struct {
	VarString string
	VarInt    int
}

func Load(conf interface{}) {
	load(reflect.ValueOf(conf), "", "")
}

func load(conf reflect.Value, envTag, envDefault string) {
	// here conf could be either a struct or just a variable
	// if it's a variable we just set its value to the value of the
	// environment variable referenced by its tag, or its default, otherwise we recursively
	// set the struct value to the value returned by load(...) of each of its
	// individual fields

	if conf.Kind() == reflect.Ptr {
		reflectedConf := reflect.Indirect(conf)
		// we should only keep going if we can set values
		if reflectedConf.IsValid() && reflectedConf.CanSet() {
			value, ok := os.LookupEnv(envTag)
			// if the env variable is not set we just use the envDefault
			if !ok {
				value = envDefault
			}
			switch reflectedConf.Kind() {
			case reflect.Struct:
				for i := 0; i < reflectedConf.NumField(); i++ {
					if tag, ok := reflectedConf.Type().Field(i).Tag.Lookup("env"); ok {
						def, _ := reflectedConf.Type().Field(i).Tag.Lookup("envDefault")
						load(reflectedConf.Field(i).Addr(), envTag+tag, def)
					}
				}
				break
			// Here for each type we should make a cast of the env variable and then set the value
			case reflect.String:
				reflectedConf.SetString(value)
				break
			case reflect.Int:
				value, _ := strconv.Atoi(value)
				reflectedConf.Set(reflect.ValueOf(value))
				break
			case reflect.Bool:
				value, _ := strconv.ParseBool(value)
				reflectedConf.Set(reflect.ValueOf(value))
			}
		}

	}

}

func main() {
	conf := Config{}
	Load(&conf)
	fmt.Printf("\nReaded conf: %+v\n", conf)
}
