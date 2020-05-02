package main

/**
 * Instructions set:
 * NOP     0b0000  No operation
 * LDA     0b0001  Load to the A register
 * ADD     0b0010  Load RAM addr in B register and store addition to the A register
 * SUB     0b0011  Load RAM addr in B and store subtraction to the A register
 * STA     0b0100  Store the content of the A register to the RAM addr
 * MVI     0b0101  Move immediate value to the A register
 * INC     0b0110  Increment immediate value to the A register
 * DEC     0b0111  Increment immediate value to the A register
 * JNC     0b1000  Jump if carry bit is clear
 * JNZ     0b1001  Jump if zero bit is clear
 * CMP     0b1010  Compare address value with A register - if true zero flag is set
 * JC      0b1011  Jump if carry bit is set
 * JZ      0b1100  Jump if zero bit is set
 * JMP     0b1101  Unconditional jump
 * OUT     0b1110  Out the result of A register
 * HLT     0b1111  Halt
 */

const (
	Inst_NOP = 0b0000
	Inst_LDA = 0b0001
	Inst_ADD = 0b0010
	Inst_SUB = 0b0011
	Inst_STA = 0b0100
	Inst_MVI = 0b0101
	Inst_JNC = 0b1000
	Inst_JNZ = 0b1001
	Inst_CMP = 0b1010
	Inst_JC  = 0b1011
	Inst_JZ  = 0b1100
	Inst_JMP = 0b1101
	Inst_OUT = 0b1110
	Inst_HLT = 0b1111
)

var (
	Instructions = [][]uint16{
		[]uint16{MI | CO, RO | II | CE, 0, 0, 0, 0},                             // NOP - 0000
		[]uint16{MI | CO, RO | II | CE, MI | IO, RO | AI, 0, 0},                 // LDA - 0001
		[]uint16{MI | CO, RO | II | CE, MI | IO, RO | BI, EO | AI | FI, 0},      // ADD - 0010
		[]uint16{MI | CO, RO | II | CE, MI | IO, RO | BI, SU | EO | AI | FI, 0}, // SUB - 0011
		[]uint16{MI | CO, RO | II | CE, IO | MI, AO | RI, 0, 0},                 // STA - 0100
		[]uint16{MI | CO, RO | II | CE, IO | AI, 0, 0, 0},                       // MVI - 0101
		[]uint16{MI | CO, RO | II | CE, IO | BI, EO | FI | AI, 0, 0},            // INC - 0110
		[]uint16{MI | CO, RO | II | CE, IO | BI, EO | FI | AI | SU, 0, 0},       // DEC - 0111
		[]uint16{MI | CO, RO | II | CE, IO | J, 0, 0, 0},                        // JNC - 1000
		[]uint16{MI | CO, RO | II | CE, IO | J, 0, 0, 0},                        // JNZ - 1001
		[]uint16{MI | CO, RO | II | CE, IO | MI, RO | BI, EO | SU | FI, 0},      // CMP - 1010
		[]uint16{MI | CO, RO | II | CE, 0, 0, 0, 0},                             // JC  - 1011
		[]uint16{MI | CO, RO | II | CE, 0, 0, 0, 0},                             // JZ  - 1100
		[]uint16{MI | CO, RO | II | CE, IO | J, 0, 0, 0},                        // JMP - 1101
		[]uint16{MI | CO, RO | II | CE, AO | OE, 0, 0, 0},                       // OUT - 1110
		[]uint16{MI | CO, RO | II | CE, HLT, 0, 0, 0},                           // HLT - 1111
	}
)
