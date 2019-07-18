CONFIG_PATH=config/dev.yml.dist

#create role hostinguser;
#create database hosting owner hostinguser;
#ALTER USER hostinguser WITH PASSWORD '1234';
#ALTER ROLE "hostinguser" WITH LOGIN;

init:
	dep ensure -v
	cp -n $(CONFIG_PATH) config/dev.yml

start:
	go build -o chat && ./chat repeater -c config/dev.yml
