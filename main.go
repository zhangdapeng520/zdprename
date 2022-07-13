package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_cmd"
	"github.com/zhangdapeng520/zdpgo_file"
)

var (
	rootCmd = &zdpgo_cmd.Command{}
	dir     string // 文件夹目录
	old     string // 旧字符串
	new     string // 新字符串
	isDir   bool   // 是否为修改文件夹名
)

func init() {
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "./", `指定文件夹`)
	rootCmd.Flags().StringVarP(&old, "old", "o", "", `旧字符串`)
	rootCmd.Flags().StringVarP(&new, "new", "n", "", `新字符串`)
	rootCmd.Flags().BoolVarP(&isDir, "isdir", "i", false, `是否为修改文件夹`)
}

func main() {
	// 创建一个根cmd对象
	rootCmd.Run = func(cmd *zdpgo_cmd.Command, args []string) {
		// 处理退出
		if old == "" {
			fmt.Println("要替换的字符串不能为空")
		} else {
			f := zdpgo_file.New()
			if isDir {
				f.ReplaceDirName(dir, old, new)
				fmt.Println("文件夹批量重命名成功")
			} else {
				f.ReplaceDirFilesName(dir, old, new)
				fmt.Println("文件批量重命名成功")
			}
		}
	}

	// 执行根rootCmd的命令
	rootCmd.Execute()
}
