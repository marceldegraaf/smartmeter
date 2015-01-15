package poller

import (
	"bufio"
	"io"

	"github.com/howeyc/crc16"
	"github.com/marceldegraaf/smartmeter/log"
	"github.com/marceldegraaf/smartmeter/types"
	"github.com/tarm/goserial"
)

const (
	device           = "/dev/ttyUSB0"
	rate             = 115200
	messageDelimiter = '\x21' // "!" character
	crcDelimiter     = '\x0a' // newline
	crcPolynomial    = 0xA001 // IBM CRC16
)

var (
	config   *serial.Config
	usb      io.ReadWriteCloser
	reader   *bufio.Reader
	Incoming = make(chan types.Telegram, 8)
)

func Initialize() error {
	var err error

	config = &serial.Config{Name: device, Baud: rate}

	usb, err = serial.OpenPort(config)
	if err != nil {
		log.Errorf("Could not open serial interface: %s", err)
		return err
	}

	reader = bufio.NewReader(usb)

	return nil
}

func Poll() {
	for {
		payload := blockingReadFromSerial(messageDelimiter)
		log.Debugf("Received payload: %s", payload)

		Incoming <- types.Telegram{Payload: payload}
	}
}

// Blocks until delimiter is received, returns the buffer
// including delimiter
func blockingReadFromSerial(delimiter byte) string {
	message, err := reader.ReadString(delimiter)
	if err != nil {
		log.Errorf("Error while reading from serial interface: %s", err)
		return ""
	}

	return message
}

func calculateCRC(message []byte) uint16 {
	table := crc16.MakeTable(crcPolynomial)

	return crc16.Checksum(message, table)
}
