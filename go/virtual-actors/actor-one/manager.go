package actor_two

import (
	"log"
	"os"
)

// Settings for AdditionVirtualActor.
type AdditionVirtualActorSettings struct {
	QueueDepth int
	Autostart  bool
}

// The AdditionVirtualActor reference type.
type AdditionVirtualActor struct {
	RequestRouter chan interface{}
	Running bool

	logger                       *log.Logger
	additionVirtualActorListener chan AdditionVirtualActorRequest
}

func NewAdditionVirtualActor(settings *AdditionVirtualActorSettings) (*AdditionVirtualActor, error) {
	a := &AdditionVirtualActor{
		RequestRouter:                make(chan interface{}, settings.QueueDepth),
		logger:                       log.New(os.Stdout, "addition-virtual-actor ", log.LstdFlags),
		additionVirtualActorListener: make(chan AdditionVirtualActorRequest, settings.QueueDepth),
	}

	if settings.Autostart {
		a.Run()
	}

	return a, nil
}

func (a *AdditionVirtualActor) deferrer() {
	if r := recover(); r != nil {
		a.logger.Printf("fatality!\n")
	}
}

func (a *AdditionVirtualActor) Run() {
	go func() {
		defer a.deferrer()

		a.Running = true

		for {
			msg := <-a.RequestRouter
			switch msg.(type) {
			case AdditionVirtualActorRequest:
				a.additionVirtualActorListener <- msg.(AdditionVirtualActorRequest)
			}
		}
	}()

	a.additionRunner()
}

type AdditionVirtualActorRequest struct {
	X, Y     int
	Response chan *AdditionVirtualActorResponse
}

type AdditionVirtualActorResponse struct {
	Result int
	Error  error
	Ok     bool
}

func (a *AdditionVirtualActor) additionRunner() {
	go func() {
		defer a.deferrer()
		for {
			msg := <- a.additionVirtualActorListener
			go a.add(msg)
		}
	}()
}

func (a *AdditionVirtualActor) add(request AdditionVirtualActorRequest) {
	defer a.deferrer()

	if request.Response != nil {
		request.Response <- &AdditionVirtualActorResponse{
			Result: request.X + request.Y,
			Error:  nil,
			Ok:     true,
		}
	}
}
