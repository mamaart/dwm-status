deploy:
	go build ./cmd/aiserver 
	sudo mv ./aiserver /usr/bin/aiserver
	sudo systemctl restart aiserver
	go build ./cmd/ask
	sudo mv ./ask /usr/bin/ask
	go build ./cmd/statusbar && sudo setcap cap_net_admin=+ep ./statusbar
	sudo mv ./statusbar /usr/bin/statusbar
	# TODO restart ui here

