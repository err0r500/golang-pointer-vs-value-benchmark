package pointerVsValue

type myBigStruct struct {
	f0  int64
	f1  int64
	f2  int64
	f3  int64
	f4  int64
	f5  int64
	f6  int64
	f7  int64
	f8  int64
	f9  int64
	f10 int64
	f11 int64
	f12 int64
	f13 int64
	f14 int64
	f15 int64
	f16 int64
	f17 int64
	f18 int64
	f19 int64
}

func newValue() myBigStruct {
	return myBigStruct{}
}

func newPointer() *myBigStruct {
	return &myBigStruct{}
}

func dummyFuncValue(mbs myBigStruct) {}

func dummyFuncPointer(mbs *myBigStruct) {}

func (mbs myBigStruct) dummyMethodValue() {}

func (mbs *myBigStruct) dummyMethodPointer() {}

func dummyFuncNoRefValue(myBigStruct) {}

func dummyFuncNoRefPointer(*myBigStruct) {}

func (myBigStruct) dummyMethodNoRefValue() {}

func (*myBigStruct) dummyMethodNoRefPointer() {}
