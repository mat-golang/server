package ws

import (
	"bytes"
	"encoding/json"
)

func init() {
	pub := &Publisher{}

	pub.HandleFunc("connect", func(sub *Subcriber) {
		var foo Foo
		err := json.NewDecoder(bytes.NewReader(sub.Msg.Data)).Decode(&foo)
		if err != nil {
			panic(err)
		}
	})

}

type Foo struct {
}
