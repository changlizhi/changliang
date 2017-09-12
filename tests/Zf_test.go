package tests

import (
	"changliang/fanshe"
	"changliang/zf"
	"log"
	"testing"
)

func TestPrintname(t *testing.T) {
	log.Println(fanshe.Fangfaming(false))
	log.Println(zf.Zfs.Fangfaming(false))
}
