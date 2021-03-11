package settings

import "github.com/GuilhermeFirmiano/grok"

//Settings ...
type Settings struct {
	Grok *grok.Settings `yaml:"grok"`
}
