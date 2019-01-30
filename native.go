package seed

import "bytes"
import "fmt"

func (seed Seed) JSX() ([]byte) {
	var jsx bytes.Buffer
	var native = seed.ReactNative()

	//Write the tag.
	if native.tag != "div" {
		jsx.WriteByte('<')
		jsx.WriteString(native.tag)
		jsx.WriteByte(' ')
		if native.attr != "" {
			jsx.WriteString(native.attr)
			jsx.WriteByte(' ')
		}
		jsx.WriteString("ref={(r) => this.")
		jsx.WriteString(fmt.Sprint(seed.id))
		jsx.WriteString(" = r}")
		
		/*if data := seed.Style.Bytes(); !seed.styled && data != nil {
			html.WriteString(" style='")
			html.Write(data)
			html.WriteByte('\'')
		}
		
		if seed.onclick != nil {
			html.WriteString(" onclick='")
			html.WriteString(script.ToJavascript(seed.onclick))
			html.WriteByte('\'')
		}
		
		if seed.onchange != nil {
			html.WriteString(" onchange='")
			html.WriteString(script.ToJavascript(seed.onchange))
			html.WriteByte('\'')
		}*/
		
		jsx.WriteByte('>')
	
	
		if native.content != nil {
			jsx.Write(seed.content)
		}
	}
	
	for _, child := range seed.children {
		jsx.Write(child.Root().JSX())
	}
	
	if native.tag != "div" {
		jsx.WriteString("</")
		jsx.WriteString(native.tag)
		jsx.WriteByte('>')
	}
	
	return jsx.Bytes()
}
