package server

import (
	"fmt"
	"net/smtp"
	"strings"
)

//发送邮件的逻辑函数
func SendMail(user, password, host, to, subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func sendTxs0(eth *[]ETHEmail) {
	//邮件主题
	var text string
	for _, v := range *eth {
		text += `
			<table cellspacing="0" >
			<tr>
				<td>Type</td>
				<td>`+ v.TxType +`</td>
			</tr>
			<tr>
				<td>TokenAddress</td>
				<td>`+ v.TokenAddress +`</td>
			</tr>
			<tr>
				<td>From</td>
				<td>`+ v.From +`</td>
			</tr>
			<tr>
				<td>To</td>
				<td>`+ v.To +`</td>
			</tr>
			<tr>
				<td>Value</td>
				<td>`+ v.Value +`</td>
			</tr>
		</table>
		`+"\n"
	}
	subject := "以太坊账户变动通知"
	body := `
    <!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title></title>
		<style>
			td{
				border: 1px solid #bbb;
			}
		</style>
	</head>
	<body>
		<h3>
		以太坊账户变动通知
		</h3>
		`+ text +`
	</body>
</html>
    `
	fmt.Println("send email")
	//执行逻辑函数
	err := SendMail(EmailUser, EmailPassword, EmailHost, EmailTo, subject, body, "html")
	if err != nil {
		fmt.Println("发送邮件失败!", EmailTo)
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!", EmailTo)
	}
}

func sendTxs(eth *[]ETHEmail) {
	//邮件主题
	var text string
	for _, v := range *eth {
		text += `
			<tr>
				<td>`+ v.TxType +`</td>
				<td>`+ v.TokenAddress +`</td>
				<td>`+ v.From +`</td>
				<td>`+ v.To +`</td>
				<td>`+ v.Value +`</td>
			</tr>
		`
	}
	subject := "以太坊账户变动通知"
	body := `
    <!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title></title>
		<style>
			td{
				border: 1px solid #bbb;
			}
		</style>
	</head>
	<body>
		<h3>
		以太坊账户变动通知
		</h3>
		<table cellspacing="0" >
			<tr>
				<td>TxType</td>
				<td>TokenAddress</td>
				<td>From</td>
				<td>To</td>
				<td>Value</td>
			</tr>
			`+ text +`
		</table>
	</body>
</html>
    `
	fmt.Println("send email")
	//执行逻辑函数
	err := SendMail(EmailUser, EmailPassword, EmailHost, EmailTo, subject, body, "html")
	if err != nil {
		fmt.Println("发送邮件失败!", EmailTo)
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!", EmailTo)
	}
}
