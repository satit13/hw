package t_test

import (
	"testing"
	"github.com/payboxth/vending/host/hw/cctalk"
	"fmt"
	"log"
	"github.com/payboxth/vending/host/model"
)

func Test_CcTalk_Connect(t *testing.T) {
	CCT := &cctalk.Driver{}
	// try to connect CCT Port
	err := CCT.OpenConnection("/dev/ttyCCTALK", 9600)
	if err != nil {
		log.Println("Error OpenConnection ccTalk state : ", err)
		t.Errorf("expect cctalk connnect true but actual is %s", err.Error())
		return
	} else {
		fmt.Println("CCT Connected...")
		CCT.Close()
		return
	}

}

func Test_BA_Connect(t *testing.T) {
	CCT := &cctalk.Driver{}
	// try to connect CCT Port
	err := CCT.OpenConnection("/dev/ttyCCTALK", 9600)
	if err != nil {
		t.Errorf("expect cctalk connect but --> %v", err.Error())
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
			BA.Close()
			return
		}
	}
}

func Test_BA_Open(t *testing.T) {
	CCT := &cctalk.Driver{}
	// try to connect CCT Port
	err := CCT.OpenConnection("/dev/ttyCCTALK", 9600)
	if err != nil {
		t.Errorf("expect cctalk connect but --> %v", err.Error())
	}
	CCT.StartPolling()
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
			return
		}
	}

	BA.SetEnable(true)
	BA.Close()
	return

}