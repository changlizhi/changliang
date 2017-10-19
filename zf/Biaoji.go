package zf

import "changliang/fanshe"

//一些特殊的标记名称
func (zf *Zf) Zhujian(xiaoxie bool) string {
	//主键，如果列带上这个标记则需要用snowflake来生成，有了这个值就不用排序字段了
	return fanshe.Fangfaming(xiaoxie)
}
func (zf *Zf) Jiance(xiaoxie bool) string {
	return fanshe.Fangfaming(xiaoxie)
}
func (zf *Zf) Zhengchang(xiaoxie bool) string {
	return fanshe.Fangfaming(xiaoxie)
}
func (zf *Zf) Wancheng(xiaoxie bool) string {
	return fanshe.Fangfaming(xiaoxie)
}
