package view

import (
	"fmt"
	"time"

	"github.com/sepuka/chat/internal/domain"
)

const (
	infoFormat = "id %s\tcreated at %s\nweb port %d\nssh port %d"
)

type InfoFormatter struct {
	host *domain.VirtualHost
}

func NewInfoFormatter(host *domain.VirtualHost) *InfoFormatter {
	return &InfoFormatter{
		host: host,
	}
}

func (f *InfoFormatter) Format() []byte {
	return []byte(
		fmt.Sprintf(
			infoFormat, f.host.Container, f.host.CreatedAt.Format(time.RFC822), f.host.WebPort, f.host.SshPort,
		),
	)
}
