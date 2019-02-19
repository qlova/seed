package service

import "bytes"

func NewWorker() Worker {
	return Worker{
		Assets: make(map[string]bool),
	}
}

type Worker struct {
	Assets map[string]bool
}

func (worker Worker) Render() []byte {
	var b bytes.Buffer
	
	b.WriteString(`self.addEventListener('install', function(event) {
  event.waitUntil(
    caches.open("cache").then(function(cache) {
      return cache.addAll(
        [".", `)

	var i = 0
	for asset := range worker.Assets {
		b.WriteByte('"')
		b.WriteString(asset)
		b.WriteByte('"')
		if i < len(worker.Assets)-1 {
			b.WriteString(", ")
		}
		i++
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
      	if (navigator.onLine) {
	        var fetchPromise = fetch(event.request).then(function(networkResponse) {
	          cache.put(event.request, networkResponse.clone());
	          return networkResponse;
	        })
	        return response || fetchPromise;
	    } else {
	    	return response;
	    }
      })
    })
  );
});
`)
	
	return b.Bytes()
}
