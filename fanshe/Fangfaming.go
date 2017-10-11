package fanshe

import (
	"runtime"
	"strings"
)

func Fangfaming(xiaoxie bool) string {
	pc, _, _, _ := runtime.Caller(1)
	ff := runtime.FuncForPC(pc)
	ffarr := strings.Split(ff.Name(), ".")

	f := ffarr[len(ffarr) - 1]
	if xiaoxie {
		return strings.ToLower(f)
	}
	return f
}
