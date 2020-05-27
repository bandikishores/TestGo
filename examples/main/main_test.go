package main

import (
	"testing"
)

func BenchmarkSetNilOneByOne(b *testing.B) {
	var nums []*int
	for i := 0; i < b.N; i++ {
		nums = make([]*int, 128)
		for i := range nums {
			nums[i] = nil
		}
	}
	if len(nums) > 999 {
		b.Fail()
	}
}

func BenchmarkSetNilInBulk(b *testing.B) {
	nils := make([]*int, 32)
	var nums []*int
	for i := 0; i < b.N; i++ {
		nums = make([]*int, 128)
		for len(nums) > 0 {
			nums = nums[copy(nums, nils):]
		}
	}
	if len(nums) > 999 {
		b.Fail()
	}
}
