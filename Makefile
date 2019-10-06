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

telegram: build
	./chat telegram -c config/dev.yml

terminal: build
	./chat terminal ${instr} -c config/dev.yml

gen_tbl: build
	./chat db generate -c config/dev.yml

joomla_unpack:
	rm -rf docker/tariff/joomla/web/src/*
	unzip -qq docker/tariff/joomla/web/Joomla_3.9.11-Stable-Full_Package.zip -d /tmp/${project}

# usage: make joomla project=vasya (where vasya is user's nickname)
joomla: joomla_unpack
	cd docker/tariff/joomla && \
	export PROJECT=${project} && \
	docker-compose up --build

stop_all:
	docker stop $(docker ps -a -q);
	docker rm $(docker ps -a -q);
	docker rmi $(docker images -q);
	docker volume rm $(docker volume ls -q)