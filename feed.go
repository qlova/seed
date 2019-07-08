package seed

import "fmt"
import "strings"
import "strconv"
import "net/http"
import "bytes"

import "github.com/qlova/seed/style/css"

import "github.com/qlova/seed/script"

//Shuold this be stored in the harvester?
var feeds = make(map[string]func(User) Food)

//A feed is used to transfer dynamic repeatable data from the server to the application.
//For example, a news feed, a blog, comments etc.
type Feed struct {
	Seed

	//Am I within another feed?
	within *Feed

	//What will we be feeding to the application?
	food interface{}

	//A reference to this feed's handler.
	handler string
}

type Food interface{}

var feed_id = 0

func NewFeed(food interface{}) Feed {
	var seed = New()
	seed.SetSize(100, Auto)
	seed.SetUnshrinkable()

	seed.SetDisplay(css.Flex)
	seed.SetFlexDirection(css.Row)
	seed.SetFlexWrap(css.Wrap)

	feed_id++

	return Feed{
		Seed: seed,
		food: food,

		handler: fmt.Sprint(feed_id),
	}
}

func NewFeedWithin(parent Feed, food interface{}) Feed {
	var feed = NewFeed(food)
	feed.within = &parent

	return feed
}

//Refresh the feeds content from the server.
func (feed Feed) Script(q Script) script.Feed {
	return script.Feed{script.Seed{
		ID: feed.id,
		Q:  q,
	}}
}

//Run a script when this feed refreshes.
func (feed Feed) OnRefresh(f func(Script)) {
	feed.OnReady(func(q Script) {
		q.Javascript(feed.Script(q).Element() + ".onrefresh = function() {")
		f(q)
		q.Javascript("}); ")
		q.Javascript(feed.Script(q).Element() + ".onrefresh()")
	})
}

//Associate this feed to a parent, call .As(template_seed) in order to display the feed.
func (feed Feed) AddTo(parent Interface) feeder {
	return feeder{
		feed: feed,
		seed: parent.Root(),
	}
}

//A feeder is a feed builder.
type feeder struct {
	feed Feed
	seed Seed
}

