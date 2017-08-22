package zfzhi

func (zfzhi *Zfzhi) Postjuese() string {
	ret := `{"Id":1,"Bianma":"ROLE_ADMIN","Mingcheng":"管理员角色","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postjuesequanxian() string {
	ret := `{"Id":1,"Juesebianma":"ROLE_ADMIN","Quanxianbianma":"SRC_VIEW","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postquanxian() string {
	ret := `{"Id":1,"Bianma":"SRC_VIEW","Mingcheng":"供浏览的资源","Lujing":"/juese","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postxinxi() string {
	ret := `{"Id":1,"Bianma":"USER_CLZ","Mingcheng":"用户clz","Youxiang":"clz@aa.com","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postxinxijuese() string {
	ret := `{"Id":1,"Xinxibianma":"USER_CLZ","Juesebianma":"ROLE_ADMIN","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postyanzheng() string {
	ret := `{"Id":1,"Bianma":"USER_CLZ","Mingcheng":"验证1","Leixing":"Mima","Zhi":"123456","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postyanzhengleixing() string {
	ret := `{"Id":1,"Bianma":"Mima","Mingcheng":"密码验证","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchjuese() string {
	ret := `{"Id":1,"Bianma":"ROLE_ADMINxiugai","Mingcheng":"管理员角色xiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchjuesequanxian() string {
	ret := `{"Id":1,"Juesebianma":"ROLE_ADMINxiugai","Quanxianbianma":"SRC_VIEWxiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchquanxian() string {
	ret := `{"Id":1,"Bianma":"SRC_VIEWxiugai","Mingcheng":"供浏览的资源xiugai","Lujing":"/juesexiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchxinxi() string {
	ret := `{"Id":1,"Bianma":"USER_CLZxiugai","Mingcheng":"用户clzxiugai","Youxiang":"clz@aa.com","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchxinxijuese() string {
	ret := `{"Id":1,"Xinxibianma":"USER_CLZxiugai","Juesebianma":"ROLE_ADMINxiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchyanzheng() string {
	ret := `{"Id":1,"Bianma":"USER_CLZxiugai","Mingcheng":"验证1xiugai","Leixing":"Mimaxiugai","Zhi":"123456","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchyanzhengleixing() string {
	ret := `{"Id":1,"Bianma":"Mimaxiugai","Mingcheng":"密码验证xiugaxiugaii","Biaoji":"Youxiao"}`
	return ret
}
