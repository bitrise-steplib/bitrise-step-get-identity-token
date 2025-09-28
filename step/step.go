package step

import (
	"github.com/bitrise-io/go-steputils/v2/export"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-get-identity-token/api"
)

const identityTokenKey = "BITRISE_IDENTITY_TOKEN"

type TokenFetcher struct {
	inputParser   stepconf.InputParser
	envRepository env.Repository
	exporter      export.Exporter
	logger        log.Logger
}

func NewTokenFetcher(inputParser stepconf.InputParser, envRepository env.Repository, exporter export.Exporter, logger log.Logger) TokenFetcher {
	return TokenFetcher{
		inputParser:   inputParser,
		envRepository: envRepository,
		exporter:      exporter,
		logger:        logger,
	}
}

func (r TokenFetcher) ProcessConfig() (Config, error) {
	var input Input
	err := r.inputParser.Parse(&input)
	if err != nil {
		return Config{}, err
	}

	stepconf.Print(input)
	r.logger.Println()
	r.logger.EnableDebugLog(input.Verbose)

	return Config{
		BuildURL:   input.BuildURL,
		BuildToken: input.BuildToken,
		Subject:    input.Subject,
		Audience:   input.Audience,
	}, nil
}

func (r TokenFetcher) Run(config Config) (Result, error) {
	client := api.NewDefaultAPIClient(config.BuildURL, config.BuildToken, r.logger)

	parameter := api.GetIdentityTokenParameter{
		Subject:  config.Subject,
		Audience: config.Audience,
	}
	token, err := client.GetIdentityToken(parameter)
	if err != nil {
		return Result{}, err
	}

	r.logger.Donef("Identity token fetched.")

	return Result{
		IdentityToken: token,
	}, nil
}

func (r TokenFetcher) Export(result Result) error {
	r.logger.Printf("The following outputs are exported as environment variables:")

	values := map[string]string{
		identityTokenKey: result.IdentityToken,
	}

	for key, value := range values {
		err := r.exporter.ExportOutput(key, value)
		if err != nil {
			return err
		}

		r.logger.Donef("$%s = %s", key, value)
	}

	return nil
}
