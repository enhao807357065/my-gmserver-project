package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&Register{})
	Processor.Register(&Back{})
}

type Hello struct {
	Name string	`json:"name"`
}

type Register struct {
	Name		string		`bson:"name"`
	Age		int		`bson:"age"`
}

type Back struct {
	Success		string		`json:"success"`
}