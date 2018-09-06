package main

import (
	"fmt"
	"github.com/payboxth/vending/host/model"
	"log"
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
	fmt.Printf("BA GetEnable -> %v \n", BA.Enable)
	fmt.Println("BA status : ", BA.Status)

	errorRetry := 0
	noteCount := 1
	noteMaxCount := int(1)

	lastNoteIndex, _, _ := BA.OnBillReceived()
	fmt.Println("lastNoteIndex Before for loop , ", lastNoteIndex)

	for {
		noteIndex, noteData, err := BA.OnBillReceived()
		fmt.Println("loop for wait to payment ...  ")
		if err != nil {
			fmt.Println("error ", err.Error())
			errorRetry++
			continue
		}
		errorRetry = 0

		if noteIndex != lastNoteIndex && noteData[1] == 1 {
			BA.TakePendingBill()
			fmt.Println("Received BA  Count ", noteCount)
			noteCount++
		}
		lastNoteIndex = noteIndex
		if noteCount > noteMaxCount {
			break
		}
	}
}
