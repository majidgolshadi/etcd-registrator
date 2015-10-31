package docker

import (
	"os"
	"log"

	"github.com/fsouza/go-dockerclient"
)

var dockerClient *docker.Client

func init() {
	dockerClient, _ = docker.NewClient(os.Getenv("DOCKER_API"))
}

func EventsListener(dispatcher EventsDispatcher) {
	var err error

	listener := make(chan *docker.APIEvents)
	err = dockerClient.AddEventListener(listener)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {

		err = dockerClient.RemoveEventListener(listener)
		if err != nil {
			log.Fatal(err)
		}

	}()

	for {
		msg := <-listener
		switch msg.Status {
		case "kill":
			dispatcher.Kill(msg)
		case "die":
			dispatcher.Die(msg)
		case "destroy":
			dispatcher.Destroy(msg)
		case "create":
			dispatcher.Create(msg)
		case "restart":
			dispatcher.Restart(msg)
		case "start":
			dispatcher.Start(msg)
		case "stop":
			dispatcher.Stop(msg)
		case "pause":
			dispatcher.Pause(msg)
		case "unpause":
			dispatcher.Unpause(msg)
		}
	}
}
