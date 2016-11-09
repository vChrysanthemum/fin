export pwd = $(shell sh -c 'pwd')

all:clean test build

clean:
	rm -rf pkg/*
	rm -rf bin/*

build_inn:
	go build -tags deadlock -o ./bin/inn ./main 

test_ui:
	go test ./lib/ui -args ${pwd}

versioning:
	bash mkversion.sh 

build_ui:
	go build -o ./bin/ui ./test/ui.go
