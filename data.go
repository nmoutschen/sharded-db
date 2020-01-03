package sharded

import "time"

//Data represents a piece of information
type Data struct {
	ID   int
	Time time.Time
	Data string
}

//NewData creates a new Data
func NewData(id int, time time.Time, data string) *Data {
	d := &Data{
		ID:   id,
		Time: time,
		Data: data,
	}

	return d
}
