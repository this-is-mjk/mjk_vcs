package commands

import "fmt"

func CatFile(ids []string) {
	for _, id := range ids {
		fmt.Println(id)
		
	}
}
