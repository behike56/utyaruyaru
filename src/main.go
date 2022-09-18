// @Title
// @Description
// @Author
// @Update
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

/*
テストデータ
*/
func main() {

	scriptPath := "./script/data_input.sh"

	dataPath := inputTestDataPath()
	// currentPath, err := os.Getwd()

	output, err := exec.Command(scriptPath, dataPath).Output()

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print("実行結果: " + string(output))
}

/*
テストデータを入力(コンソール)
*/
func inputTestDataPath() string {

	fmt.Print("テストデータのディレクトリを教えてください。")

	scanner := bufio.NewScanner(os.Stdin)

	var dataPath string

	for {

		scanner.Scan()
		inputDir := scanner.Text()

		fmt.Println("ディレクトリ:> ", inputDir)

		judge := checkDirExistence(inputDir)

		switch judge {
		case true:
			dataPath = inputDir
			goto END

		case false:
			fmt.Println("入力されたディレクトリは存在しません。")
			continue

		default:
			fmt.Println("不正な操作がおこなわれました。")
			continue
		}
	}
END:
	return dataPath
}

/*
ディレクトリが存在するか判定する
*/
func checkDirExistence(dir string) bool {

	if len(dir) > 0 {

		return false
	}

	if _, err := os.Stat(dir); err == nil {

		return true
	} else {

		return false
	}
}
