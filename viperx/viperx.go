package viperx

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Viperx struct {
	client *viper.Viper
}

func NewViper() *Viperx {
	return &Viperx{
		client: viper.New(),
	}
}

func (v *Viperx) Read(target interface{}, path, name string) error {
	// register config file path
	v.client.AddConfigPath(path)

	// register and read config file name
	v.client.SetConfigName(name)
	err := v.client.MergeInConfig()
	if err != nil {
		return err
	}

	return v.client.Unmarshal(target, func(dc *mapstructure.DecoderConfig) {
		// TODO: make this configurable if there's ever a need to honor snake case json struct tag
		// dc.TagName = "json"
	})
}
