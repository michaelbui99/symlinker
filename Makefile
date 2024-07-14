clean:
	@rm ./bin

build: 
	@go build -o ./bin/symlinker ./main.go