package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func CallSendLuckCmd(amount float64,num int)(error,[]string){
	str1 := strconv.FormatFloat(amount,'f',5,32)
	str2 := strconv.Itoa(num)
	command := exec.Command("node", "test.js","callSendLuck", str1, str2)
	//给标准输入以及标准错误初始化一个buffer，每条命令的输出位置可能是不一样的，
	//比如有的命令会将输出放到stdout，有的放到stderr
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	//执行命令，直到命令结束
	err := command.Run()
	if err != nil{
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(command.Stderr.(*bytes.Buffer).String())
		return err,nil
	}
	//打印命令行的标准输出
	var privateKey []string
	str3 := command.Stdout.(*bytes.Buffer).String()
	reg1 := regexp.MustCompile(`'(?s:(.*?))'`)
	result1 := reg1.FindAllStringSubmatch(str3, -1)
	for _,v := range result1{
		privateKey = append(privateKey,v[1])
	}
	return nil,privateKey
}


func CallClaim(newAccountId,privateKey string) error{
	command := exec.Command("node", "test.js","claimLuck", newAccountId, privateKey)
	//给标准输入以及标准错误初始化一个buffer，每条命令的输出位置可能是不一样的，
	//比如有的命令会将输出放到stdout，有的放到stderr
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	//执行命令，直到命令结束
	err := command.Run()
	if err != nil{
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(command.Stderr.(*bytes.Buffer).String())
		return err
	}
	//打印命令行的标准输出
	str3 := command.Stdout.(*bytes.Buffer).String()
	fmt.Println(str3)
	if strings.Contains(str3,"Error"){
		return errors.New("fail to claim the drop")
	}
	return nil
}

func CallWithdraw(newAccountId,num string) error{
	command := exec.Command("node", "test.js","withdraw", newAccountId, num)
	//给标准输入以及标准错误初始化一个buffer，每条命令的输出位置可能是不一样的，
	//比如有的命令会将输出放到stdout，有的放到stderr
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	//执行命令，直到命令结束
	err := command.Run()
	if err != nil{
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(command.Stderr.(*bytes.Buffer).String())
		return err
	}
	//打印命令行的标准输出
	str3 := command.Stdout.(*bytes.Buffer).String()
	fmt.Println(str3)
	if strings.Contains(str3,"Error"){
		return errors.New("fail to claim the drop")
	}
	return nil
}