package seed

import "fmt"
import "strconv"
import "net/http"

type Feed Seed

func (feed Feed) Refresh(q Script) {
	q.Javascript(Seed(feed).Script(q).Element()+".onready();")
}

func (feed Feed) OnRefresh(f func(Script)) {
	Seed(feed).OnReady(func(q Script) {
		q.Javascript(Seed(feed).Script(q).Element()+".onrefresh = function() {")
		f(q)
		q.Javascript("}); ")
		q.Javascript(Seed(feed).Script(q).Element()+".onrefresh()")
	})
}

var feeds = make(map[string]func(User))

func (seed Seed) AddFeed(template Interface, feed func(User)) Feed {
	var WrapperSeed = New()
	WrapperSeed.SetSize(100, Auto)
	WrapperSeed.SetUnshrinkable()

	minified, err := mini(template.Root().HTML(Default))
	if err != nil {
		//Panic?
	}
	
	var ReplaceList string = ".replace(/"+template.Root().id+"/g,'"+template.Root().id+"-'+i)"
	//Each id needs to be replaced with an id with a unique suffix.
	//TODO support recursion.
	for _, child := range template.Root().children {
		ReplaceList += ".replace(/"+child.Root().id+"/g,'"+child.Root().id+"-'+i)"
	}


	var id = fmt.Sprint(feed)
	
	feeds[id] = feed
	
	WrapperSeed.OnReady(func(q Script) {
		q.Javascript(WrapperSeed.Script(q).Element()+".onready = function() {")
		q.Javascript(`let request = new XMLHttpRequest(); request.open("GET", "/feeds/`+id+`"); request.onload = function() {`)
			q.Javascript(`if (request.response.length <= 0) return;`)
		
			q.Javascript(`let json = JSON.parse(request.response);`)

			q.Javascript(WrapperSeed.Script(q).Element()+`.data = json;`)
			
			q.Javascript(WrapperSeed.Script(q).Element()+`.innerHTML = "";`)
			q.Javascript(`for (let i = 0; i < json.length; i++) {`)
				q.Javascript(WrapperSeed.Script(q).Element()+`.innerHTML += `+strconv.Quote(string(minified))+ReplaceList)
			q.Javascript(`}`)

			q.Javascript(`for (let i = 0; i < json.length; i++) {`)

			q.Javascript(`get("`+template.Root().id+`-"+i).data = json[i];`)
			
			//Figure out what content to replace.
			for _, child := range template.Root().children {

				var text = string(child.Root().content)
				if len(text) < 2 {
					continue
				}
				
				if text[0] == '{' && text[len(text)-1] == '}' {
					text = text[1:len(text)-1]
					var id = child.Root().id

					q.Javascript(`get("`+id+`-"+i).innerHTML = json[i]["`+text+`"];`)
				}
			}
			q.Javascript(`}`)
			//TODO do this properly.
			
		
		q.Javascript(`}; request.send();`)
		q.Javascript(`};`)
		q.Javascript(WrapperSeed.Script(q).Element()+".onready();")
		q.Javascript(`if (`+WrapperSeed.Script(q).Element()+".onrefresh) "+WrapperSeed.Script(q).Element()+".onrefresh();")
	})

	seed.Add(WrapperSeed)
	return Feed(WrapperSeed)
}

func feedHandler(w http.ResponseWriter, r *http.Request, id string) {
	if feed, ok := feeds[id]; ok {
		feed(User{}.FromHandler(w, r))
	}
}
