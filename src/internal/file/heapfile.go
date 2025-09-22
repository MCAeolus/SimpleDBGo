package file

import (
	"hash/fnv"
	"os"

	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/transaction"
	"nathan.simpledb/src/internal/tuple"
)

func NewHeapFile(f *os.File, path string, td tuple.TupleDesc) HeapFile {
	// generate hash
	h := fnv.New32a()
	h.Write([]byte(path))

	return HeapFile{
		backingFile: f,
		backingDescription: td,
		id: int(h.Sum32()),
	}
}

type HeapFile struct {
	backingFile *os.File
	backingDescription tuple.TupleDesc
	id int //ref to table id
}

func (hf *HeapFile) GetFile() *os.File {
	return hf.backingFile
}

func (hf *HeapFile) GetId() int {
	return hf.id
}

func (hf *HeapFile) GetTupleDesc() tuple.TupleDesc {
	return hf.backingDescription
}

// this will return a heap page
func (hf *HeapFile) ReadPage(pid interfaces.PageId) interfaces.Page {
	//TODO: fix this up
	var page interfaces.Page
	return page
}

func (hf *HeapFile) WritePage(page interfaces.Page) error {
	// todo
	return nil
}

func (hf *HeapFile) NumPages() int {
	return -1 //TODO
}

func (hf *HeapFile) AddTuple(tid transaction.TransactionId, tup tuple.Tuple) error {
	return nil
}

func (hf *HeapFile) DeleteTuple(tid transaction.TransactionId, tup tuple.Tuple) error {
	return nil
}

func (hf *HeapFile) Iterator(tid transaction.TransactionId) interfaces.DbFileIterator {
	return nil
}