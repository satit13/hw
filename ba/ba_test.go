package ba_test

import (
	"bytes"
	"testing"
	"gitlab.com/paybox/hw/ba"
)

func BillAcceptorInit() (ba ba.BillAcceptor, err error) {
	cctalk := &ba.CCTalkDriver
	err = cctalk.OpenConnection("/dev/ttyUSB0", 9600)
	if err != nil {
		return
	}
	go cctalk.StartPolling()
	ba = ba.BillAcceptor{DeviceID: 0x28}
	ba.CCTalkDriver = *cctalk
	return
}

func TestBillAcceptorIsOnline(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	if online, _ := ba.IsOnline(); !online {
		t.Fatalf("Expected online but got offline")
	}
}

func TestBillAcceptorGetSerialNumberMustEqual166138(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := 166138
	a := ba.GetSerialNumber()
	if e != a {
		t.Fatalf("Expected serial number %d but got %d", e, a)
	}
}

func TestBillAcceptorSetInhibitFalseMustGetInhibitFalse(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := false
	ba.SetInhibit(false, 255, 255)
	ba.SetMasterInhibit(false)
	a, err := ba.GetInhibit()
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestBillAcceptorSetInhibitTrueMustGetInhibitTrue(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := true
	ba.SetInhibit(true, 0, 0)
	ba.SetMasterInhibit(true)
	a, err := ba.GetInhibit()
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill1MustGotTH0020B(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH0020B"
	a, err := ba.GetBillId(1)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill2MustGotTH0050B(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH0050B"
	a, err := ba.GetBillId(2)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill3MustGotTH0100B(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH0100B"
	a, err := ba.GetBillId(3)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill4MustGotTH0500B(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH0500B"
	a, err := ba.GetBillId(4)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill5MustGotTH0500A(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH0500A"
	a, err := ba.GetBillId(5)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill6MustGotTH1000B(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH1000B"
	a, err := ba.GetBillId(6)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestBillAcceptorGetBillIdForBill7MustGotTH1000A(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := "TH1000A"
	a, err := ba.GetBillId(7)
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if e != a {
		t.Fatalf("Expected bill id %s but got %s", e, a)
	}
}

func TestReceiveNote20BathMustGetBillType9(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := []byte{9}
	ba.ResetDevice()
	status := ba.PerformSelfCheck()
	for status != 0 {
		status = ba.PerformSelfCheck()
	}
	ba.SetInhibit(false, 255, 255)
	ba.SetMasterInhibit(false)
	count, bills, err := ba.OnBillReceived()
	for bills[0] == 0 {
		count, bills, err = ba.OnBillReceived()
		if bytes.Contains(e, []byte{bills[0]}) {
			ba.TakePendingBill()
		}
	}
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if count > 0 && bytes.IndexByte(e, bills[0]) == -1 {
		t.Fatalf("Expected %v but got %v", e, bills[0])
	}
}

func TestReceiveNote50BathMustGetBillType8(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := []byte{8}
	ba.ResetDevice()
	status := ba.PerformSelfCheck()
	for status != 0 {
		status = ba.PerformSelfCheck()
	}
	ba.SetInhibit(false, 255, 255)
	ba.SetMasterInhibit(false)
	count, bills, err := ba.OnBillReceived()
	for bills[0] == 0 {
		count, bills, err = ba.OnBillReceived()
		if bytes.Contains(e, []byte{bills[0]}) {
			ba.TakePendingBill()
		}
	}
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if count > 0 && bytes.IndexByte(e, bills[0]) == -1 {
		t.Fatalf("Expected %v but got %v", e, bills[0])
	}
}

func TestReceiveNote100BathMustGetBillType3or11(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := []byte{3, 11}
	ba.ResetDevice()
	status := ba.PerformSelfCheck()
	for status != 0 {
		status = ba.PerformSelfCheck()
	}
	ba.SetInhibit(false, 255, 255)
	ba.SetMasterInhibit(false)
	count, bills, err := ba.OnBillReceived()
	for bills[0] == 0 {
		count, bills, err = ba.OnBillReceived()
		if bytes.Contains(e, []byte{bills[0]}) {
			ba.TakePendingBill()
		}
	}
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if count > 0 && bytes.IndexByte(e, bills[0]) == -1 {
		t.Fatalf("Expected %v but got %v", e, bills[0])
	}
}

func TestReceiveNote500BathMustGetBillType4(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := []byte{4}
	ba.ResetDevice()
	status := ba.PerformSelfCheck()
	for status != 0 {
		status = ba.PerformSelfCheck()
	}
	ba.SetInhibit(false, 255, 255)
	ba.SetMasterInhibit(false)
	count, bills, err := ba.OnBillReceived()
	for bills[0] == 0 {
		count, bills, err = ba.OnBillReceived()
		if bytes.Contains(e, []byte{bills[0]}) {
			ba.TakePendingBill()
		}
	}
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if count > 0 && bytes.IndexByte(e, bills[0]) == -1 {
		t.Fatalf("Expected %v but got %v", e, bills[0])
	}
}

func TestReceiveNote1000BathMustGetBillType6(t *testing.T) {
	ba, err := BillAcceptorInit()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer ba.Close()
	e := []byte{6, 12}
	ba.ResetDevice()
	status := ba.PerformSelfCheck()
	for status != 0 {
		status = ba.PerformSelfCheck()
	}
	ba.SetInhibit(false, 255, 255)
	ba.SetMasterInhibit(false)
	count, bills, err := ba.OnBillReceived()
	for bills[0] == 0 {
		count, bills, err = ba.OnBillReceived()
		if bytes.Contains(e, []byte{bills[0]}) {
			ba.TakePendingBill()
		}
	}
	if err != nil {
		t.Fatalf("Expected %v but got error", e)
	} else if count > 0 && bytes.IndexByte(e, bills[0]) == -1 {
		t.Fatalf("Expected %v but got %v", e, bills[0])
	}
}

