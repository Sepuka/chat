package http

import (
	"net/http"
	"net/url"

	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/def"
)

const ClientDef = "http.client"

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: ClientDef,
			Build: func(ctx def.Context) (interface{}, error) {
				client := &http.Client{}
				var (
					cfg       = ctx.Get(def.CfgDef).(def.Config)
					transport = &http.Transport{}
				)

				if len(cfg.HttpClient.Proxy) > 0 {
					proxy, err := url.Parse(cfg.HttpClient.Proxy)
					if err != nil {
						return nil, err
					}
					transport.Proxy = http.ProxyURL(proxy)
				}

				client.Transport = transport

				return client, nil
			},
		})
	})
}
