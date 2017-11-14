package pointerVsValue

import "testing"

func BenchmarkUnexportedNewValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newValue()
	}
}

func BenchmarkUnexportedNewPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newPointer()
	}
}

func BenchmarkUnexportedFuncValue(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncValue(big)
	}
}

func BenchmarkUnexportedFuncPointer(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncPointer(big)
	}
}

func BenchmarkUnexportedMethodValue(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodValue()
	}
}

func BenchmarkUnexportedMethodPointer(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodPointer()
	}
}

func BenchmarkUnexportedFuncValueNoRef(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncNoRefValue(big)
	}
}

func BenchmarkUnexportedFuncPointerNoRef(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncNoRefPointer(big)
	}
}

func BenchmarkUnexportedMethodValueNoRef(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodNoRefValue()
	}
}

func BenchmarkUnexportedMethodPointerNoRef(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodNoRefPointer()
	}
}
