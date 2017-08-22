package fanshe

import (
	"runtime"
	"strings"
	"changliang/zfzhi"
)

func Fangfaming(xiaoxie bool) string {
	pc, _, _, _ := runtime.Caller(zfzhi.Zhi.Shuzi1())
	ff := runtime.FuncForPC(pc)
	ffarr := strings.Split(ff.Name(), zfzhi.Zhi.Dh())

	f := ffarr[len(ffarr) - zfzhi.Zhi.Shuzi1()]
	if xiaoxie {
		return strings.ToLower(f)
	}
	return f
}