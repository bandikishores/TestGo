package data

import (
	"database/sql/driver"
	fmt "fmt"
	"time"

	types "github.com/gogo/protobuf/types"
)

// Scan ...
func (ts *Timestamp) Scan(value interface{}) error {
	switch t := value.(type) {
	case time.Time:
		var err error
		ts.Timestamp, err = types.TimestampProto(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Not a protobuf Timestamp")
	}
	return nil
}

// Value ...
func (ts *Timestamp) Value() (driver.Value, error) {
	return types.TimestampFromProto(ts.Timestamp)
}
