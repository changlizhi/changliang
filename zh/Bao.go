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
func (zh *Zh) Beegocorsbao() string {
	//"github.com/astaxie/beego/plugins/cors"
	ret := zfzhi.Zhi.Syh() +
		zf.Zfs.Github(true) +
		zfzhi.Zhi.Dh() +
		zf.Zfs.Com(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Astaxie(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Beego(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Plugins(true) +
		zfzhi.Zhi.Xx() +
		zf.Zfs.Cors(true) +
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

func (zh *Zh) Zfszhtrue(mokuai string) string {
	return zfszh(mokuai, zf.Zfs.True(true))
}
func (zh *Zh) Zfszhfalse(mokuai string) string {
	return zfszh(mokuai, zf.Zfs.False(true))
}

//zh.Zhs.Xxx()
func Zhszh(fangfa string) string {
	ret := zf.Zfs.Zh(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhs(false) + zfzhi.Zhi.Dh() + fangfa + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	return ret
}

//zf.Zfs.Xxx(true/false)
func zfszh(mokuai string, boolstr string) string {
	ret := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) +
		zfzhi.Zhi.Dh() + mokuai + zfzhi.Zhi.Xkhz() +
		boolstr + zfzhi.Zhi.Xkhy()
	return ret
}

//zfzhi.Zhi.xxx()
func (zh *Zh) Zhiszh(zhi string) string {
	ret := zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) +
		zfzhi.Zhi.Dh() + zhi + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	return ret
}

//log.Println("xxx",xxx)
func (zh *Zh) Logszh(zhi string) string {
	ret := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Syh() + zhi +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyh() +
		zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() + zhi + zfzhi.Zhi.Xkhy()
	return ret
}

func (zh *Zh) AllowMethods() []string {
	ret := []string{}
	ret = append(ret, zf.Zfs.GET(false))
	ret = append(ret, zf.Zfs.POST(false))
	ret = append(ret, zf.Zfs.PUT(false))
	ret = append(ret, zf.Zfs.DELETE(false))
	ret = append(ret, zf.Zfs.OPTIONS(false))
	return ret
}
func (zh *Zh) AllowHeaders() []string {
	ret := []string{}
	ret = append(ret, zf.Zfs.Origin(false))
	ret = append(ret, zf.Zfs.Authorization(false))
	//Access-Control-Allow-Origin
	acao := zf.Zfs.Access(false) + zfzhi.Zhi.Jian() + zf.Zfs.Control(false) +
		zfzhi.Zhi.Jian() + zf.Zfs.Allow(false) + zfzhi.Zhi.Jian() + zf.Zfs.Origin(false)
	ret = append(ret, acao)

	//Access-Control-Allow-Headers
	acah := zf.Zfs.Access(false) + zfzhi.Zhi.Jian() + zf.Zfs.Control(false) +
		zfzhi.Zhi.Jian() + zf.Zfs.Allow(false) + zfzhi.Zhi.Jian() + zf.Zfs.Headers(false)
	ret = append(ret, acah)
	//Content-Type
	ct := zf.Zfs.Content(false) + zfzhi.Zhi.Jian() + zf.Zfs.Type(false)
	ret = append(ret, ct)
	return ret
}
func (zh *Zh) ExposeHeaders() []string {
	ret := []string{}
	//Content-Length
	cl := zf.Zfs.Content(false) + zfzhi.Zhi.Jian() + zf.Zfs.Length(false)
	ret = append(ret, cl)
	//Access-Control-Allow-Origin
	acao := zf.Zfs.Access(false) + zfzhi.Zhi.Jian() + zf.Zfs.Control(false) +
		zfzhi.Zhi.Jian() + zf.Zfs.Allow(false) + zfzhi.Zhi.Jian() + zf.Zfs.Origin(false)
	ret = append(ret, acao)

	//Access-Control-Allow-Headers
	acah := zf.Zfs.Access(false) + zfzhi.Zhi.Jian() + zf.Zfs.Control(false) +
		zfzhi.Zhi.Jian() + zf.Zfs.Allow(false) + zfzhi.Zhi.Jian() + zf.Zfs.Headers(false)
	ret = append(ret, acah)
	//Content-Type
	ct := zf.Zfs.Content(false) + zfzhi.Zhi.Jian() + zf.Zfs.Type(false)
	ret = append(ret, ct)
	return ret
}
func (zh *Zh) Httpbao() string {
	//"net/http"
	ret := zfzhi.Zhi.Syh() + zf.Zfs.Net(true) + zfzhi.Zhi.Xx() + zf.Zfs.Http(true) + zfzhi.Zhi.Syh()
	return ret

}

func (zh *Zh) Httptestbao() string {
	//"net/http/httptest"
	ret := zfzhi.Zhi.Syh() + zf.Zfs.Net(true) + zfzhi.Zhi.Xx() + zf.Zfs.Http(true) +
		zfzhi.Zhi.Xx() + zf.Zfs.Httptest(true) + zfzhi.Zhi.Syh()
	return ret

}
func (zh *Zh) Conveybao() string {
	//. "github.com/smartystreets/goconvey/convey"
	ret := zfzhi.Zhi.Dh() + zfzhi.Zhi.Syh() + zf.Zfs.Github(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Com(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Smartystreets(true) + zfzhi.Zhi.Xx() + zf.Zfs.Goconvey(true) +
		zfzhi.Zhi.Xx() + zf.Zfs.Convey(true) + zfzhi.Zhi.Syh()
	return ret
}
