package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
    "crypto/sha1"
    "encoding/hex"
    "io/ioutil"
	"log"
)

func main() {
	sourcefileList := searchFiles("source/")
	targetFileList := searchFiles("target/")

	diffs := getDiffFiles(sourcefileList, targetFileList)
	//fmt.Printf("%+v\n", getDiffFiles(sourcefileList, targetFileList))
    
    newFiles := []string{}
    deletedFiles := []string{}
    
	for _, diff := range diffs {
        
        for _, fileOnSourceDir := range sourcefileList{
            if strings.Compare(diff, fileOnSourceDir)==0{
                newFiles = append(newFiles, diff)
                //continue
                break
            }
        }
        
        for _, fileOnTargetDir := range targetFileList{
            if strings.Compare(diff, fileOnTargetDir)==0{
                 deletedFiles = append(deletedFiles, diff)
                 break
            }
        }
	}
	
    //fmt.Println("+--list new files:")
	for _, newFile := range newFiles{
        fmt.Printf("%s NEW \n",newFile)
    }
    //fmt.Println()
    //fmt.Println("+--list deleted files:")
    for _, deletedFile := range deletedFiles{
        fmt.Printf("%s DELETED \n",deletedFile)
    }
    modifiedFiles := []string{}
    matchs := getMatchingFiles(sourcefileList, targetFileList)
    //fmt.Println(matchs)
	for _, file := range matchs{
            if compareCheckSum("source/"+file, "target/"+file)==false{
                modifiedFiles = append(modifiedFiles, file)
            }
            
        }
    for _, modifiedFile := range modifiedFiles{
        fmt.Printf("%s MODIFIED \n",modifiedFile)
    }
    
    //fmt.Printf("compare result is %v",compareCheckSum("source/folder1/file1.txt", "target/folder1/file1.txt"))
}

func searchFiles(dir string) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		firstPath := strings.Split(path, "/")
		editedPath := path[len(firstPath[0]):len(path)]
		if strings.Contains(editedPath, ".txt")==true{
            fileList = append(fileList, editedPath)
        }
        return nil
	})

	return fileList
}

func getDiffFiles(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}

			if !found {
				diff = append(diff, s1)
			}
		}

		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
func getMatchingFiles(slice1 []string, slice2 []string) []string {
	var matching []string

	for i := 0; i < 1; i++ {
		for _, s1 := range slice1 {
			for _, s2 := range slice2 {
				if s1 == s2 {
                    matching = append(matching, s1)
					break
				}
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
		
	}

	return matching
}

func compareCheckSum(file1 string,  file2 string) bool{
    content1, err := ioutil.ReadFile(file1)
	if err != nil {
		log.Fatal(err)
	}
	sum1 := sha1.Sum(content1)
	checkSumFile1 := hex.EncodeToString(sum1[:])
    
    
    content2, err := ioutil.ReadFile(file2)
	if err != nil {
		log.Fatal(err)
	}
	sum2 := sha1.Sum(content2)
	checkSumFile2 := hex.EncodeToString(sum2[:])
    
    if strings.Compare(checkSumFile1, checkSumFile2)!=0{
        //false means file was modified
        return false
    }else{
        return true
    }

}
