package models

import (
	"github.com/astaxie/beego/logs"
)

type Words struct{
	Id int64
	Word string `xorm:"varchar(255)"`
	Phonetic string `xorm:"varchar(255) null"`
	Interpretation string `xorm:"varchar(255) null"`
	MusicUrl string `xorm:"varchar(100) null"`
	WordType *WordType
}
func (this *Words) Insert(){
	_,err := DBWrite.Insert(this)
	if err!=nil{
		logs.Error(err)
	}

}
