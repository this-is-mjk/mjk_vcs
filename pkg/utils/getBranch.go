package utils

import (
	"path/filepath"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func GetCurrentBranch() string {
	// get the current branch
	HEAD, _ := fileUtils.ReadFile(filepath.Join(".", ".mjk", "HEAD"))
	branchPointerPath := filepath.Join(".mjk", string(HEAD)[5:])
	return branchPointerPath
}
