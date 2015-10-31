package main
import "github.com/majidgolshadi/registrator/docker"

func main() {
	// connect to docker event api
	docker.EventsListener()
}
