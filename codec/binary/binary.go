package binary

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/goobjfmt"
)

type binaryCodec struct {
}

func (self *binaryCodec) Name() string {
	return "binary"
}

func (self *binaryCodec) Encode(msgObj interface{}) (data interface{}, err error) {

	return goobjfmt.BinaryWrite(msgObj)

}

func (self *binaryCodec) Decode(data interface{}, msgObj interface{}) error {

	return goobjfmt.BinaryRead(data.([]byte), msgObj)
}

func init() {

	cellnet.RegisterCodec(new(binaryCodec))
}
