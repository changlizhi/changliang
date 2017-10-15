package zfzhi

import "changliang/zf"

func (zfzhi *Zfzhi) Mhsydata() string {
	return zf.Zfs.Mhsydata(false)
}
