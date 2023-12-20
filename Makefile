deploy:
	go build ./cmd/statusbar && sudo setcap cap_net_admin=+ep ./statusbar
