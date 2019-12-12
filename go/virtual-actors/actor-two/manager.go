package actor_two

import (
	"log"
	"os"
)

// Settings for SubtractionVirtualActor.
type SubtractionVirtualActorSettings struct {
	QueueDepth int
	Autostart  bool
}

// The SubtractionVirtualActor reference type.
type SubtractionVirtualActor struct {
	RequestRouter chan interface{}
	Running bool

	logger                          *log.Logger
	subtractionVirtualActorListener chan SubtractionVirtualActorRequest
}

func NewSubtractionVirtualActor(settings *SubtractionVirtualActorSettings) (*SubtractionVirtualActor, error) {
	a := &SubtractionVirtualActor{
		RequestRouter:                   make(chan interface{}, settings.QueueDepth),
		logger:                          log.New(os.Stdout, "subtraction-virtual-actor", log.LstdFlags),
		subtractionVirtualActorListener: make(chan SubtractionVirtualActorRequest, settings.QueueDepth),
	}

	if settings.Autostart {
		a.Run()
	}

	return a, nil
}

func (s *SubtractionVirtualActor) deferrer() {
	if r := recover(); r != nil {
		s.logger.Printf("fatality!\n")
	}
}

func (s *SubtractionVirtualActor) Run() {
	go func() {
		defer s.deferrer()

		s.Running = true

		for {
			msg := <-s.RequestRouter
			switch msg.(type) {
			case SubtractionVirtualActorRequest:
				s.subtractionVirtualActorListener <- msg.(SubtractionVirtualActorRequest)
			}
		}
	}()

	s.subtractionRunner()
}

type SubtractionVirtualActorRequest struct {
	X, Y     int
	Response chan *SubtractionVirtualActorResponse
}

type SubtractionVirtualActorResponse struct {
	Result int
	Error  error
	Ok     bool
}

func (s *SubtractionVirtualActor) subtractionRunner() {
	go func() {
		defer s.deferrer()
		for {
			go s.subtract(<-s.subtractionVirtualActorListener)
		}
	}()
}

func (s *SubtractionVirtualActor) subtract(request SubtractionVirtualActorRequest) {
	defer s.deferrer()

	if request.Response != nil {
		request.Response <- &SubtractionVirtualActorResponse{
			Result: request.X - request.Y,
			Error:  nil,
			Ok:     true,
		}
	}
}
