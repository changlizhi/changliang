package zfzhi

import "strconv"

func (zfzhi *Zfzhi) Shuzifu1() int {
	return -1
}
func (zfzhi *Zfzhi) Shuzi0() int {
	return 0
}
func (zfzhi *Zfzhi) Shuzi1() int {
	return 1
}
func (zfzhi *Zfzhi) Shuzi2() int {
	return 2
}
func (zfzhi *Zfzhi) Shuzi3() int {
	return 3
}
func (zfzhi *Zfzhi) Shuzi4() int {
	return 4
}
func (zfzhi *Zfzhi) Shuzi5() int {
	return 5
}
func (zfzhi *Zfzhi) Shuzi6() int {
	return 6
}
func (zfzhi *Zfzhi) Shuzi7() int {
	return 7
}
func (zfzhi *Zfzhi) Shuzi8() int {
	return 8
}
func (zfzhi *Zfzhi) Shuzi9() int {
	return 9
}
func (zfzhi *Zfzhi) Shuzi10() int {
	return 10
}

func (zfzhi *Zfzhi) Shuzifu1w() string {
	return "-1"
}
func (zfzhi *Zfzhi) Shuzi0w() string {
	return "0"
}
func (zfzhi *Zfzhi) Shuzi1w() string {
	return "1"
}
func (zfzhi *Zfzhi) Shuzi2w() string {
	return "2"
}
func (zfzhi *Zfzhi) Shuzi3w() string {
	return "3"
}
func (zfzhi *Zfzhi) Shuzi4w() string {
	return "4"
}
func (zfzhi *Zfzhi) Shuzi5w() string {
	return "5"
}
func (zfzhi *Zfzhi) Shuzi6w() string {
	return "6"
}
func (zfzhi *Zfzhi) Shuzi7w() string {
	return "7"
}
func (zfzhi *Zfzhi) Shuzi8w() string {
	return "8"
}
func (zfzhi *Zfzhi) Shuzi9w() string {
	return "9"
}
func (zfzhi *Zfzhi) Shuzi10w() string {
	return "10"
}
func (zfzhi *Zfzhi) Ishuzi1(zifu bool) interface{} {
	if !zifu {
		return 1
	}
	return strconv.Itoa(1)
}
