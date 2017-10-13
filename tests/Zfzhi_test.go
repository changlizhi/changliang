package tests

import (
	"changliang/zfzhi"
	"log"
	"testing"
)

func TestPrintzhi(t *testing.T) {
	log.Println(zfzhi.Zhi.Baifenhao())
}

func TestIshuzi1(t *testing.T) {
	zi := zfzhi.Zhi.Ishuzi1(false)
	zs := zfzhi.Zhi.Ishuzi1(true)
	log.Println(zi.(int))
	log.Println(zs.(string))
}