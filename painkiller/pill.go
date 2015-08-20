package painkiller

type Pill int

//go:generate stringer -type=Pill
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
