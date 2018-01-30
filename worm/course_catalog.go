package worm

import (
	"github.com/biaocheng/word_worm/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/logs"
)

func course_catalog_worm(typeWord *models.WordType,classId string)  {
	url := "http://word.iciba.com/?action=courses&classid="+classId
	doc2 ,err := goquery.NewDocument(url)
	if err!=nil{
		beego.Error(err)
	}
	doc2.Find(".mid li").Each(func(i int, selection5 *goquery.Selection) {
		wordType := new(models.WordType)
		wordType.Name = strings.TrimSpace(selection5.Find("h4").Text())
		wordType.Parent = typeWord
		wordType.HasChild = false

		logs.Debug("============courseCatalog:",wordType.Name,"==============")
		wordType.Insert()

		wordsUrlId,_ := selection5.Attr("course_id")
		wordsUrlId = strings.TrimSpace(wordsUrlId)
		getWords(wordType,classId,wordsUrlId)

	})
}