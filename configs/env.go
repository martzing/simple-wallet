package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
}

func NewEnvironment() *Environment {
	env := Environment{}
	godotenv.Load()

	return &env
}

func (env *Environment) GetString(key string, opts ...string) *string {
	var _default string
	for _, opt := range opts {
		_default = opt
	}
	s := os.Getenv(key)
	if s == "" {
		s = _default
	}
	return &s
}

func (env *Environment) GetInt(key string, opts ...int) *int {
	var _default string
	for _, opt := range opts {
		_default = fmt.Sprint(opt)
	}
	s := env.GetString(key)
	if *s == "" {
		s = &_default
	}
	i, err := strconv.Atoi(*s)

	if err != nil {
		panic(err)
	}

	return &i
}

func (env *Environment) GetBool(key string, opts ...bool) *bool {
	var _default string

	for _, opt := range opts {
		_default = fmt.Sprint(opt)
	}

	s := env.GetString(key)

	if *s == "" {
		s = &_default
	}
	i, err := strconv.ParseBool(*s)

	if err != nil {
		panic(err)
	}

	return &i
}
