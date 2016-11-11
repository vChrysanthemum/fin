export pwd = $(shell sh -c 'pwd')

all:clean test build_in

install:
	@./install.sh
	@echo "success."

clean:
	rm -rf pkg/*
	rm -rf bin/*

build_in:
	go build -tags deadlock -o ./bin/in ./main 

test_ui:
	go test ./lib/ui -args ${pwd}

versioning:
	bash mkversion.sh 

build_ui:
	go build -o ./bin/ui ./test/ui.go
