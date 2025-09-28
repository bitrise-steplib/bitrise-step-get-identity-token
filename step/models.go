package step

import "github.com/bitrise-io/go-steputils/v2/stepconf"

type Input struct {
	BuildURL   string          `env:"build_url,required"`
	BuildToken stepconf.Secret `env:"build_api_token,required"`
	Subject    string          `env:"subject,required"`
	Audience   string          `env:"audience,required"`
	Verbose    bool            `env:"verbose,opt[true,false]"`
}

type Config struct {
	BuildURL   string
	BuildToken stepconf.Secret
	Subject    string
	Audience   string
}

type Result struct {
	IdentityToken string
}
