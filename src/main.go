package main

import (
	"encoding/json"
	"fmt"

	"github.com/holla-world/golibr/awsnqs"
)

const (
	MsgCount      = 1
	SNSARNLOW     = "arn:aws:sns:ap-northeast-1:529673077012:mailbox-test"
	SNSARNMID     = "arn:aws:sns:ap-northeast-1:529673077012:mailbox-middle-prod"
	SNSARNPRIVATE = "arn:aws:sns:ap-northeast-1:529673077012:mailbox-message-prod"
)

type TestCMD struct {
	ToUid   []string `json:"to_uid"`
	Content string   `json:"content"`
	Channel string   `json:"channel"`
	Src     string   `json:"src"`
	Appid   int      `json:"appid"`
}

type TestContent struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

func main() {
	body := "this is a test1"
	writeSNS("holla_1", body, SNSARNPRIVATE)
}

func writeSNS(uid, content, arn string) {
	for i := 0; i < MsgCount; i++ {
		msg := TestCMD{
			ToUid:   []string{uid},
			Content: content,
			Channel: "tx",
			Src:     "holla",
		}
		b, _ := json.Marshal(msg)

		res, err := awsnqs.SNSPublish("test", arn, string(b))
		fmt.Println(res, err)
	}
}
