package script

func (q Script) SubString(s String, start, end Int) String {
	return q.js.Call(s.LanguageType().Raw()+".substr", start, end).String()
}
