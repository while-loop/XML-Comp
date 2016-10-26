package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CompareContainingFolders looks to two different directories given called pathA & pathB, and creates a file named
// "missingFolders.txt"
func CompareContainingFolders(pathA, pathB string) error {
	if (pathA == "") || (pathB == "") {
		return fmt.Errorf("Empty path")
	}
	missing, err := missingFolders(pathA, pathB)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/missingFolders.txt", pathB))
	if err != nil {
		return err
	}
	defer file.Close()
	mfldr := []byte("<--Missing Folders-->\n")
	file.Write(mfldr)
	for _, v := range missing {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	return nil
}

func missingFolders(A, B string) ([]string, error) {
	var DirA []string
	var DirB []string
	var missing []string

	files, err := ioutil.ReadDir(A)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		DirA = append(DirA, v.Name())
	}
	files, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		DirB = append(DirB, v.Name())
	}
	for _, vA := range DirA {
		for iB, vB := range DirB {
			if (vA != vB) && (iB == len(DirB)-1) {
				missing = append(missing, vA)
			}
			continue
		}
	}
	if len(missing) > 0 {
		return missing, nil
	}
	return nil, fmt.Errorf("No folders missing!")
}