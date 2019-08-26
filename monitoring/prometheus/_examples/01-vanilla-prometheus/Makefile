# Docker Swarm - Vanilla Prometheus
# 

# Initialise Docker Swarm
init:
	docker swarm init

# Start the 'vanilla-prometheus' stack.
start:
	docker stack deploy vanilla-prometheus --compose-file=./vanilla-prometheus.yaml

# Get the status of the 'vanilla-prometheus' stack
status:
	docker stack ls
	docker stack services prometheus-basic
	docker stack ps prometheus-basic

# Start the 'vanilla-prometheus' stack.
stop:
	docker stack rm vanilla-prometheus

open:
	x-www-browser http://localhost:9090/

# Clean-up.
clean: stop
	docker swarm leave --force
