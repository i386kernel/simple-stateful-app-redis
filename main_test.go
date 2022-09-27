package main

import (
	"testing"
)

func TestAddUP(t *testing.T) {
	sum := addUp(2, 2)
	if sum != 4 {
		t.Errorf("Test Failed!!!, Wanted 4 got %v", sum)
	}
}

//
//var num = 1000
//
//func BenchmarkRedisReq(t *testing.B) {
//	fmt.Println("====Querying 1000 items from Redis=====")
//	for i := 0; i < num; i++ {
//		redisGET("IST")
//	}
//}

//func TestgetCurrentTimeWithTimeDifference(t *testing.T) {
//
//	time1 := time.Now()
//	time2 := time.Now()
//
//	if time1 != time2 {
//		t.Error("Test Failed!!!!, issues with time sync")
//	}
//}
