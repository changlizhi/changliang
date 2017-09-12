package jsonlie

import (
	"changliang/zf"
)

//json所指定的内容，国际化指定了Yuyan类型读取文件
func (jl *Jl) Guojihuazhiding() string {
	return zf.Zfs.Yuyan(false)
}

//json所指定的内容，设置没有指定特别的文件
func (jl *Jl) Shezhizhiding() string {
	return zf.Zfs.Xitong(false)
}
