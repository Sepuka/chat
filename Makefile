init:
	dep ensure -v

start:
	go build -o chat && ./chat repeater -c config/dev.yml
