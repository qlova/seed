package interfaces 

import "github.com/qlova/seed/style"

//Everything is an app.
type App interface {
	ID() string
	
	SetName(name string)
	SetDescription(description string)
	SetIcon(path string)
	
	GetStyle() *style.Style
	
	Add(App)
	GetParent() App 
	SetParent(App)
	GetChildren() []App
	
	Page() bool
	
	SetContent(content string)

	Render() []byte
}
