package format

import (
	"fmt"
	"strings"
	"time"

	"github.com/sepuka/chat/internal/domain"
)

type ShortHostsListFormatter struct {
	source domain.ClientSource
}

func NewShortHostsListFormatter(source domain.ClientSource) *ShortHostsListFormatter {
	return &ShortHostsListFormatter{
		source: source,
	}
}

func (f *ShortHostsListFormatter) Format(hosts []*domain.VirtualHost) string {
	var (
		length = len(hosts)
		result string
	)
	switch f.source {
	case domain.Telegram:
		result = fmt.Sprintf("You have %d hosts:\n%s", length, f.extendList(hosts))
	case domain.Terminal:
		result = fmt.Sprintf(`You have %d hosts`, length)
	}

	return result
}

func (f *ShortHostsListFormatter) extendList(hosts []*domain.VirtualHost) string {
	var (
		details = make([]string, 0, len(hosts))
		id      string
	)
	for _, host := range hosts {
		id = host.Container[:12]
		format := host.CreatedAt.Format(time.RFC822)
		details = append(details, fmt.Sprintf("%s@%s created at %s", id, host.Pool.Address, format))
	}

	return strings.Join(details, "\n")
}
