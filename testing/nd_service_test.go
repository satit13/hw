package testing

import (
	"gitlab.com/paybox/hw/mock"
	"testing"
	"gitlab.com/paybox/hw/nd"
	"github.com/magiconair/properties/assert"
	//"gitlab.com/paybox/hw/hardware"
	//"gitlab.com/paybox/hw/hardware"
	"gitlab.com/paybox/hw/hardware"
)

const port = "/dev/ttyS0"
const speed = 9600

var mode = "mock"

func prepare_nd() (nds nd.Service, err error) {
	var repo nd.Repository
	switch mode {
	case "mock":
		repo, _ = mock.NewNDMockRepo()
	case "hw":
		repo, _ = hardware.NewNDRepo()
	}

	//ndrepo, _ := hardware.NewNDRepo()
	nds, err = nd.NewService(repo)

	err = nds.Open(port, speed)
	//if err != nil {
	//	return nil, err
	//}
	return nds, err
}

func Test_open_nd(t *testing.T) {
	_, err := prepare_nd()

	assert.Equal(t, err, nil)
}

func Test_dispense_nd(t *testing.T) {
	nds, err := prepare_nd()
	err = nds.Dispense(10)
	assert.Equal(t, err, nil)
}

func Test_status_nd(t *testing.T) {
	nds, err := prepare_nd()
	assert.Equal(t, err, nil)
	id, _ := nds.Status()
	assert.Equal(t, id, "M09001")
}

func Test_open_dispense(t *testing.T) {
	nds, err := prepare_nd()
	if err != nil {
		t.Fatalf("error mock service ", err.Error())

	}
	err = nds.Open(port, speed)
	if err != nil {
		t.Fatalf("error mock service ", err.Error())

	}
	err = nds.Dispense(2)
	if err != nil {
		t.Fatalf("error mock service ", err.Error())

	}
}



