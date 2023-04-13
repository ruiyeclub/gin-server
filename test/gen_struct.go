package main

import (
	"fmt"
	"github.com/tiantianlikeu/converter"
)

func main() {
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		StructNameToHump:  true,
		RmTagIfUcFirsted:  false, // 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		TagToLower:        false, // tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		UcFirstOnly:       false, // 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		JsonTagToHump:     true,  // 驼峰处理
		JsonTagFirstLower: true,  // json  tag 驼峰首字母小写
	})
	err := t2t.
		SavePath("./src/refer_reward.go").
		//SavePath("./wx_user.go").
		Dsn("root:password@tcp(127.0.0.1:3306)/rosetta_monitor?charset=utf8mb4").
		TagKey("gorm").
		PackageName("refer_reward").
		RealNameMethod("dex_refer_reward").
		EnableJsonTag(true).
		EnableFormTag(true).
		Table("dex_refer_reward").
		Run()
	fmt.Println(err)
}
