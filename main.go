package main

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func result(list []func() error) error{

	var errs []error
	for _, name := range list{
		err := name()
		if err != nil{
			errs = append(errs, err)
		}
	}

	var resultErr error
	for _, err := range(errs){
		resultErr = multierror.Append(resultErr, err)
	}
	return resultErr

}

func main(){

	list := []func() error{result1, result2, result3}

	fmt.Println(result(list))
}

func result1() error{
	return fmt.Errorf("result1 error")
}

func result2() error{
	return fmt.Errorf("result2 error")
}

func result3() error{
	return fmt.Errorf("result3 error")
}