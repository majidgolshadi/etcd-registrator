package docker
import "github.com/fsouza/go-dockerclient"

type EventsDispatcher interface {
	Kill(docker.APIEvents)
	Die(docker.APIEvents)
	Destroy(docker.APIEvents)
	Create(docker.APIEvents)
	Start(docker.APIEvents)
	Stop(docker.APIEvents)
	Restart(docker.APIEvents)
	Pause(docker.APIEvents)
	Unpause(docker.APIEvents)
}