//Add this feed to the parent as described in the template.
func (f feeder) As(template Template) Feed {
	f.seed.Add(template)

	//Subfeed.
	if f.feed.within != nil {

		var handler = f.feed.food.(func(User) Food)

		//var parent_handler = f.feed.within.handler

		//Add the handler to a map..
		feeds[f.feed.handler] = func(user User) Food {
			//TODO user.SetValue(feeds[parent_handler](user))
			return handler(user)
		}

	} else {
		//Top level feed.
		var handler func(User) Food

		switch food := f.feed.food.(type) {

		case func() Food:
			handler = func(User) Food {

				return food()
			}

		case func(User) Food:
			handler = food

		//Try to wrap food in a handler and send it to the application.
		default:
			handler = func(User) Food {
				return f.feed.food
			}

		}

		//Add the handler to a map..
		feeds[f.feed.handler] = handler
	}

	f.feed.OnReady(func(q Script) {
		q.Javascript(f.feed.Script(q).Element() + ".index = window.localStorage.getItem('" + f.feed.Script(q).ID + "_index') || '0';")
		//Call this refresh instead of onready?
		q.Javascript(f.feed.Script(q).Element() + ".onready = function() {")

		q.Javascript(`let index = "";`)

		for parent := &f.feed; parent != nil; parent = parent.within {
			if parent.within != nil {
				q.Javascript(`index += "/"+(` + parent.Script(q).Element() + `.index || "0");`)
			}
		}

		q.Javascript(`let request = new XMLHttpRequest(); request.open("GET", "/feeds/` + f.feed.handler + `"+index); request.onload = function() {`)
		q.Javascript(`if (request.response.length <= 0) return;`)

		q.Javascript(`let json = JSON.parse(request.response);`)

		q.Javascript(f.feed.Script(q).Element() + `.data = json;`)
		q.Javascript(f.feed.Script(q).Element() + `.innerHTML = "";`)

		q.Javascript(`for (let i = 0; i < json.length; i++) {`)

		q.Javascript("let data = json[i];")
		//Here we need to generate Javascript that can construct a seed from a Template.
		q.Javascript(f.feed.Script(q).Element() + ".appendChild(" + template.Render(q) + ");")

		q.Javascript(`}`)
		//TODO do this properly.

		q.Javascript(`}; request.send();`)
		q.Javascript(`};`)
		q.Javascript(f.feed.Script(q).Element() + ".onready();")
		q.Javascript(`if (` + f.feed.Script(q).Element() + ".onrefresh) " + f.feed.Script(q).Element() + ".onrefresh();")

	})

	f.seed.Add(f.feed)

	return f.feed

	//Minify the template's HTML.
	minified, err := mini(template.Root().HTML(Default))
	if err != nil {
		//Panic?
	}

	//We replace all of the id's with the items's index as a suffix.
	var ReplaceList string = ".replace(/id=" + template.Root().id + "/g,'id=" + template.Root().id + "-'+i)"

	//Each id needs to be replaced with an id with a unique suffix.
	//TODO support recursion.
	for _, child := range template.Root().children {
		ReplaceList += ".replace(/id=" + child.Root().id + "/g,'id=" + child.Root().id + "-'+i)"
	}

	//TODO handle mutexes below.

	//The template onready callback.
	var onready bytes.Buffer
	template.Root().buildOnReady(0, &onready)

	//This is the refresh function of the feed, send a request to the server, recieve feed and populate children with the feed's data.
	f.feed.OnReady(func(q Script) {
		q.Javascript(f.feed.Script(q).Element() + ".index = window.localStorage.getItem('" + f.feed.Script(q).ID + "_index') || '0';")
		//Call this refresh instead of onready?
		q.Javascript(f.feed.Script(q).Element() + ".onready = function() {")

		q.Javascript(`let index = "";`)

		for parent := &f.feed; parent != nil; parent = parent.within {
			if parent.within != nil {
				q.Javascript(`index += "/"+(` + parent.Script(q).Element() + `.index || "0");`)
			}
		}

		q.Javascript(`let request = new XMLHttpRequest(); request.open("GET", "/feeds/` + f.feed.handler + `"+index); request.onload = function() {`)
		q.Javascript(`if (request.response.length <= 0) return;`)

		q.Javascript(`let json = JSON.parse(request.response);`)

		q.Javascript(f.feed.Script(q).Element() + `.data = json;`)

		q.Javascript(f.feed.Script(q).Element() + `.innerHTML = "";`)
		q.Javascript(`for (let i = 0; i < json.length; i++) {`)
		q.Javascript(f.feed.Script(q).Element() + `.innerHTML += ` + strconv.Quote(string(minified)) + ReplaceList + ";")

		q.Javascript(`}`)

		q.Javascript(`for (let i = 0; i < json.length; i++) {`)

		q.Javascript(`get("` + template.Root().id + `-"+i).data = json;`)
		q.Javascript(`get("` + template.Root().id + `-"+i).index = +i+1;`)

		//Run OnReady
		q.Javascript("eval(" + strconv.Quote(onready.String()) + ReplaceList + ");")

		//Figure out what content to replace.
		for _, child := range template.Root().children {

			var text = string(child.Root().content)
			if len(text) < 2 {
				continue
			}

			if text[0] == '{' && text[len(text)-1] == '}' {
				text = text[1 : len(text)-1]
				var id = child.Root().id

				q.Javascript(`get("` + id + `-"+i).innerHTML = json[i]["` + text + `"];`)
			}
		}
		q.Javascript(`}`)
		//TODO do this properly.

		q.Javascript(`}; request.send();`)
		q.Javascript(`};`)
		q.Javascript(f.feed.Script(q).Element() + ".onready();")
		q.Javascript(`if (` + f.feed.Script(q).Element() + ".onrefresh) " + f.feed.Script(q).Element() + ".onrefresh();")
	})

	f.seed.Add(f.feed)

	return f.feed
}

func feedHandler(w http.ResponseWriter, r *http.Request, id string) {
	var splits = strings.Split(id, "/")
	var indices []int = make([]int, len(splits)-1)
	var err error

	if len(splits) > 1 {
		for i := range splits[1:] {
			indices[i], err = strconv.Atoi(splits[1+i])
			if err != nil {
				return
			}
			indices[i]--
		}
	}
	if feed, ok := feeds[splits[0]]; ok {
		var user = User{}.FromHandler(w, r)
		user.SetIndices(indices)
		user.Send(feed(user))
	}
}
