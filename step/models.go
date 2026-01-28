package step

import "github.com/bitrise-io/go-steputils/v2/stepconf"

type Input struct {
	BuildURL   string          `env:"build_url,required"`
	BuildToken stepconf.Secret `env:"build_api_token,required"`
	Audience   string          `env:"audience,required"`
	Verbose    bool            `env:"verbose,opt[true,false]"`
}

type Config struct {
	BuildURL   string
	BuildToken stepconf.Secret
	Audience   string
	Verbose    bool
}

type Result struct {
	IdentityToken string
}
