package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 5, 6, 8, 2}
	//a := []int{5, 4, 1, 8, 6, 2}
	
	fmt.Println("+-draw vertical barchart--+")
	drawBarChart(a)
    
    insertionSortAsc(a)
    insertionSortDesc(a)
}

func insertionSortAsc(param []int){
    fmt.Println("+-start sort asc--+")
    size := len(param)
    for i := 0; i < size; i++{
        key := param[i]
        beforeKey := i-1
        
        for beforeKey >= 0 && param[beforeKey] > key{
            temp := param[beforeKey+1]
            param[beforeKey+1] = param[beforeKey]
            param[beforeKey] = temp
            drawBarChart(param)
            fmt.Println(param)
            beforeKey = beforeKey-1
        }
       /*
        for beforeKey >= 0 && beforeKey > key{
            fmt.Print(param[beforeKey])
            temp := param[beforeKey]
            param[beforeKey] = param[key]
            param[key] = temp
            //beforeKey = beforeKey-1
        }
         */
        //fmt.Println(param)
        param[beforeKey+1] = key
        
        
    }
    fmt.Println("+--result sort asc--+")
    fmt.Println(param)
}

func insertionSortDesc(param []int){
    fmt.Println("+--start sort desc-+")
    size := len(param)
    for i := 0; i < size; i++{
        key := param[i]
        beforeKey := i-1
        
        for beforeKey >= 0 && param[beforeKey] < key{
            temp := param[beforeKey+1]
            param[beforeKey+1] = param[beforeKey]
            param[beforeKey] = temp
            drawBarChart(param)
            fmt.Println(param)
            beforeKey = beforeKey-1
        }
        param[beforeKey+1] = key
    }
    fmt.Println("+--result sort asc--+")
    fmt.Println(param)
}

func drawBarChart(param []int){
    max := getMax(param)
    //fmt.Printf("max value is %v \n",max)
    size := len(param)
    //fmt.Printf("size array is %v \n",size)
    
	for i := 0; i < max; i++ {
        
       reverseValue := max - (i+1)
        for x := 0; x < size; x++{
            //fmt.Print(x)
            o := param[x]
            //reverseValue := max - o
            
            if o <= reverseValue{
                fmt.Print(" ")
            } else {
                 fmt.Print("|")
                 
            }
        }
        fmt.Println(" ")
	}
	for i := 0; i<size;i++{
        //fmt.Print(" ")
        fmt.Print(param[i])
        
    }
    fmt.Println()
}

func getMax(param []int) int{
    var max int = param[0]
    
    for _, value := range param {
        if max < value {
            max = value
        }
    }
    return max
}


