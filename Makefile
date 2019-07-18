CONFIG_PATH=config/dev.yml.dist

#create role hostinguser;
#create database hosting owner hostinguser;
#ALTER USER hostinguser WITH PASSWORD '1234';
#ALTER ROLE "hostinguser" WITH LOGIN;

init:
	dep ensure -v
	cp -n $(CONFIG_PATH) config/dev.yml

build:
	go build -o chat

start: build
	./chat telegram -c config/dev.yml

gen_tbl: build
	./chat db generate -c config/dev.yml
