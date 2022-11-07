package main

import (
	"fmt"
	"time"
)

var heroname string
var attack int

type Hero interface {
	Attack()
	Combo()
	say()
} //没用到接口
type Yasuo struct {
	Name         string
	HP           int
	MP           int
	Aggressivity int
	Skill        int
}

type Qiyana struct {
	Name         string
	HP           int
	MP           int
	Aggressivity int
	Skill        int
}
type Efls struct {
	Name         string
	HP           int
	MP           int
	Aggressivity int
	Skill        int
}
type Robort struct {
	Name string
	HP   int
	MP   int
}

func (Y Yasuo) say() {
	fmt.Println("死亡如风,常伴吾身。")
}

func (Q Qiyana) say() {
	fmt.Println("我超！模！")
}
func (Y Efls) say() {
	fmt.Println("千万利器，莫过于你的信念。")
}

//func Attack(H Hero) {
//
//	R := Robort{
//		Name: "Robort",
//		HP:   10000,
//		MP:   500,
//	}
//	fmt.Println("选择攻击方式\n1.平A\n2.技能连招\n3.停止攻击")
//	switch attack {
//	case 1:
//		R.HP -= H.Aggressivity
//		if R.HP < 0 {
//			R.HP = 0
//		}
//		fmt.Printf("假人剩余的血量=%d\n", R.HP)
//	case 2:
//		R.HP -= H.Skill
//		if R.HP < 0 {
//			R.HP = 0
//		}
//		fmt.Printf("假人剩余的血量=%d\n", R.HP)
//	case 3:
//		return
//
//	}
//}

func main() {
	Y := Yasuo{
		Name:         "Yasuo",
		HP:           1000,
		MP:           0,
		Aggressivity: 70,
		Skill:        200,
	}
	Q := Qiyana{
		Name:         "Qiyana",
		HP:           800,
		MP:           500,
		Aggressivity: 70,
		Skill:        10000,
	}
	E := Efls{
		Name:         "Efls",
		HP:           700,
		MP:           500,
		Aggressivity: 60,
		Skill:        250,
	}
	R := Robort{
		Name: "Robort",
		HP:   10000,
		MP:   500,
	}
	for {

		fmt.Printf("现在开始进行英雄训练营，请请输入你的命令\n1.选择你的英雄\n2.获取英雄列表\n3.退出程序\n")
		option := 0
		sum := 0
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			_, err := fmt.Scanf("%s\n", &heroname)
			if err != nil {
				fmt.Println("找不到该英雄")
			}
			if heroname == "Yasuo" {
				sum = 1
				var heroname Yasuo
				heroname.say()
				time.Sleep(time.Second)
				for {
					fmt.Println("选择攻击方式\n1.平A\n2.技能连招\n3.停止攻击")
					fmt.Scanf("%d\n", &attack)
					switch attack {
					case 1:
						R.HP -= Y.Aggressivity
						if R.HP < 0 {
							R.HP = 0
						}
						fmt.Printf("假人剩余的血量=%d\n", R.HP)
					case 2:
						R.HP -= Y.Skill
						if R.HP < 0 {
							R.HP = 0
						}
						fmt.Printf("假人剩余的血量=%d\n", R.HP)
					case 3:
						sum = 4

					}
					if sum == 4 {
						break
					}
				}

			}
			if heroname == "Qiyana" {
				sum = 2
				var heroname Qiyana
				heroname.say()
				time.Sleep(time.Second)
				for {
					fmt.Println("选择攻击方式\n1.平A\n2.技能连招\n3.停止攻击")
					fmt.Scanf("%d\n", &attack)
					switch attack {
					case 1:
						R.HP -= Q.Aggressivity
						if R.HP < 0 {
							R.HP = 0
						}
						fmt.Printf("假人剩余的血量=%d\n", R.HP)
					case 2:
						R.HP -= Q.Skill
						if R.HP < 0 {
							R.HP = 0
						}
						fmt.Printf("假人剩余的血量=%d\n", R.HP)
					case 3:
						sum = 4

					}
					if sum == 4 {
						break
					}
				}
			}
			if heroname == "Efls" {
				sum = 3
				var heroname Efls
				heroname.say()
				time.Sleep(time.Second)
				for {
					fmt.Println("选择攻击方式\n1.平A\n2.技能连招\n3.停止攻击")
					fmt.Scanf("%d\n", &attack)
					switch attack {
					case 1:
						R.HP -= E.Aggressivity
						if R.HP < 0 {
							R.HP = 0
						}
						fmt.Printf("假人剩余的血量=%d\n", R.HP)
					case 2:
						R.HP -= E.Skill
						if R.HP < 0 {
							R.HP = 0
						}
						fmt.Printf("假人剩余的血量=%d\n", R.HP)
					case 3:
						sum = 4
					}
					if sum == 4 {
						break
					}
				}
			}
			if sum == 0 {
				fmt.Println("找不到该英雄")
			}

		case 2:
			fmt.Println("Yasuo", "Qiyana", "Efls")
		case 3:
			return
		default:
			fmt.Println("输入信息有误")
		}
	}
}
