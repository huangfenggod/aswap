package utils

import (
	"errors"
	"math"
)

func GetFNumber(in int64)(int64,error)  {
	n :=float64(0)
	sum :=float64(0)
	sum1 :=float64(0)  //上一层总和
	for i:=float64(0);i<1000;i++ {
		pow := math.Pow(3, i)
		sum +=pow
		if float64(in)<=sum {
			n=i
			break
		}
		sum1 +=pow
	}
	if n==0 {
		return 0,errors.New("out of range")
	}
	for i:=1;i<4;i++{
		y :=float64(in)-float64(i)*math.Pow(3,n-1)
		if y<=sum1{
			return int64(y),nil
		}
	}
	return 0,errors.New("range out 1000")
}
