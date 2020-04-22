package js

var imports = make(map[string]string)

//NewImport creates a new import that can be used with the #import macro.
func NewImport(path string, data string) {
	imports[path] = data
}

//Imports returns a map of js imports.
func Imports() map[string]string {
	return imports
}
