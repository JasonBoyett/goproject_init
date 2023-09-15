package builder

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

// Run is the main entry point for the builder package.
func Run(tgt, project, rootName string) error {
  return nil   
}

// RunEmpty creates a new project with all the default
// options.
func RunEmpty() error {
  tgt, err := os.Getwd()
  if err != nil { return err }
  rootName, err := getValidRootName(tgt, "goproject")
  if err != nil { return err }
  return Run(tgt, "goproject", rootName)
}

// getValidRootName checks to see if the base name is
// available in the target directory. If it is not, it
// will append a number to the end of the base name to
// insure it is unique.
func getValidRootName(tgt, base string) (string, error){
  files, err := ioutil.ReadDir(tgt) 
  if err != nil { return "", err }
  for _, f := range files {
    if f.Name() == base {
      return formatRootName(base, files, 1), nil
    }
  }
  return base, nil
}

// formatRootName recursively appends an incrementing
// integer to the base name until it is unique in the
// target directory.
func formatRootName(
  base string, 
  files []fs.FileInfo, 
  inc int,
) string {
  for _, f := range files {
    if f.Name() == fmt.Sprintf("%s%d", base, inc) {
      return formatRootName(base, files, inc+1)
    }
  }
  return fmt.Sprintf("%s%d", base, inc)
}
