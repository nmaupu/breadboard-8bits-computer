package main

import (
	"log"
	"os"
)

const (
	OutputFilename = "output.bin"
	EepromSize     = 8192
)

var (
	// Digits are coded to display the correct segments on a 7 segments display such as:
	// https://docs.broadcom.com/doc/AV02-1363EN
	// Pinout are as follow:
	// 3 -
	// 2| |4
	// 1 -
	// 5| |7
	// 6 - .8
	Digits = []byte{
		0b01111110, // 0
		0b00010010, // 1
		0b10111100, // 2
		0b10110110, // 3
		0b11010010, // 4
		0b11100110, // 5
		0b11101110, // 6
		0b00110010, // 7
		0b11111110, // 8
		0b11110110, // 9
		0b11111010, // a
		0b11001110, // b
		0b01101100, // c
		0b10011110, // d
		0b11101100, // e
		0b11101000, // f
		0b10000000, // -
		0b11111111, // all on
		0b00000000, // all off
	}
)

const (
	DisplayAllOff = 18
	DisplayAllOn  = 17
	DisplayMinus  = 16
)

// Generates a file ready to be burnt on an EEPROM (64K - 8K x 8 like the AT28C64B)
// https://www.mouser.fr/datasheet/2/268/doc0270-1108115.pdf
func main() {
	var err error

	file, err := os.OpenFile(OutputFilename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Cannot open file %s", OutputFilename)
	}
	defer file.Close()

	// Writing units, tens and hundreds 0 ... 255
	addr := 0
	for i := 0; i <= 0xff; i++ {
		units := i % 10
		tens := (i / 10) % 10
		hundreds := (i / 100) % 10
		writeByteToFile(file, Digits[units], int64(addr))
		writeByteToFile(file, Digits[tens], int64(addr+256))
		writeByteToFile(file, Digits[hundreds], int64(addr+512))
		writeByteToFile(file, Digits[DisplayAllOff], int64(addr+768))
		addr++
	}

	// Writing units, tens and hundreds -128 ... 127
	addr += 768
	for i := -128; i <= 127; i++ {
		iAbs := Abs(i)
		units := iAbs % 10
		tens := (iAbs / 10) % 10
		hundreds := (iAbs / 100) % 10
		a := int64(byte(i)) + int64(addr)
		writeByteToFile(file, Digits[units], int64(a))
		writeByteToFile(file, Digits[tens], int64(a+256))
		writeByteToFile(file, Digits[hundreds], int64(a+512))
		if i < 0 {
			writeByteToFile(file, Digits[DisplayMinus], int64(a+768))
		} else {
			writeByteToFile(file, Digits[DisplayAllOff], int64(a+768))
		}
	}

	// fill in the rest of the file to fit eeprom size
	addr += 1024
	for i := addr; i < EepromSize; i++ {
		writeByteToFile(file, Digits[DisplayAllOn], int64(i))
	}
}

func writeByteToFile(f *os.File, b byte, off int64) {
	n, err := f.WriteAt([]byte{b}, off)
	if n != 1 || err != nil {
		log.Printf("Error writing byte %08b (0x%02x) at address %02x\n", b, b, off)
	}
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}
