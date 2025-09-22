package page

import "nathan.simpledb/src/internal/interfaces"

type RecordId interface {
	TupleNo() int
	GetPageId() interfaces.PageId
	Equals(o RecordId) bool
}

type CRecordId struct {
	tupleId int
	pageid interfaces.PageId
}

func (c *CRecordId) TupleNo() int {
	return c.tupleId
}

func (c *CRecordId) GetPageId() interfaces.PageId {
	return c.pageid
}

func (c *CRecordId) Equals(o RecordId) bool {
	return c.tupleId == o.TupleNo() && c.pageid == o.GetPageId()
}