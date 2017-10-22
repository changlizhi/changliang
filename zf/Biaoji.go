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
func (zf *Zf) Yiban(xiaoxie bool) string { //一般的数据表都不做特殊处理
	return fanshe.Fangfaming(xiaoxie)
}
func (zf *Zf) Yemianlianjie(xiaoxie bool) string {
	//这个标记是给ziyuan表用的，如果是资源表则生成一个特别的控制器展示数据
	return fanshe.Fangfaming(xiaoxie)
}
