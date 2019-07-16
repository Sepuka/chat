package http

import (
	"net/http"
	"net/url"

	"github.com/sarulabs/di"
	"github.com/sepuka/chat/src/def"
)

const ClientDef = "http.client"

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: ClientDef,
			Build: func(ctx def.Context) (interface{}, error) {
				client := &http.Client{}
				var cfg = ctx.Get(def.CfgDef).(def.Config)
				proxy, err := url.Parse(cfg.HttpClient.Proxy)
				if err != nil {
					return nil, err
				}

				client.Transport = &http.Transport{
					Proxy: http.ProxyURL(proxy),
				}

				return client, nil
			},
		})
	})
}
