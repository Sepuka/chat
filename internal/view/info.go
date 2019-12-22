package view

import (
	"fmt"
	"time"

	"github.com/sepuka/chat/internal/domain"
)

const (
	infoFormat = "%s%s\tcreated at %s\nweb\t%s:%d\nssh\t%s:%d"
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
	var alias string

	if f.host.Alias != "" {
		alias = fmt.Sprintf(` (%s)`, f.host.Alias)
	} else {
		alias = ""
	}

	return []byte(
		fmt.Sprintf(
			infoFormat,
			f.host.Container,
			alias,
			f.host.CreatedAt.Format(time.RFC822),
			f.host.Pool.Address,
			f.host.WebPort,
			f.host.Pool.Address,
			f.host.SshPort,
		),
	)
}
