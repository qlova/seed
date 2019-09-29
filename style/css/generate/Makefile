all:
	go build 
	./generate properties > ../properties.go
	./generate values > ../values.go
	go fmt ../*.go
