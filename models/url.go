/*
* date: 2019-04-25 10:01
* author: liaodong
* email: liaoshp@qq.com
 */

package models

import (
	"time"

	"github.com/marksalpeter/token"
)

type Url struct {
	Id         int64  `xorm:"pk autoincr"`
	Code       string `xorm:"varchar(15) index(code) unique(urlcode)"`
	Target     string `xorm:"varchar(3000)"`
	Hijeck     string `xorm:"varchar(3000) default('')"`
	Title      string `xorm:"varchar(500) default('')"`
	PageTitle  string `xorm:"varchar(500) default('')"`
	UserId     int64
	Expire     int64
	Status     int
	Note       string    `xorm:"varchar(250)"`
	CreateTime time.Time `xorm:"created" json:"create_time"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"`
}

type RequestLog struct {
	Id           int64  `xorm:"pk autoincr"`
	UrlId        int64  `xorm:"index(url_id)"`
	UserId       int64  `xorm:"index(user_id)"`
	Uuid         string `xorm:"varchar(50) index(uuid)"`
	Target       string `xorm:"varchar(3000)"`
	Referrer     string `xorm:"varchar(250) index(referrer)"`
	ReqIp        string `xorm:"varchar(30) index(ip)"`
	ReqUseragent string `xorm:"varchar(200)"`
	ReqHeader    string `xorm:"varchar(500)"`
	Country      string `xorm:"varchar(30)"`
	Region       string `xorm:"varchar(30)"`
	City         string `xorm:"varchar(30)"`
	Isp          string `xorm:"varchar(30)"`
	Mobile       bool
	Bot          bool
	Mozilla      string `xorm:"varchar(30)"`
	Platform     string `xorm:"varchar(30)"`
	Os           string `xorm:"varchar(30)"`
	EngineName   string `xorm:"varchar(30)"`
	EngineVer    string `xorm:"varchar(30)"`
	BrowserName  string `xorm:"varchar(30)"`
	BrowserVer   string `xorm:"varchar(30)"`

	CreateTime time.Time `xorm:"created index(ctime)" json:"create_time"`
}

func GetCode() string {
	// 默认生成code长度为5的短网址
	code := token.New(5)
	return code.Encode()
}
