package queue

type Node struct {
	Value int
}

type Stack struct {
	nodes	[]*Node
	count	int
}

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) Queue {
    //a := make([]interface, size)
	return Queue
}

func Push(key interface{}){

}

func Contains(key interface{}) bool{
    return true
}
/*
func Keys() []interface{}{

    return 
}
*/
