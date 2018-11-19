package seed

var UsingMap bool

func Map() Seed {
	var m = New()
	
	UsingMap = true
	
	var RealMap = New()
	RealMap.SetWidth("100%")
	RealMap.SetHeight("100%")
	m.Add(RealMap)
	
	var Script = New()
	Script.SetContent(`
	<div class="circle" style="left:50vw; top:50vh; z-index: 100000; position: fixed; pointer-events: none;"></div>
	<script>
var map = L.map("`+RealMap.ID()+`", {
    center: [51.505, -0.09],
    zoom: 13
});
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
	maxZoom: 19,
	attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);
</script>`)
	m.Add(Script)
	
	return m
}
