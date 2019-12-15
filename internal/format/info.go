package format

import (
	"fmt"
	"time"

	"github.com/sepuka/chat/internal/domain"
)

const (
	infoFormat = "#%d\t%s\tcreated at %s"
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
	return []byte(fmt.Sprintf(infoFormat, f.host.Id, f.host.Container, f.host.CreatedAt.Format(time.RFC822)))
}
