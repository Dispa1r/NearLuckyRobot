package main

import (
	"NearHackathon/common"
	"NearHackathon/global"
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init(){
	InitConfig()
	global.DBEngine = common.InitDB()
}

func InitConfig(){
	workDir,_ :=os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+"/config")
	err :=viper.ReadInConfig()
	if err!=nil{
		log.Println("read config failed")
	}
}

func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func main() {
	//err := CallClaim("dispa1r1.testnet","2hTSdCSMoweRCDXnNzpKZzVbxxDFXgWAGveHSrbXqBoKkiDuVMvNERJ26C9H2UMj7jycUeai5VbZMw7ZhQcHejiB")
	//num,_ := strconv.Atoi(numOnly)
	//fmt.Println(result)
	//_, privateKey := CallSendLuckCmd(5.5,5)
	//fmt.Println(privateKey)
	fmt.Println(global.DBEngine)
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "127.0.0.1:7890",

		Token:       "1999455938:AAH0DvOuixbjWgOarKCi0ywNZzBRnS_uBro",
		Poller:      &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
	})

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
		b.Send(m.Chat,"fuck u" + m.Text)
		b.Send(m.Chat,"let us make a easy game")
		b.Handle("/xiazhu",func(m *tb.Message) {
			b.Send(m.Chat,"fuck u" + "nixiade s ")
		})
		fmt.Println(m.Sender)
		fmt.Println(m.Text)
		fmt.Println(m.Chat)
		fmt.Println(m)
	})

	b.Handle("/bind",func(m *tb.Message) {
		if common.CheckBinded(m.Sender.ID) {
			b.Send(m.Sender, "u have bind the account!!!")
			return
		}else {
			b.Send(m.Sender, "Please input your near account id")
			b.Handle(tb.OnText, func(m *tb.Message) {
				// all the text messages that weren't
				// captured by existing handlers
				nearAccount := m.Text
				if !strings.HasSuffix(nearAccount, ".testnet") {
					b.Send(m.Sender, "please input the right testnet account!")
				}
				err := common.BindNearAccount(nearAccount, m.Sender.ID, m.Sender.Username)
				if err != nil {
					b.Send(m.Sender, "fail to bind near account")
				} else {
					b.Send(m.Sender, "success!")
				}

			})
		}
	})
	b.Handle("/deposit",func(m *tb.Message) {
		b.Send(m.Sender, "please transfer to the near account:dispa1r.testnet")
		b.Send(m.Sender, "after transfer, please input your txn hash")
		b.Send(m.Sender, "the deposit must be bigger than 1 near")
		b.Handle(tb.OnText, func(m *tb.Message) {
			// all the text messages that weren't
			// captured by existing handlers
			url := "https://explorer.testnet.near.org/transactions/" + m.Text
			result := Get(url)
			//fmt.Println(result)
			reg1 := regexp.MustCompile(`"receiverId":"(?s:(.*?)).testnet",`)
			result1 := reg1.FindAllStringSubmatch(result, -1)
			receiverId := result1[0][1]
			reg2 := regexp.MustCompile(`"deposit":"(?s:(.*?))",`)
			result1 = reg2.FindAllStringSubmatch(result, -1)
			tmp := result1[0][1]
			numOnly := strings.TrimSuffix(tmp, "\"}}],\"status\":\"SuccessValue")
			var numOnly1 string
			numOnly1 = strings.TrimSuffix(numOnly, "0000000000000000000000")
			num,err := strconv.Atoi(numOnly1)
			num1 := float64(num)/100
			if err!=nil || num <100{
				b.Send(m.Sender, "invalid number")
				return
			}
			if receiverId != "dispa1r" {
				b.Send(m.Sender, "invalid transaction")
				return
			}
			err = common.GenerateTxn(m.Text)
			if err !=nil{
				b.Send(m.Sender, "invalid transaction")
				return
			}
			common.DepositMoney(m.Sender.ID, num1)
			num2 := strconv.FormatFloat(num1,'f',5,32)
			str3 := "success to deposit "+num2+" near"
			b.Send(m.Sender, str3)
			return
		})

	})

	b.Handle("/lucky",func(m *tb.Message) {
		//chat := m.Chat
		b.Send(m.Sender, "input the number of the near:")
		var amount float64

		b.Handle(tb.OnText, func(m *tb.Message) {
			amount,err = strconv.ParseFloat(m.Text,10)
			result := common.CheckIfHave(m.Sender.ID,amount)
			if err !=nil || !result{
				b.Send(m.Sender, "invalid number, please go to deposit")
			}
			b.Send(m.Sender, "input the number of the red-packets:")
			b.Handle(tb.OnText, func(m *tb.Message) {
				var num int
				num,err = strconv.Atoi(m.Text)
				if err !=nil{
					b.Send(m.Sender, "invalid number")
				}
				var privateKey []string
				err,privateKey = CallSendLuckCmd(amount,num)
				if err!=nil{
					common.SendLuck(m.Sender.ID,amount)
				}
				b.Send(m.Sender, "success to call the send luck,now give u the private key")
				for i := range privateKey{
					b.Send(m.Sender, privateKey[i])
				}
				return
			})
		})
	})

	b.Handle("/claim",func(m *tb.Message) {
		b.Send(m.Sender, "please input the private key to claim the drop")
		result := common.CheckBinded(m.Sender.ID)
		if !result{
			b.Send(m.Sender,"please first bind the near account")
			return
		}
		b.Handle(tb.OnText, func(m *tb.Message) {
			privateKey := m.Text
			accountId := common.GetAccountId(m.Sender.ID)
			err = CallClaim(accountId,privateKey)
			//result := common.CheckIfHave(m.Sender.ID,amount)
			if err !=nil || !result{
				b.Send(m.Sender, "fail to claim the drop")
			}
			b.Send(m.Sender, "success to claim the drop!")
			return
		})

	})

	b.Handle("/withdraw",func(m *tb.Message) {
		result := common.CheckBinded(m.Sender.ID)
		if !result{
			b.Send(m.Sender,"please first bind the near account")
			return
		}
		num := common.GetAmount(m.Sender.ID)
		num2 := strconv.FormatFloat(num,'f',5,32)
		str3 := "now you have " + num2 +" near"
		b.Send(m.Sender, str3)
		b.Send(m.Sender, "please input the near amount to withdraw")
		b.Handle(tb.OnText, func(m *tb.Message) {
			num3 := m.Text
			num4, _ := strconv.ParseFloat(num3,10)
			result := common.CheckIfHave(m.Sender.ID,num4)
			if !result{
				b.Send(m.Sender, "please input the number lower than you have")
				return
			}
			err = CallClaim(common.GetAccountId(m.Sender.ID),num3)
			//result := common.CheckIfHave(m.Sender.ID,amount)
			if err !=nil || !result{
				b.Send(m.Sender, "fail to withdraw the money")
			}
			b.Send(m.Sender, "success to withdraw the money!")
			return
		})

	})


	b.Start()
}
