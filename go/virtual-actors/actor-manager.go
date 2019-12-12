package main

import (
	"fmt"
	"log"
	"os"
	"time"

	adder "./actor-one"
	subtractor "./actor-two"
)

type VirtualActorRequestType string

const (
	AdditionRequest    VirtualActorRequestType = "add"
	SubtractionRequest VirtualActorRequestType = "subtract"
)

type VirtualActorManagerSettings struct {
	QueueDepth                      int
	Autostart                       bool
	AdditionVirtualActorSettings    *adder.AdditionVirtualActorSettings
	SubtractionVirtualActorSettings *subtractor.SubtractionVirtualActorSettings
}

type VirtualActorManager struct {
	RequestRouter chan VirtualActorRequest
	Running bool

	logger                  *log.Logger
	additionVirtualActor    *adder.AdditionVirtualActor
	subtractionVirtualActor *subtractor.SubtractionVirtualActor
}

func NewVirtualActorManager(settings *VirtualActorManagerSettings) (*VirtualActorManager, error) {
	v := &VirtualActorManager{
		RequestRouter: make(chan VirtualActorRequest, settings.QueueDepth),
		logger:        log.New(os.Stdout, "virtual-actor-manager ", log.LstdFlags),
	}

	var err error

	v.additionVirtualActor, err = adder.NewAdditionVirtualActor(settings.AdditionVirtualActorSettings)
	if err != nil {
		return &VirtualActorManager{}, err
	}

	v.subtractionVirtualActor, err = subtractor.NewSubtractionVirtualActor(settings.SubtractionVirtualActorSettings)
	if err != nil {
		return &VirtualActorManager{}, err
	}

	if settings.Autostart {
		v.Run()
	}

	return v, nil
}

func (v *VirtualActorManager) deferrer() {
	if r := recover(); r != nil {
		v.logger.Printf("fatality!\n")
	}
}

func (v *VirtualActorManager) Run() {
	go func() {
		defer v.deferrer()

		v.Running = true

		for {
			msg := <-v.RequestRouter

			v.logger.Printf("received %s", msg.Type)

			switch msg.Type {
			case AdditionRequest:
				v.additionVirtualActor.RequestRouter <- msg.Payload
			case SubtractionRequest:
				v.subtractionVirtualActor.RequestRouter <- msg.Payload
			}
		}
	}()
}

type VirtualActorRequest struct {
	Type    VirtualActorRequestType
	Payload interface{}
}

func main() {

	autoStart := true
	queueDepth := 10

	vaOpts := &VirtualActorManagerSettings{
		QueueDepth: 10,
		Autostart:  autoStart,
		AdditionVirtualActorSettings: &adder.AdditionVirtualActorSettings{
			QueueDepth: queueDepth,
			Autostart:  autoStart,
		},
		SubtractionVirtualActorSettings: &subtractor.SubtractionVirtualActorSettings{
			QueueDepth: queueDepth,
			Autostart:  autoStart,
		},
	}

	v, err := NewVirtualActorManager(vaOpts)
	if err != nil {
		panic(err)
	}

	// v.Run() is not needed because we're autostarting the virtual actor pool.

	// prepare the response channel.
	arespc := make(chan *adder.AdditionVirtualActorResponse, 1)

	// send the request.
	x, y := 1, 1
	v.RequestRouter <- VirtualActorRequest{
		Type: AdditionRequest,
		Payload: adder.AdditionVirtualActorRequest{
			X:        x,
			Y:        y,
			Response: arespc,
		},
	}

	// other things can be done now since we'll just read from the channel when we need the response.
	time.Sleep(time.Second * 3)

	// read the response.
	aresp := <-arespc
	close(arespc)

	if !aresp.Ok {
		fmt.Println(fmt.Errorf("adder error: %s", aresp.Error))
	}

	fmt.Printf("added %d + %d = %d\n", x, y, aresp.Result)

	// prep, make, and wait for the response.
	srespc := make(chan *subtractor.SubtractionVirtualActorResponse, 1)
	x, y = 1, 1
	v.RequestRouter <- VirtualActorRequest{
		Type: SubtractionRequest,
		Payload: subtractor.SubtractionVirtualActorRequest{
			X:        x,
			Y:        y,
			Response: srespc,
		},
	}
	sresp := <-srespc
	close(srespc)

	if !sresp.Ok {
		fmt.Println(fmt.Errorf("subtractor error: %s", sresp.Error))
	}

	fmt.Printf("subtracted %d - %d = %d\n", x, y, sresp.Result)
}
