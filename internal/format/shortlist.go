package format

import (
	"fmt"
	"strings"

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
	details := make([]string, 0, len(hosts))
	for _, host := range hosts {
		details = append(details, fmt.Sprintf(`host %s created at %s`, host.Pool.Address, host.CreatedAt))
	}

	return strings.Join(details, "\n")
}
