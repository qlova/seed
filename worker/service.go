package worker

import "bytes"

type Service struct {
	Assets []string
}

func (worker Service) Render() []byte {
	var b bytes.Buffer
	
	b.WriteString(`self.addEventListener('install', function(event) {
  event.waitUntil(
    caches.open("cache").then(function(cache) {
      return cache.addAll(
        ["/", `)
	
	for i, asset := range worker.Assets {
		b.WriteByte('"')
		b.WriteString(asset)
		b.WriteByte('"')
		if i < len(worker.Assets)-1 {
			b.WriteString(", ")
		}
	}
	
	b.WriteString(`]
      );
    })
  );
});
	
self.addEventListener('fetch', function(event) {
  event.respondWith(
    caches.open('mysite-dynamic').then(function(cache) {
      return cache.match(event.request).then(function(response) {
        var fetchPromise = fetch(event.request).then(function(networkResponse) {
          cache.put(event.request, networkResponse.clone());
          return networkResponse;
        })
        return response || fetchPromise;
      })
    })
  );
});
`)
	
	return b.Bytes()
}
