package ben

import (
	"sync"
	"testing"
)

func BenchmarkMyMap_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mm1 := MyMap{Map: map[interface{}]interface{}{}}
		mm1.Store(i, "aaaaa")
		mm1.Load(i)
		mm1.Delete(i)
		//mm1.LoadAndDelete(i)
		mm1.LoadOrStore(i, "aaaaa")
	}
}

func BenchmarkMyMap_All2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mm1 := sync.Map{}
		mm1.Store(i, "bbbbb")
		mm1.Load(i)
		mm1.Delete(i)
		mm1.LoadOrStore(i, "bbbbb")
	}
}

func BenchmarkMyMap_Store(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var mm1 = MyMap{Map: map[interface{}]interface{}{}}

		mm1.Store(2, []int{4,5,6})

		mm1.Delete(2)

	}
}

func BenchmarkMyMap_Store2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var mm1 = sync.Map{}

		mm1.Store(2, []int{7,8,9})

		mm1.Delete(2)

	}
}

func BenchmarkMyMap_Load(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var mm1 = MyMap{Map: map[interface{}]interface{}{}}

		mm1.Load(i)

	}
}


func BenchmarkMyMap_Load2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var mm1 = sync.Map{}
		mm1.Load(i)
	}
}


