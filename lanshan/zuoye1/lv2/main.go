package main

import "fmt"

type Movie struct {
	Name string
	Star float64
	Time string
	Date string
}

func main() {
	m := Movie{
		Name: "西线无战事",
		Star: 8.8,
		Time: "148min",
		Date: "2022-10-28",
	}
	fmt.Printf("请输入你的命令\n1.获得名字\n2.查看星数\n3.查看时长\n4.查看上映时间\n5.退出程序\n")
	var option int
	for {
		_, err := fmt.Scanf("%d\n", &option)
		if err != nil {
			fmt.Println("输入的信息有误")
		}
		switch option {
		case 1:
			fmt.Printf("%s\n", m.Name)
		case 2:
			fmt.Printf("%v\n", m.Star)
		case 3:
			fmt.Printf("%v\n", m.Time)
		case 4:
			fmt.Printf("%v\n", m.Date)
		case 5:
			return
		default:
			fmt.Println("输入有误")

		}
	}
}
