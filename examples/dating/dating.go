package dating

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"sort"
	"time"

	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

type Holiday struct {
	Name     string
	Image    string
	Distance string

	Time time.Time `mirror:"ignore"`

	distance time.Duration
	nextTime func() time.Time
}

var Holidays = []Holiday{}

func readPopular(r io.Reader) {
	var rawHolidays []HolidayJSON
	var err = json.NewDecoder(r).Decode(&rawHolidays)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, hol := range rawHolidays {
		t := hol.Start
		if hol.Substitute {
			continue
		}
		Holidays = append(Holidays, Holiday{
			Name:     hol.Name,
			Time:     hol.Start,
			Image:    "https://loremflickr.com/500/500/" + hol.Name + "?lock=1",
			nextTime: func() time.Time { return t },
		})
	}
}

var Custom = []Holiday{}

func AddCustom(name string, date time.Time, hours string) {
	Custom = append(Custom, Holiday{
		Name:     name,
		Image:    "https://picsum.photos/100?" + fmt.Sprint(time.Now().UnixNano()),
		Time:     date,
		nextTime: func() time.Time { return date },
	})
}

func update(h []Holiday) {
	for i := range h {
		if h[i].nextTime == nil {
			t := h[i].Time
			h[i].nextTime = func() time.Time { return t }
		}
		h[i].distance = h[i].nextTime().Sub(time.Now())
		h[i].Distance = fmt.Sprintf("%v days", int(math.Ceil(h[i].distance.Hours()/24)))
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i].distance < h[j].distance
	})
}

func GetHolidays() []Holiday {
	update(Holidays)

	return Holidays
}

func GetCustom() []Holiday {
	update(Custom)

	return Custom
}

func SaveCustom() client.Script {
	b, err := json.Marshal(Custom)
	if err != nil {
		panic(err)
	}

	return js.Func("window.localStorage.setItem").Run(client.NewString("custom.dates"), client.NewString(string(b)))
}

func LoadCustom(custom string) {
	err := json.Unmarshal([]byte(custom), &Custom)
	if err != nil {
		fmt.Println(err)
	}
}
