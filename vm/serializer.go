package vm

import (
	"simple-db-go/backend"
	"simple-db-go/frontend"
)

func Serialize(insertRow frontend.Row, cursor *backend.Cursor) {
	rowBytes := frontend.RowToBytes(insertRow)
	copy(cursor.GetValue(), rowBytes[:])
	backend.TABLE.RowNumber += 1
}

func Deserialize(bytes []byte) (row frontend.Row) {
	var lastRow [backend.RowSize]byte
	copy(lastRow[:], bytes[:backend.RowSize])
	row = frontend.BytesToRow(lastRow)
	return row
}
