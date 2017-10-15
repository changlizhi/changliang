package zfzhi

import "changliang/zf"

func (zfzhi *Zfzhi) Guilei() string {
	return zf.Zfs.Guilei(false)
}

func (zfzhi *Zfzhi) Bianma() string {
	return zf.Zfs.Bianma(false)
}
func (zfzhi *Zfzhi) Mingcheng() string {
	return zf.Zfs.Bianma(false)
}
func (zfzhi *Zfzhi) Zhi() string {
	return zf.Zfs.Bianma(false)
}
