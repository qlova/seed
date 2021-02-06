package dating

import (
	"encoding/json"
	"fmt"
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

	Time time.Time

	distance time.Duration
	nextTime func() time.Time
}

var Holidays = []Holiday{
	{
		Name:  "New Year's Eve",
		Image: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/48/Fanciful_sketch_by_Marguerite_Martyn_of_a_New_Years_Eve_celebration.jpg/1280px-Fanciful_sketch_by_Marguerite_Martyn_of_a_New_Years_Eve_celebration.jpg",
		nextTime: func() time.Time {
			var now = time.Now()
			return time.Date(
				now.Year(), 12, 31, 0, 0, 0, 0, time.Local,
			)
		},
	},

	{
		Name:  "Christmas",
		Image: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8f/NativityChristmasLights2.jpg/1280px-NativityChristmasLights2.jpg",
		nextTime: func() time.Time {
			var now = time.Now()
			return time.Date(
				now.Year(), 12, 25, 0, 0, 0, 0, time.Local,
			)
		},
	},
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
