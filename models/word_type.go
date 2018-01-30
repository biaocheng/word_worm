/**
单词分类
 */
package models

import (
	"github.com/astaxie/beego/logs"
)

type WordType struct{
	Id int64
	Name string `xorm:"varchar(255) null"`
	Parent *WordType `xorm:"null"`
	HasChild bool
}

func (this *WordType) Insert(){
	_,err := DBWrite.Insert(this)
	if err!=nil{
		logs.Error(err)
	}

}