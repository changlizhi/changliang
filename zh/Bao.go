package zh

import (
	"changliang/zf"
	"changliang/zfzhi"
)

func (zh *Zh) Beegobao() string {
	//"github.com/astaxie/beego"
	ret := zfzhi.Zhi.Syh() +
		zf.Zfs.Github(true) +
		zfzhi.Zhi.Dh() +
		zf.Zfs.Com(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Astaxie(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Beego(true) +
		zfzhi.Zhi.Syh() +
		zfzhi.Zhi.Hhf()
	return ret
}
func (zh *Zh) Beegoormbao() string {
	//"github.com/astaxie/beego/orm"
	ret := zfzhi.Zhi.Syh() +
		zf.Zfs.Github(true) +
		zfzhi.Zhi.Dh() +
		zf.Zfs.Com(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Astaxie(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Beego(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Orm(true) +
		zfzhi.Zhi.Syh() +
		zfzhi.Zhi.Hhf()
	return ret
}
func (zh *Zh) Mysqlbao() string {
	// _ "github.com/go-sql-driver/mysql"
	ret := zfzhi.Zhi.Xhx() +
		zfzhi.Zhi.Syh() +
		zf.Zfs.Github(true) +
		zfzhi.Zhi.Dh() +
		zf.Zfs.Com(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Go(true) +
		zfzhi.Zhi.Jian() +
		zf.Zfs.Sql(true) +
		zfzhi.Zhi.Jian() +
		zf.Zfs.Driver(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Mysql(true) +
		zfzhi.Zhi.Syh() +
		zfzhi.Zhi.Hhf()
	return ret
}

func (zh *Zh)Zfszhtrue(mokuai string) string {
	return zfszh(mokuai, zf.Zfs.True(true))
}
func (zh *Zh)Zfszhfalse(mokuai string) string {
	return zfszh(mokuai, zf.Zfs.False(true))
}
func zfszh(mokuai string, boolstr string) string {

	ret := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) +
		zfzhi.Zhi.Dh() + mokuai + zfzhi.Zhi.Xkhz() +
		boolstr + zfzhi.Zhi.Xkhy()
	return ret
}
//zfzhi.Zhi.Kzf()
func (zh *Zh)Zhiszh(zhi string) string {
	ret := zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) +
		zfzhi.Zhi.Dh() + zhi + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	return ret
}