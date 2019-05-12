package error

import (
	"errors"
	"fmt"
	"testing"
)




func TestPanVxExit(t *testing.T){
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("恢复错误",err)
		}
	}()
	fmt.Println("Start")
	//os.Exit(15)
	panic(errors.New("Something wrong!"))
}