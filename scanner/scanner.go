package scanner

// Checker Scanner defines an interface to which all the checks adhere
type Checker interface {
	Check(host string, port uint64) *Result
}

// Result defines the outcome of a check
type Result struct {
	Vulnerable bool
	Details    string
}
