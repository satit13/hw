package t_test

import (
	"testing"
	"github.com/payboxth/vending/host/hw/cctalk"
	"fmt"
	//"log"
	"github.com/payboxth/vending/host/model"
	"github.com/magiconair/properties/assert"
)

func Test_CcTalk_Open(t *testing.T) {
	CCT := &cctalk.Driver{}
	// try to connect CCT Port
	err := CCT.OpenConnection("/dev/ttyCCTALK", 9600)
	defer CCT.Close()
	assert.Equal(t, err.Error(), "open /dev/ttyCCTALK: no such file or directory", "errror @!!!")
}

func Test_BA_open(t *testing.T) {
	CCT := &cctalk.Driver{}
	// try to connect CCT Port
	err := CCT.OpenConnection("/dev/ttyCCTALK", 9600)
	defer CCT.Close()
	if err != nil {
		//t.Errorf("expect cctalk connect but --> %v", err.Error())
		assert.Matches(t, err.Error(), "open /dev/ttyCCTALK: no such file or directory")
	}

	BA := model.BillAcceptor{
		DeviceID:       0x28,
		Status:         "idle",
		Send:           make(chan *model.Message, 10),
		//TakeRejectNote: make(chan bool),
		EscrowValue:    make(chan float64, 1),
		StackedValue:   make(chan float64, 1),
	}

	if CCT.Online && !CCT.Busy {
		fmt.Println(" try BA.Open(*CCT)...")
		err := BA.Open(*CCT)
		if err != nil {
			t.Errorf("expect BA connnected true but actual is ", err.Error())
			fmt.Println("BA Open Error : ", err.Error())
			assert.Matches(t, err.Error(), "open /dev/ttyCCTALK: no such file or directory")
		}
		defer BA.Close()
	}

}
