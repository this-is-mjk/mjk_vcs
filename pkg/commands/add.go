package commands

import (
	creator "github.com/this-is-mjk/mjk/pkg/utils/componentCreater"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Add() {
	pwd := fileUtils.Get_pwd()
	creator.StageFile(pwd)
	// fileUtils.WriteFile(pwd+"/.mjk/index", *bytes.NewBuffer([]byte(creator.StageFile(pwd))))
	// files := strings.Split(string(fileUtils.ReadFile(pwd+"/.mjk/index")), "\n")
	// for _, file := range files {
	// 	if file == "" {
	// 		continue
	// 	}
	// 	fmt.Println(file)
	// }
}
