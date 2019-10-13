package http

import (
	"github.com/sepuka/chat/src/def"
	"net/http"
	"net/url"
	"time"

	"github.com/sarulabs/di"
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
				if cfg.HttpClient.Timeout > 0 {
					client.Timeout = time.Second * time.Duration(cfg.HttpClient.Timeout)
				}

				return client, nil
			},
		})
	})
}
