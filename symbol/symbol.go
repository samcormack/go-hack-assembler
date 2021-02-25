package symbol

type SymbolTable struct {
	table map[string]int64
} 

func NewSymbolTable() *SymbolTable {
	st := new(SymbolTable)
	st.table = map[string]int64{
		"SP": 0,
		"LCL": 1,
		"ARG": 2,
		"THIS": 3,
		"THAT": 4,
		"R0": 0,
		"R1": 1,
		"R2": 2,
		"R3": 3,
		"R4": 4,
		"R5": 5,
		"R6": 6,
		"R7": 7,
		"R8": 8,
		"R9": 9,
		"R10": 10,
		"R11": 11,
		"R12": 12,
		"R13": 13,
		"R14": 14,
		"R15": 15,
		"SCREEN": 16384,
		"KBD": 24576,
	}
	return st
}

// Add new symbol to the table
func (st *SymbolTable) AddEntry(symbol string, address int64) {
	st.table[symbol] = address
}

// Return true if symbol is the table, else false
func (st *SymbolTable) Contains(symbol string) bool {
	_, present := st.table[symbol]
	return present 
}

// Get the address corresponding to the symbol
func (st *SymbolTable) GetAddress(symbol string) int64 {
	address := st.table[symbol]
	return address
}