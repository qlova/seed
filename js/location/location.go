package location

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

func Replace(url js.AnyString) script.Script {
	return js.Run("window.location.replace", url)
}
