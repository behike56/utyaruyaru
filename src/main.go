// @Title
// @Description
// @Author
// @Update
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

/*
テストデータ
*/
func main() {

	msgWelcome()

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

func makeDirFilesList(dir string) []string {

	nodes, err := ioutil.ReadDir(dir)
	if err != nil {

		panic(err)
	}

	var paths []string
	for _, node := range nodes {

		if node.IsDir() {

			paths = append(paths, makeDirFilesList(filepath.Join(dir, node.Name()))...)
			continue
		}

		paths = append(paths, filepath.Join(dir, node.Name()))
	}

	return paths
}

func listFiles(dir string) ([]string, error) {

	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func msgWelcome() {

	welcome := ` 
 _  _   __    ____  __  __ 
( \/ ) /__\  (  _ \(  )(  )
 \  / /(__)\  )   / )(__)( 
 (__)(__)(__)(_)\_)(______)
 `

	fmt.Println(welcome)
}
