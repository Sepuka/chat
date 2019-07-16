CONFIG_PATH=config/dev.yml.dist

init:
	dep ensure -v
	cp -n $(CONFIG_PATH) config/dev.yml

start:
	go build -o chat && ./chat repeater -c config/dev.yml
