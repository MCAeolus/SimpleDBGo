package interfaces

import (
	"io"

	"nathan.simpledb/src/internal/enums"
	"nathan.simpledb/src/internal/transaction"
	"nathan.simpledb/src/internal/tuple"
)

type Type interface {
	// bytes
	GetLen() int
	Parse(io.Reader) Field
}

type Field interface {
	Serialize(io.Writer) error
	Compare(op enums.PredicateOp, value Field) (bool, error)
	GetType() Type
}

type PageId interface {
	Serialize() []int
	GetTableId() int
	HashCode() int
	Equals(PageId) bool
	PageNo() int
}

type Page interface {
	GetId() PageId
	IsDirty() transaction.TransactionId
	MarkDirty(dirty bool, tid transaction.TransactionId)
	GetPageData() []byte
	GetBeforeImage() Page
	SetBeforeImage() Page
}

type DbFile interface {
	ReadPage(id PageId) Page
	WritePage(p Page)
	// returns: modified pages
	AddTuple(tid transaction.TransactionId, t tuple.Tuple) []Page
	DeleteTuple(tid transaction.TransactionId, t tuple.Tuple) Page
	Iterator(tid transaction.TransactionId) DbFileIterator
	GetId() int
	GetTupleDesc() tuple.TupleDesc
}

type DbFileIterator interface {
	Open() error
	HasNext() (bool, error)
	Next() (tuple.Tuple, error)
	Rewind() error
	Close()
}