.PHONY: run-backends run-lb run-all

run-backends:
	PORT=9001 go run ./fileserver/ &
	PORT=9002 go run ./fileserver/ &
	PORT=9003 go run ./fileserver/

run-lb:
	go run ./loadbalancer/

run-all:
	PORT=9001 go run ./fileserver/ &
	PORT=9002 go run ./fileserver/ &
	PORT=9003 go run ./fileserver/ &
	go run ./loadbalancer/
