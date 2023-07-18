package types 

type cType interface {
	From()
}

type (
	cInt     int
	cInt8    int8
	cInt16   int16
	cInt32   int32
	cInt64   int64
	cUint    uint
	cUint8   uint8
	cUint16  uint16
	cUint32  uint32
	cUint64  uint64
	cFloat32 float32
	cFloat64 float64
	cString  string
	cBool    bool
)

func (cInt) From() {}

func (cInt8) From() {}

func (cInt16) From() {}

func (cInt32) From() {}

func (cInt64) From() {}

func (cUint) From() {}

func (cUint8) From() {}

func (cUint16) From() {}

func (cUint32) From() {}

func (cUint64) From() {}

func (cFloat32) From() {}

func (cFloat64) From() {}

func (cString) From() {}

func (cBool) From() {}