package main

import (
	"fmt"
	"github.com/payboxth/vending/host/model"
	"log"
	//"time"
	//"time"
)

func main() {

	CCT := &model.CCTalkDriver{}

	// try to connect CCT Port
	err := CCT.OpenConnection("/dev/ttyCCTALK", 9600)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("CCT Connected...")
		go CCT.StartPolling()
	}

	BA := model.BillAcceptor{
		DeviceID:       0x28,
		Status:         "idle",
		Send:           make(chan *model.Message, 10),
		TakeRejectNote: make(chan bool),
		EscrowValue:    make(chan float64, 1),
		StackedValue:   make(chan float64, 1),
	}
	fmt.Println("status of CCT.Online ", CCT.Online)
	//  if connectd CCT to port
	//  then Open BA
	if CCT.Online && !CCT.Busy {
		fmt.Println(" try BA.Open(*CCT)...")
		BA.Open(*CCT)
	}

	// ---- TEST --
	fmt.Println(" สั่งเปิด Test รอรับ ธนบัตร .... Start")
	BA.SetEnable(true)
	BA.SetInhibit4(false, 1, 1, 1, 1)
	fmt.Printf("BA GetEnable -> %v \n", BA.Enable)
	fmt.Println("BA status : ", BA.Status)

	errorRetry := 0
	noteCount := 0
	noteMaxCount := int(2)
	lastNoteIndex, _, _ := BA.OnBillReceived()
	//fmt.Println("lastNoteIndex Before for loop , ", lastNoteIndex)

	for i := 1; i <= 18; i++ {
		bill, err := BA.GetBillId(byte(i))
		if err == nil {
			fmt.Printf("Bill type %d : %s \n", i, bill)
		}
	}

	for {
		noteIndex, noteData, err := BA.OnBillReceived()
		fmt.Printf("Recived payment ...  noteIndex - %v, noteData - %v", noteIndex, noteData)
		if err != nil {
			fmt.Println("error ", err.Error())
			errorRetry++
			continue
		}
		errorRetry = 0
		//time.Sleep(3000)
		if noteIndex != lastNoteIndex && noteData[1] == 1 {
			BA.TakePendingBill()
			noteCount++
			fmt.Println("Received BA  Count ", noteCount)
		}
		lastNoteIndex = noteIndex
		if noteCount >= noteMaxCount {
			break
		}
	}
}
