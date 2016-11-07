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

build_editor:
	go build -o ./bin/editor ./test/editor.go

build_select:
	go build -o ./bin/select ./test/select.go

build_table:
	go build -o ./bin/table ./test/table.go

build_script:
	go build -o ./bin/script ./test/script.go
