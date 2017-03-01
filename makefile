export pwd = $(shell sh -c 'pwd')

all:clean test build_fin

install:
	@mkdir -p "$(HOME)/.fin"
	@rm -f "$(HOME)/.fin/res"
	@ln -sf "$(pwd)/res" "$(HOME)/.fin/res"
	@echo "success."

clean:
	rm -rf pkg/*
	rm -rf bin/*

build_fin:
	go install -tags deadlock ./src/main 
	@mv bin/main bin/fin

test:test_ui test_script

test_ui:
	go test ./lib/ui -args ${pwd}

test_script:
	go test ./lib/script -args ${pwd}

versioning:
	bash mkversion.sh 

build_ui:
	go install -tags deadlock ./src/ui
