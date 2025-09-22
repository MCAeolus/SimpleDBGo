package transaction

import "sync/atomic"

var counter atomic.Uint64

type TransactionId struct {
	myId uint64
}

func NewTransactionId() TransactionId {
	return TransactionId{
		myId: counter.Add(1), //todo: is this ok
	}
}

func (t TransactionId) GetId() uint64 {
	return t.myId
} 

func (t TransactionId) Equals(o TransactionId) bool {
	return t.myId == o.myId
} 

func (t TransactionId) HashCode() int {
	return int(t.myId)
}


