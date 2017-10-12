package tests

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"testing"
)

func TestPrintbiaos(t *testing.T) {
	log.Println(zfzhi.Zhi.Bianmachangdu())
}

func TestPrintbiaolie(t *testing.T) {
	liearr := []string{
		"Id",
		"Bianma",
		"Paixu",
		"Biaoji",
		"Chuangjianriqi",
		"Caozuoriqi",
		"Youxiaoxing",
		"Caozuoren",
	}
	biaoarr := []string{
		"Juese",
		"Zhanghao",
		"Ziyuan",
		"Zhanghaojuese",
		"Jueseziyuan",
		"Xiangmu",
		"Xiangmuzu",
		"Xiangmuzuxiangmu",
		"Shijian",
		"Xiangmuzushijian",
		"Zhanghaoshijian",
		"Zidian",
		"Yonghu",
		"Jibing",
		"Wenzhang",
		"Liuyan",
		"Fuwufankui",
		"Xiangmushuju",
		"Chujinfangan",
		"Shebei",
		"Yinpin",
		"Yinpinbofang",
		"Yinpinxiazai",
	}
	for _, b := range biaoarr {
		buffer := &bytes.Buffer{}
		for _, l := range liearr {
			//	func (sjkmhsydata *Sjkmhsydata) Id(xiaoxie bool) string {
			//	}
			funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.Sjk(true) +
				zf.Zfs.Mhsydata(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
				zf.Zfs.Sjk(false) + zf.Zfs.Mhsydata(true) + zfzhi.Zhi.Xkhy()
			buffer.WriteString(funstr)
			csstr := b + l + zfzhi.Zhi.Xkhz() + zf.Zfs.Xiaoxie(true) + zfzhi.Zhi.Kgf() +
				zf.Zfs.Bool(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
			buffer.WriteString(csstr)
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			//return zf.Zfs.Dtziyuan(xiaoxie) + zf.Zfs.Leixing(xiaoxie)
			restr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
				zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + b + zfzhi.Zhi.Xkhz() + zf.Zfs.Xiaoxie(true) +
				zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Jia() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
				zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + l + zfzhi.Zhi.Xkhz() +
				zf.Zfs.Xiaoxie(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(restr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf())
		log.Println(buffer.String())
	}

}
