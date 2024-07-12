package toolbox

import "github.com/dawsonalex/site-builder/tpl"

func GetComponent[t tpl.ComponentType](name string) tpl.Component[t] {
	switch {
	case "hero":
		return tpl.Component[tpl.HeroConfig]{}
	}
}
