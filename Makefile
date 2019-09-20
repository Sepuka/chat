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

repeater: build
	./chat telegram --mode=repeater -c config/dev.yml

hosting: build
	./chat telegram --mode=hosting -c config/dev.yml

gen_tbl: build
	./chat db generate -c config/dev.yml

joomla_unpack:
	rm -rf docker/tariff/joomla/web/src/*
	unzip -q docker/tariff/joomla/web/Joomla_3.9.11-Stable-Full_Package.zip -d docker/tariff/joomla/web/src/

joomla: joomla_unpack
	cd docker/tariff/joomla && docker-compose up --build

stop_all:
	docker stop $(docker ps -a -q);
	docker rm $(docker ps -a -q);
	docker rmi $(docker images -q);
	docker volume rm $(docker volume ls -q)