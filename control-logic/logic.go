package main

const (
	// Control word's bits
	// Halt
	HLT = 1 << 15
	// Memory address in
	MI = 1 << 14
	// Ram out
	RO = 1 << 13
	// Ram in
	RI = 1 << 12
	// Instruction register in
	II = 1 << 11
	// Instruction register out
	IO = 1 << 10
	// Counter out
	CO = 1 << 9
	// Program counter enabled
	CE = 1 << 8
	// A register out
	AO = 1 << 7
	// A register in
	AI = 1 << 6
	// Sum out
	EO = 1 << 5
	// Subtract enabled
	SU = 1 << 4
	// Flags in
	FI = 1 << 3
	// B register in
	BI = 1 << 2
	// Out enabled
	OE = 1 << 1
	// Program counter address in (jump)
	J = 1

	// eeprom address pins
	StepBits        = 0b0000000111
	InstructionBits = 0b0001111000
	SelectBit       = 0b0010000000 // "bit select" (right or left eeprom - pin a7)
	FlagCarryBit    = 0b0100000000 // Flag Carry (pin a8)
	FlagZeroBit     = 0b1000000000 // Flag Zero (pin a9)

)

func GetControlWordLeft(w uint16) byte {
	return byte((w >> 8) & 0xFF)
}

func GetControlWordRight(w uint16) byte {
	return byte(w & 0xFF)
}
