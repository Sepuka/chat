package cloud

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sepuka/chat/internal/config"

	"github.com/sepuka/chat/internal/domain"
)

type Cloud struct {
	client *ClientBuilder
	config *config.Config
}

func NewCloud(
	clientBuilder *ClientBuilder,
	cfg *config.Config,
) *Cloud {
	return &Cloud{
		client: clientBuilder,
		config: cfg,
	}
}

func (c *Cloud) Run(pool *domain.Pool, cmd domain.RemoteCmd) ([]byte, error) {
	secret, err := c.readKey(pool.Secret)
	if err != nil {
		return []byte{}, errors.Wrap(err, `unable to read private key file`)
	}

	client, err := c.client.Build(pool.Address.String(), secret)
	if err != nil {
		return []byte{}, errors.Wrap(err, `unable to build cloud client`)
	}

	return client.RemoteCmd(cmd)
}

func (c *Cloud) readKey(privateKey string) ([]byte, error) {
	return ioutil.ReadFile(privateKey)
}
