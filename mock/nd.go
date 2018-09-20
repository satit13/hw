package mock

import (
	"gitlab.com/paybox/hw/nd"
)

func NewNDMockRepo() (nd.Repository, error) {
	nd := ndrepo{}
	return &nd, nil
}

type ndrepo struct {
	nd.NoteDispenser
}

func (nd *ndrepo) Connect(portName string, speed int) (err error) {
	return nil
}

func (nd *ndrepo) ResetDispenser() error {
	return nil
}

func (nd *ndrepo) RequestMachineStatus() (status byte, data byte, err error) {

	return 255, 0, nil
}

func (nd *ndrepo) PayNote(qty int) error {
	return nil
}

func (nd *ndrepo) GetStatus() (string, error) {
	return " id from Mock  M0001", nil
}