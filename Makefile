CONFIG_PATH=config/dev.yml.dist
PROGRAM_NAME=hosting

init:
	dep ensure -v
	cp -n $(CONFIG_PATH) config/dev.yml

dependencies:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure

build:
	go build -o $(PROGRAM_NAME)

telegram: build
	./${PROGRAM_NAME} telegram -c config/dev.yml

terminal: build
	./${PROGRAM_NAME} terminal ${instr} -c config/dev.yml -u ${user}

db:
	docker build -t hosting_db docker/parts/pg/
	docker run -p 5433:5432 hosting_db

pgdump:
	PGPASSWORD="1234" pg_dump -h localhost -p 5433 -U hosting -f docker/parts/pg/dump.sql hosting

gen_tbl: build
	./hosting db generate -c config/dev.yml

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