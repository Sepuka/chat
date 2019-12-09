package cloud

import (
	"time"

	"github.com/pkg/errors"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/domain"
	"golang.org/x/crypto/ssh"
)

type ClientBuilder struct {
	appCfg *config.Config
}

func NewClientBuilder(cfg *config.Config) *ClientBuilder {
	return &ClientBuilder{
		appCfg: cfg,
	}
}

func (c *ClientBuilder) Build(addr string, key []byte) (*SshClient, error) {
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, errors.Wrap(err, `cannot parse private key`)
	}

	cfg := &ssh.ClientConfig{
		Config: ssh.Config{},
		User:   c.appCfg.Pool.Login,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return NewSslClient(cfg, addr), nil
}

type SshClient struct {
	cfg  *ssh.ClientConfig
	addr string
}

func NewSslClient(
	clientConfig *ssh.ClientConfig,
	addr string,
) *SshClient {
	return &SshClient{
		cfg:  clientConfig,
		addr: addr,
	}
}

func (c *SshClient) RemoteCmd(cmd domain.RemoteCmd) ([]byte, error) {
	client, err := ssh.Dial(`tcp`, c.addr+`:22`, c.cfg)
	if err != nil {
		return []byte{}, errors.Wrap(err, `cannot connect to server`)
	}

	session, err := client.NewSession()
	if err != nil {
		return []byte{}, errors.Wrap(err, `cannot init session`)
	}
	defer func() {
		_ = session.Close()
	}()

	return session.CombinedOutput(string(cmd))
}
