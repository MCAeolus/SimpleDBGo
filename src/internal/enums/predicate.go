package enums

type PredicateOp int

const (
	EQUALS PredicateOp = iota
	GREATER_THAN
	LESS_THAN
	LESS_THAN_OR_EQ
	GREATER_THAN_OR_EQ
	LIKE
	NOT_EQUALS
)
