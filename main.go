package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sys/windows/registry"
)

func getRegQUERY_VALUE(str string) string {
	tKey, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`+str, registry.QUERY_VALUE)
	stringValue, _, _ := tKey.GetStringValue("")
	defer tKey.Close()
	return stringValue
}

func delALL(str []string) {
	tKey, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`, registry.QWORD)
	defer tKey.Close()
	for _, v := range str {
		delKey(v)
	}
}

func delKey(str string) bool {
	tKey, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`, registry.QWORD)
	defer tKey.Close()
	return registry.DeleteKey(tKey, str) == nil

}

func showIcons() []string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace`, registry.QWORD)
	if err != nil {
		return nil
	}
	defer k.Close()
	keys, _ := k.ReadSubKeyNames(0)
	return keys
}

func main() {
	s := showIcons()
	if len(s) < 1 {
		fmt.Println("没找到相关磁盘图标！3秒后自动退出")
		time.Sleep(time.Second * 3)
		return
	}

	var input string

	for {
		s := showIcons()
		if len(s) < 1 {
			fmt.Println("已经没有可删除的图标了")
			break
		}
		fmt.Println("检测到存在以下图标,请输入相关序号进行删除(输入a可全部删除)：")
		for i, key_subkey := range s {
			fmt.Println(i, getRegQUERY_VALUE(key_subkey))
		}
		fmt.Println("请输入要删除的图标序号")
		n, err := fmt.Scan(&input)
		if n > 1 && err != nil {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("输入有误,请重试。")
			continue
		}
		if input == "a" || input == "A" {
			delALL(s)
			break
		}
		id, err2 := strconv.Atoi(input)
		if err2 != nil {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("输入有误,请重试。")
			continue
		}

		if id > len(s)-1 {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("输入有误,请重试。")
			continue
		}
		if delKey(s[id]) {
			fmt.Println("删除成功！")
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}
	fmt.Println("清理完毕！3秒后自动退出")
	time.Sleep(time.Second * 3)
}
