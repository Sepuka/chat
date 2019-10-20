package def

import (
	"github.com/sarulabs/di"
)

var (
	Container  di.Container
	components []creatorFn
)

type (
	creatorFn func(builder *di.Builder, cfg Config) error
	Context    = di.Container
)

func Register(fn creatorFn) {
	components = append(components, fn)
}

func Build(params Config) error {
	builder, err := di.NewBuilder()
	if err != nil {
		return err
	}

	for _, fnc := range components {
		if err := fnc(builder, params); err != nil {
			return err
		}
	}

	Container = builder.Build()

	return nil
}

func GetByTag(tag string) []interface{} {
	var defs []interface{}

	for _, def := range Container.Definitions() {
		for _, defTag := range def.Tags {
			if defTag.Name == tag {
				var content interface{}
				if err := Container.Fill(def.Name, &content); err != nil {
					panic(err)
				}
				defs = append(defs, content)
			}
		}
	}

	return defs
}
