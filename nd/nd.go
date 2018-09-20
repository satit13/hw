package nd

import (
	"github.com/tarm/serial"
)

type Repository interface {
	Connect(string, int) error
	PayNote(int) error
	GetStatus() (string, error)
}

type Message struct {
	Device  string      `json:"device"`
	Type    string      `json:"type"`
	Command string      `json:"command"`
	Result  bool        `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Time    int         `json:"time,omitempty"`
	Size    int         `json:"size,omitempty"`
	Color   string      `json:"color,omitempty"`
}
type NoteDispenser struct {
	MachineId       string `json:"machine_id"`
	DeviceId        byte // Default 0x00
	Enable          bool
	Online          bool
	Remain          int
	Dispensed       int
	ChangeNoteValue int
	Send            chan *Message
	serial.Port
}

var NoteDispenserStatus = map[int]string{
	0xAA: "Payout successful",
	0xBB: "Payout fails",
	0:    "Status fine",
	1:    "Empty note",
	2:    "Stock less",
	3:    "Note jam",
	4:    "Over length",
	5:    "Note not exit",
	6:    "Sensor error",
	7:    "Double note error",
	8:    "Motor error",
	9:    "Dispensing busy",
	10:   "Sensor adjusting",
	11:   "Checksum error",
	12:   "Low power error",
}
