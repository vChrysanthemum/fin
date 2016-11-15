export pwd = $(shell sh -c 'pwd')

all:clean test build_in

install:
	@./install.sh
	@echo "success."

clean:
	rm -rf pkg/*
	rm -rf bin/*

build_in:
	go install -tags deadlock ./src/main 
	@mv bin/main bin/in

test_ui:
	go test ./lib/ui -args ${pwd}

test_script:
	go test ./lib/script -args ${pwd}

versioning:
	bash mkversion.sh 

build_ui:
	go install -tags deadlock ./src/ui
