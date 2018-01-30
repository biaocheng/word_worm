package worm

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/astaxie/beego/logs"
	"github.com/biaocheng/word_worm/models"
)

func Start(){
	doc,err:= goquery.NewDocument("http://word.iciba.com/?action=index&reselect=y")
	if err!=nil{
		logs.Error(err)
	}
	doc.Find(".c_nav").Each(func(i int, selection *goquery.Selection) {
		rootWordType := new(models.WordType)
		rootWordType.Name = strings.TrimSpace(selection.Find("p").Text())
		rootWordType.HasChild = true

		logs.Debug("============one:",rootWordType.Name,"==============")
		rootWordType.Insert()

		childClassId,_:= selection.Attr("cid")
		doc.Find(".main_l").Each(func(i int, selection2 *goquery.Selection) {

			cFilter,_ := selection2.Attr("c_filter")


			if cFilter==childClassId{

				selection2.Find(".cl>li").Each(func(i int, selection3 *goquery.Selection) {
					twoWordType := new(models.WordType)
					twoWordType.Parent = rootWordType
					twoWordType.HasChild = true
					twoWordType.Name = strings.TrimSpace(selection3.Find("h3").Text())
					logs.Debug("============two:",twoWordType.Name,"==============")
					twoWordType.Insert()

					if selection3.Find(".main_l_box").Size()==0{
						classId,_ := selection3.Attr("class_id")
						course_catalog_worm(twoWordType,classId)
						return
					}

					selection3.Find(".main_l_box .nobt li").Each(func(i int, selection4 *goquery.Selection) {
						threeWordType := new(models.WordType)
						threeWordType.Name = strings.TrimSpace(selection4.Find("a h4").Text())
						threeWordType.Parent = twoWordType
						threeWordType.HasChild = true

						logs.Debug("============three:",threeWordType.Name,"==============")
						threeWordType.Insert()

						classId,_ := selection4.Attr("class_id")
						classId = strings.TrimSpace(classId)

						course_catalog_worm(threeWordType,classId)

					})

				})
			}
		})
	})
}