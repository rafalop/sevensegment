package sevensegment

import (
	"fmt"
	"log"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/host"
)

// Sevensegment with a buffer to hold desired
// segment state
type SevenSegment struct {
	d      *i2c.Dev
	Buffer [10]uint16
}

func NewSevenSegment(addr byte) (ss *SevenSegment) {
	ss = new(SevenSegment)
	host.Init()
	if b, err := i2creg.Open(""); err != nil {
		log.Fatal(err)
		return
	} else {
		if b.SetSpeed(100 * physic.KiloHertz); err != nil {
			log.Fatal(err)
			return
		}
		ss.d = &i2c.Dev{Addr: uint16(addr), Bus: b}
		ss.OscillatorOn()
		ss.DisplayOn()
		ss.SetBrightness(4)
	}
	return
}

// Write raw data to bus, useful for setting up
// the seven segment
func (ss *SevenSegment) WriteRaw(data []byte) {
	read := make([]byte, 1)
	if err := ss.d.Tx(data, read); err != nil {
		log.Fatal(err)
	} else {
	}
}

// Set brightness from 0-15
func (ss *SevenSegment) SetBrightness(brightness int) {
	//Command for brightness is 0xe + "dimming" level 0x0-0xf
	cmd := []byte{0xe0 + byte(brightness)}
	ss.WriteRaw(cmd)
}

// Command to turn on display is 0x81
func (ss *SevenSegment) DisplayOn() {
	cmd := []byte{0x81}
	ss.WriteRaw(cmd)
}

// Command to turn on oscillator
func (ss *SevenSegment) OscillatorOn() {
	cmd := []byte{0x21}
	ss.WriteRaw(cmd)
}

//Write out my buffer, this will actually display segments
func (ss *SevenSegment) WriteData() {
	read := make([]byte, 1)
	data := make([]byte, len(ss.Buffer))
	i := 0
	for _, item := range ss.Buffer {
		data[i] = byte(item)
		i++
	}
	//We try writing the data a few times
	for j := 0; j < 3; j++ {
		if err := ss.d.Tx(data, read); err != nil {
			log.Println("Error writing data, trying again...")
		} else {
			break
			//fmt.Println("Buffer:", ss.Buffer, "Wrote data: ", data)
		}
	}
}

// Clear all segments
func (ss *SevenSegment) Clear() {
	for i, _ := range ss.Buffer {
		ss.Buffer[i] = uint16(0)
	}
	ss.WriteData()
}
