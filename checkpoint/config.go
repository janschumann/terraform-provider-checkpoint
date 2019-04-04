package checkpoint

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/service/host"
)

type Config struct {
	User        string
	Password    string
	Hostname    string
	Insecure    bool
	SessionName string
}

type CPClient struct {
	host *host.HostService
}

func (c *Config) Client() (interface{}, error) {
	config := checkpoint.NewConfig()
	if c.User != "" && c.Password != "" {
		config.WithCredentials(credentials.NewStaticCredentials(c.User, c.Password))
	}
	if c.Hostname != "" {
		config.WithApiHost(c.Hostname)
	}
	if c.Insecure {
		config.WithInsecure(c.Insecure)
	}
	if c.SessionName != "" {
		config.WithSessionName(c.SessionName)
	}

	sdkClient := client.Must(client.New(config))

	cpClient := &CPClient{
		host: host.New(sdkClient),
	}

	return cpClient, nil
}
