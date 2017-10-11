package zfzhi

func (zhi *Zfzhi) Idchangdu() int {
	return zhi.Shuzi10()
}
func (zhi *Zfzhi) Biaojichangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Bianmachangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Fubianmachangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Mingchengchangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Juesebianmachangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Quanxianbianmachangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Lujingchangdu() int {
	return zhi.Shuzi3() * zhi.Shuzi10() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Youxiangchangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Leixingchangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Zhichangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Xinxibianmachangdu() int {
	return zhi.Shuzi5() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Paixuchangdu() int {
	return zhi.Shuzi10()
}
func (zhi *Zfzhi) Chuangjianriqichangdu() int {
	return zhi.Shuzi2() * zhi.Shuzi10()
}
func (zhi *Zfzhi) Xiugairiqichangdu() int {
	return zhi.Shuzi2() * zhi.Shuzi10()
}
