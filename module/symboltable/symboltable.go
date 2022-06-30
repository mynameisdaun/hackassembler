package symboltable

type symbolTable struct {
	memoryMap   map[string]int
	memoryIndex int
}

var table *symbolTable

var memoryMap = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

func SymbolTable() *symbolTable {
	if table == nil {
		table = &symbolTable{
			memoryMap:   memoryMap,
			memoryIndex: 15,
		}
	}
	return table
}

func GetAddress(table *symbolTable, symbol string) int {
	return table.memoryMap[symbol]
}

func Contains(table *symbolTable, symbol string) bool {
	if _, ok := table.memoryMap[symbol]; ok {
		return true
	}
	return false
}

func (s *symbolTable) AddVariable(symbol string) {
	index := s.memoryIndex
	memoryMap := s.memoryMap
	memoryMap[symbol] = index + 1
	s.memoryIndex = index + 1
}

func (s *symbolTable) AddCommand(symbol string, lineNumber int) {
	memoryMap := s.memoryMap
	memoryMap[symbol] = lineNumber
}
