package common

import (
	"NearHackathon/global"
	"NearHackathon/model"
	"errors"
	"fmt"
)

func BindNearAccount(nearAccount string,tgId int,username string) error{
	user := model.User{
		TgAccount:   tgId,
		NearAccount: nearAccount,
		Money:       0,
		UserName: username,
	}
	err := global.DBEngine.Create(&user).Error
	if err!=nil{
		return err
	}
	return nil
}

func CheckBinded(tgid int) bool{
	user := model.User{}
	check := global.DBEngine.Where("tg_account = ?",tgid).Find(&user).RecordNotFound()
	fmt.Println(check)
	if !check{
		return true
	}
	return false
}

func DepositMoney(tgid int,money float64) error{
	user := model.User{}
	result := global.DBEngine.Where("tg_account = ?",tgid).Find(&user).RecordNotFound()
	if result{
		return errors.New("fail to find the account")
	}
	user.Money += money
	global.DBEngine.Save(&user)
	return nil
}

func GenerateTxn(txnHash string) error{
	txn := model.Transaction{

		Txn:   txnHash,
	}
	err := global.DBEngine.Create(&txn).Error
	if err!=nil{
		return err
	}
	return nil
}

func CheckIfHave(tgid int,amount float64) bool{
	user := model.User{}
	global.DBEngine.Where("tg_account = ?",tgid).Find(&user)
	if user.Money >= amount{
		return true
	}
	return false
}

func SendLuck(tgid int,amount float64){
	user := model.User{}
	global.DBEngine.Where("tg_account = ?",tgid).Find(&user)
	user.Money-= amount
	global.DBEngine.Save(&user)
}

func GetAccountId(tgId int) string{
	user := model.User{}
	global.DBEngine.Where("tg_account = ?",tgId).Find(&user)
	return user.NearAccount

}

func GetAmount(tgid int)float64{
	user := model.User{}
	global.DBEngine.Where("tg_account = ?",tgid).Find(&user)
	return user.Money
}