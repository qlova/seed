package seed

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	//"strings"
	//"strconv"
	//"bufio"
	//"fmt"
	//"encoding/json"
	//"path"
)

var singleLocalConnection = false

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //r.Header.Get("Origin") == "https://realmoforder.com"
	},
}

//This will be single threaded.
/*type PostProduction map[string]map[string]string

var post PostProduction

func init() {
	post = make(PostProduction)

	fmt.Println(path.Dir(os.Args[0])+"/style.sss")
	file, err := os.Open(path.Dir(os.Args[0])+"/style.sss")
	if err != nil {
		fmt.Println("No "+path.Dir(os.Args[0])+"/style.sss could be found", err.Error());
		return
	}

	err = json.NewDecoder(file).Decode(&post)
	if err != nil {
		post = make(PostProduction)
		return
	}

	file.Close()
}

func (seed Seed) postProduction() {
	if style, ok := post[seed.manifest.Name+">"+seed.manifest.Description]; ok {
		for property, value := range style {
			if property == "text" {
				value, err := strconv.Unquote(strings.TrimSpace(value))
				if err != nil {
					continue
				}
				seed.SetContent(value)
			} else {
				seed.Set(property, value)
			}
		}
	}
}

func (p PostProduction) Add(seed Seed, style string) {
	if seed.manifest.Name == "" && seed.manifest.Description == "" {
		return
	}
	var table = p[seed.manifest.Name+">"+seed.manifest.Description]
	if table == nil {
		table = make(map[string]string)
	}

	var StyleReader = bufio.NewReader(strings.NewReader(style))
	for {
		property, err := StyleReader.ReadString(':')
		if err != nil {
			break
		}
		property = property[:len(property)-1]

		value, err := StyleReader.ReadString(';')
		if err != nil {
			break
		}
		value = value[:len(value)-1]

		table[property] = value
	}

	p[seed.manifest.Name+">"+seed.manifest.Description] = table
	post.Save()
}

func (p PostProduction) Save() {
	file, err := os.Create("style.sss")
	if err != nil {
		fmt.Println("Could not create style.sss!");
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(&post)
	if err != nil {
		fmt.Println("Could not save SeedStyleSheet file!");
		return
	}
}*/

var localSockets = make(map[string]*websocket.Conn)
var reloading = false

func socket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	localSockets[r.RemoteAddr] = c

	reloading = false

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			singleLocalConnection = localClients == 1

			if singleLocalConnection && !reloading {
				cleanup()
				os.Exit(0)
			} else {
				delete(localSockets, r.RemoteAddr)
				localClients--
				return
			}
		}

		/*var StyleModification = string(data)
		if len(StyleModification) == 0 || StyleModification[0] != '#' {
			continue
		}

		fmt.Println(StyleModification)

		var reader = bufio.NewReader(strings.NewReader(StyleModification))

		id, err := reader.ReadString('{')
		if err != nil {
			continue
		}
		id  = id[:len(id)-1]
		id = strings.TrimSpace(id)
		id  = id[1: len(id)]
		fmt.Println("Saving changes for ", id)

		css, err := reader.ReadString('}')
		css = css[:len(css)-1]
		css = strings.TrimSpace(css)
		fmt.Println("(", css, ")")


		var seed, ok = allSeeds[id]
		if ok {
			post.Add(Seed{seed}, css)
		}*/
	}
}
