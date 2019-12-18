package view

import (
	"fmt"

	"github.com/sepuka/chat/internal/domain"
)

type ContainersListFormatter struct {
	hosts []*domain.VirtualHost
}

func NewContainersListFormatter(hosts []*domain.VirtualHost) *ContainersListFormatter {
	return &ContainersListFormatter{
		hosts: hosts,
	}
}

func (f *ContainersListFormatter) Format() (list string) {
	for _, host := range f.hosts {
		list += fmt.Sprintf("%s\n", host.Container)
	}

	return list
}
