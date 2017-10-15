package zfzhi

import "changliang/zf"

func (zfzhi *Zfzhi) Guojihua() string {
	return zf.Zfs.Guojihua(false)
}
