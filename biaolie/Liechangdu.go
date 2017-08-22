package biaolie

import (
	"changliang/zfzhi"
)

func (bl *Bl) Idchangdu() int {
	return zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Biaojichangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Bianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Mingchengchangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Juesebianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Quanxianbianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Lujingchangdu() int {
	return zfzhi.Zhi.Shuzi3() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Youxiangchangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Leixingchangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Zhichangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (bl *Bl) Xinxibianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
