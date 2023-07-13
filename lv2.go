package main

import "fmt"

type Movie struct {
	Name         string //剧名
	director     string //导演
	scriptwriter string //编剧
	actor        string //演员
	mold         string //类型
	nation       string //国家
	language     string //语言
	date         string //上映日期
	duration     string //时长
	alias        string //别名
	IMDb         string
}

func main() {
	m := Movie{Name: "西线无战事",
		director:     "爱德华·贝尔格",
		scriptwriter: "爱德华·贝尔格 / 莱斯莉·帕特森 / 伊恩·斯托克尔 / 埃里希·玛利亚·雷马克",
		actor:        "费利克斯·卡默雷尔 / 阿尔布雷希特·舒赫 / 亚伦·希尔默 / 丹尼尔·布鲁赫 / 塞巴斯蒂安·胡克 / 更多...",
		mold:         "剧情 / 动作 / 战争",
		nation:       "德国 / 美国 / 英国",
		language:     "德语 / 法语 / 英语",
		date:         "2022-09-12(多伦多电影节) / 2022-09-29(德国) / 2022-10-28(德国网络)",
		duration:     "148分钟",
		alias:        "新西线无战事 / All Quiet on the Western Front",
		IMDb:         "tt1016150",
	}
	fmt.Printf("请输入你的命令\n1.获得名字\n2.退出程序\n")
	var option int
	for {
		fmt.Scanf("%d", &option)
		if option == 1 {
			fmt.Println(m.Name)
			fmt.Printf("请输入你的命令\n1.查看详细内容\n2.退出程序\n")
			var op int
			fmt.Scan(&op)
			if op == 1 {
				fmt.Println("导演：", m.director)
				fmt.Println("编剧：", m.scriptwriter)
				fmt.Println("主演：", m.actor)
				fmt.Println("类型：", m.mold)
				fmt.Println("制片地：", m.nation)
				fmt.Println("语言：", m.language)
				fmt.Println("上映日期：", m.date)
				fmt.Println("片长：", m.duration)
				fmt.Println("别名：", m.alias)
				fmt.Println("IMDb：", m.IMDb)
			} else if op == 2 {
				return
			}
		} else if option == 2 {
			return
		}
	}
}
