package main

import (
	"log"
	"os"
)

const (
	OutputFilename = "output.bin"
	EepromSize     = 8192
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

	// Create a zero file first to have eeprom file matching the eeprom size
	for i := 0; i < EepromSize; i++ {
		WriteByteToFile(file, 0, int64(i))
	}

	// Address is given using:
	// Intstruction (4 bits) + step (3 bits)
	// First bit (IO7) is high to address the eeprom with the 8 MSB of the control word (left eeprom)
	// and low to get the 8 LSB from the other eeprom (right eeprom)
	// So both eeproms have the exact same microcode
	// a9 a8  a7  a6 a5 a4 a3  a2 a1 a0
	// FC FZ  BS  I3 I2 I1 I0  S2 S1 S0

	for addr := 0; addr <= 0b1111111111; addr++ {
		step := (addr & StepBits)
		if step >= 6 {
			continue
		}
		inst := (addr & InstructionBits) >> 3
		bs := (addr & SelectBit) >> 7
		carryFlag := (addr & FlagCarryBit) >> 8
		zeroFlag := (addr & FlagZeroBit) >> 9

		cw := Instructions[inst][step]

		if step == 2 {
			// Overriding default instruction when zero flag is set
			if zeroFlag == 1 {
				switch inst {
				case Inst_JZ:
					cw = Instructions[Inst_JMP][step]
				case Inst_JNZ:
					cw = Instructions[Inst_NOP][step]
				}
			}

			// Overriding default instruction when carry flag is set
			if carryFlag == 1 && inst == Inst_JC {
				switch inst {
				case Inst_JC:
					cw = Instructions[Inst_JMP][step]
				case Inst_JNC:
					cw = Instructions[Inst_NOP][step]
				}
			}
		}

		log.Printf("Registering instruction %01b %01b %01b %04b %03b => %016b", zeroFlag, carryFlag, bs, inst, step, cw)
		// Writing control word
		if bs == 0 {
			WriteByteToFile(file, GetControlWordRight(cw), int64(addr))
		} else {
			WriteByteToFile(file, GetControlWordLeft(cw), int64(addr))
		}
	}

	// Registering all instructions defined
	//for k, cw := range Instructions {
	//	log.Printf("Registering instruction %04b", k)
	//	WriteToFile(file, byte(k), cw)
	//}
}

// WriteByteToFile writes a byte to a file at a given offset
func WriteByteToFile(f *os.File, b byte, off int64) {
	n, err := f.WriteAt([]byte{b}, off)
	if n != 1 || err != nil {
		log.Printf("Error writing byte %08b (0x%02x) at address %02x\n", b, b, off)
	}
}
