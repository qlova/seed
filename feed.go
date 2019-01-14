package seed

import "fmt"
import "strconv"
import "net/http"

type Feed Seed

func (feed Feed) Refresh(q Script) {
	q.Javascript(q.Get(Seed(feed)).Element()+".onready();")
}

var feeds = make(map[string]func(Client))

func (seed Seed) AddFeed(template Seed, feed func(Client)) Feed {
	var WrapperSeed = New()
	WrapperSeed.SetSize(100, Auto)
	WrapperSeed.SetUnshrinkable()

	minified, err := mini(template.HTML())
	if err != nil {
		//Panic?
	}
	
	var ReplaceList string = ".replace(/"+template.GetSeed().id+"/g,'"+template.GetSeed().id+"-'+i)"
	//Each id needs to be replaced with an id with a unique suffix.
	//TODO support recursion.
	for _, child := range template.children {
		ReplaceList += ".replace(/"+child.GetSeed().id+"/g,'"+child.GetSeed().id+"-'+i)"
	}


	var id = fmt.Sprint(feed)
	
	feeds[id] = feed
	
	WrapperSeed.OnReady(func(q Script) {
		q.Javascript(q.Get(WrapperSeed).Element()+".onready = function() {")
		q.Javascript(`let request = new XMLHttpRequest(); request.open("GET", "/feeds/`+id+`"); request.onload = function() {`)
			q.Javascript(`if (request.response.length <= 0) return;`)
		
			q.Javascript(`let json = JSON.parse(request.response);`)
			
			q.Javascript(q.Get(WrapperSeed).Element()+`.innerHTML = "";`)
			q.Javascript(`for (let i = 0; i < json.length; i++) {`)
				q.Javascript(q.Get(WrapperSeed).Element()+`.innerHTML += `+strconv.Quote(string(minified))+ReplaceList)
			q.Javascript(`}`)

			q.Javascript(`for (let i = 0; i < json.length; i++) {`)

			q.Javascript(`get("`+template.id+`-"+i).data = json[i];`)
			
			//Figure out what content to replace.
			for _, child := range template.children {

				var text = string(child.GetSeed().content)
				if len(text) < 2 {
					continue
				}
				
				if text[0] == '{' && text[len(text)-1] == '}' {
					text = text[1:len(text)-1]
					var id = child.GetSeed().id

					q.Javascript(`get("`+id+`-"+i).innerHTML = json[i]["`+text+`"];`)
				}
			}
			q.Javascript(`}`)
			//TODO do this properly.
			
		
		q.Javascript(`}; request.send();`)
		q.Javascript(`};`)
		q.Javascript(q.Get(WrapperSeed).Element()+".onready();")
	})

	seed.Add(WrapperSeed)
	return Feed(WrapperSeed)
}

func feedHandler(w http.ResponseWriter, r *http.Request, id string) {
	if feed, ok := feeds[id]; ok {
		feed(Client{client{
			Request: r,
			ResponseWriter: w, 
		}})
	}
}
