all: build copy

build:
	go build .

copy:
	mkdir ${HOME}/.proxcli
	cp ./proxcli.yml ${HOME}/.proxcli

