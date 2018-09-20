package cctalk

import (
	"github.com/tarm/serial"
	"gopkg.in/oleiade/lane.v1"
	"time"
	"log"
	"errors"
	"fmt"
)

type Driver struct {
	serial.Port
	Busy             bool
	Online           bool
	RequestPool      *lane.Queue
	SerialPortConfig *serial.Config
	Polling          bool
}

type Command struct {
	Request  []byte
	Response chan []byte
}

const (
	SimplePoll = 254
	AddressPoll = 253
	RequestStatus = 248
	RequestManId = 246
	RequestCatId = 245
	RequestProductCode = 244
	RequestDatabaseVersion = 243
	RequestSerial = 242
	RequestSwRev = 241
	TestSolenoids = 240
	OperateMotors = 239
	PerformSelfCheck = 232
	ModifyInhibit = 231
	RequestInhibit = 230
	RequestBufferedCredit = 229
	ModifyMasterInhibit = 228
	RequestMasterInhibit = 227
	RequestInsertionCounter = 226
	RequestAcceptCounter = 225
	CalculateRomChecksum = 197
	RequestRejectCounter = 194
	RequestFraudCounter = 193
	ModifyCoinId = 185
	RequestCoinId = 184
	RequestBufferedBillEvents = 159
	ModifyBillId = 158
	RequestBillId = 157
	RouteBill = 154
	RequestCurrencyRevision = 145
	UploadBillTables = 144
	BeginBillTableUpgrade = 143
	FinishBillTableUpgrade = 142
	RequestFirmwareUpgradeCapability = 141
	UploadFirmware = 140
	BeginFirmwareUpgrade = 139
	FinishFirmwareUpgrade = 138
	SmartEmpty = 51
	SetInhibitPeripheralDeviceValue = 50
	SetPeripheralDeviceMasterInhibit = 49
	RequestStatusCurrency = 47
	FloatByDenominationCurrency = 45
	PayoutByDenominationCurrency = 44
	SetDenominationAmountCurrency = 43
	GetDenominationAmountCurrency = 42
	PayoutAmountCurrency = 39
	FloatByDenomination = 33
	PayoutByDenomination = 32
	RequestHoperStatus = 29
	SetDenominationAmount = 27
	GetDenominationAmount = 26
	GetMinimumPayout = 25
	Empty = 24
	PayoutAmount = 22
	ResetDevice = 1
)

// OpenConnection เป็นเมธอดเปิดพอร์ตอนุกรมที่ใช้สื่อสารกับอุปกรณ์ ccTalk


// StartPolling เป็นเมธอดคอยตรวจสอบว่ามีการหยอดเหรียญหรือใส่ธนบัตรเข้ามาหรือไม่ เหรียญอะไร ธนบัตรอะไร
func (cctalk *Driver) StartPolling() {
	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in ccTalk.StartPolling", r)
		}
		cctalk.Polling = false
	}()
	fmt.Println("CCT.StartPolling...", cctalk.RequestPool.Size())
	cctalk.Polling = true
	for {
		if cctalk.RequestPool.Head() != nil {
			ccTalkCmd := cctalk.RequestPool.Dequeue()
			ccTalkRequest := ccTalkCmd.(*Command)

			n, err := cctalk.Write(ccTalkRequest.Request)
			if err != nil {
				cctalk.Close()        // Close cctalk serial port.
				cctalk.Online = false // Set cctalk offline.
				log.Println(err)
				ccTalkRequest.Response <- []byte{}
				continue
			}

			buf := make([]byte, 256)
			response := []byte{}
			for {
				n, _ = cctalk.Read(buf)
				response = append(response, buf[:n]...)
				if n == 0 {
					break
				}
			}
			ccTalkRequest.Response <- response
		} else {
			time.Sleep(time.Millisecond * 50)
		}
	}
}