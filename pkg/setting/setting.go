package setting

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting initializes the basic preference of blogie
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, cfg := range configs {
		if cfg != "" {
			vp.AddConfigPath(cfg)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()

	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		// in is callback parameter in OnConfigChange
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			err := s.ReloadAllSection()
			if err != nil {
				log.Fatalf("section reload failed: %v", err)
			}
		})
	}()
}
