package hardware

import (
	"gitlab.com/paybox/hw/nd"
	"time"
	"fmt"
	"github.com/tarm/serial"
	"log"

	"gitlab.com/paybox/hw/util"
)


func NewNDRepo() (nd.Repository, error) {
	nd := repo{}
	return &nd, nil
}

type repo struct {
	nd.NoteDispenser
}

func (nd *repo) Connect(portName string, speed int) (err error) {
	serialPortConfig := &serial.Config{Name: portName, Baud: 9600, Parity: serial.ParityEven, ReadTimeout: time.Millisecond * 100}
	serialPort, err := serial.OpenPort(serialPortConfig)
	if err != nil {
		log.Println(err)
		return err
	}
	nd.Port = *serialPort

	//s, d, err := nd.RequestMachineStatus()
	//if err != nil {
	//	log.Printf("NoteDispenser status error : status=%v, data=%v, error=%s ", s, d, err)
	//	return err
	//}
	//err = nd.ResetDispenser()
	//if err != nil {
	//	log.Println("NoteDispenser reset error : ", err)
	//	return err
	//}
	nd.Online = true
	return nil
}

func (nd *repo) ResetDispenser() error {
	ict_reset_dispenser_cmd := []byte{0x1, 0x10, nd.DeviceId, 0x12, 0x0}
	checksum := util.CheckSum8Modulo256(ict_reset_dispenser_cmd)
	dataWithChecksum := append(ict_reset_dispenser_cmd, checksum)
	log.Printf("Request resetDispenser : % X\n", dataWithChecksum)
	_, err := nd.Write(dataWithChecksum)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (nd *repo) RequestMachineStatus() (status byte, data byte, err error) {
	ictRequestMachineStatusCmd := []byte{0x1, 0x10, nd.DeviceId, 0x11, 0x0}
	checksum := util.CheckSum8Modulo256(ictRequestMachineStatusCmd)
	dataWithChecksum := append(ictRequestMachineStatusCmd, checksum)
	var response []byte
	n, err := nd.Write(dataWithChecksum)
	if err != nil {
		return 255, 0, err
	}

	time.Sleep(time.Millisecond * 100)
	buf := make([]byte, 6)
	response = []byte{}
	for {
		n, _ = nd.Read(buf)
		if n == 0 {
			break
		}
		response = append(response, buf[:n]...)
	}
	log.Printf("Response requestMachineStatus : % X\n", response)
	if len(response) < 6 {
		return 255, 0, fmt.Errorf("data length incorrect")
	} else {
		for len(response) >= 6 {
			if util.CheckSum8Modulo256(response[:5]) != uint8(response[5]) {
				status, data, err = 255, 0, fmt.Errorf(fmt.Sprintf("checksum must %X got %X.\n", util.CheckSum8Modulo256(response[:5]), uint8(response[5])))
			} else {
				status, data, err = response[3], response[4], nil
			}
			response = response[6:]
		}
	}
	return
}

func (nd *repo) PayNote(noteCount byte) error {
	ict_dispense_note_cmd := []byte{0x1, 0x10, nd.DeviceId, 0x10, noteCount}
	checksum := util.CheckSum8Modulo256(ict_dispense_note_cmd)
	dataWithChecksum := append(ict_dispense_note_cmd, checksum)
	log.Printf("Request dispenseNote : % X\n", dataWithChecksum)

	_, err := nd.Write(dataWithChecksum)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func (nd *repo) GetStatus() (string, error) {
	return "ND001", nil
}