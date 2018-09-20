package testing

import (
	"gitlab.com/paybox/hw/mock"
	"testing"
	"gitlab.com/paybox/hw/nd"
	"github.com/magiconair/properties/assert"
)

func mockService() (nds nd.Service, err error) {
	mockND, _ := mock.NewNDMockRepo()
	//ndrepo, _ := hardware.NewNDRepo()
	nds, err = nd.NewService(mockND)
	return nds, nil

}

func Test_open_nd(t *testing.T) {
	nds, err := mockService()
	err = nds.Open("/dev/ttyS0", 9600)
	assert.Equal(t, err, nil)
}

func Test_dispense_nd(t *testing.T) {
	nds, err := mockService()
	err = nds.Dispense(10)
	assert.Equal(t, err, nil)
}

func Test_status_nd(t *testing.T) {
	nds, err := mockService()
	if err != nil {
		t.Fatalf("error mock service ", err.Error())
		return
	}
	id, err := nds.Status()
	assert.Equal(t, id, "M0001")
}
