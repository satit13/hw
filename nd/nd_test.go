package nd_test

import (
	"testing"
	"github.com/payboxth/vending/host/model"
	"github.com/payboxth/vending/host/hw"
	"log"
	"github.com/magiconair/properties/assert"
)

func ND_Init() error {
	ND := &model.NoteDispenser{
		DeviceId:        0,
		ChangeNoteValue: 20,
	}
	err := ND.Open("/dev/ttyS0")
	if err != nil {
		if err.Error() == hw.ND_WIRING_WRONG_CABLE_WITH_OTHER_DEVICE {
			log.Println("สงสัยเสียบสายอนุกรมสลับกับเครื่องสแกนบาร์โค้ด")
		}
		log.Println("Unknown Error:", err)
		return err
	}
	return nil
}

func Test_nd_NoDevice_Should_Fail(t *testing.T) {
	err := ND_Init()
	assert.Equal(t, err.Error(), "open /dev/ttyS0: permission denied", "")
}
