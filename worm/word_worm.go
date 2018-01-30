package worm

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"path"
	"github.com/astaxie/beego"
	"strings"
	"github.com/biaocheng/word_worm/models"
	"github.com/astaxie/beego/logs"
	"regexp"
)

func getWords(wordType *models.WordType,classId string,wordsUrlId string){

	wordsUrl := fmt.Sprintf("http://word.iciba.com/?action=words&class=%s&course=%s",classId,wordsUrlId)
	doc3,err := goquery.NewDocument(wordsUrl)
	if err!=nil{
		beego.Error(err)
	}
	doc3.Find(".word_main_list").Each(func(i int, selection6 *goquery.Selection) {
		selection6.Find("li").Each(func(i int, selection7 *goquery.Selection) {
			word := new(models.Words)
			defer func(){
				word.Insert()

				if err:=recover();err!=nil{
					logs.Error(err)
				}

			}()
			//获得单词
			wordName := strings.TrimSpace(selection7.Find(".word_main_list_w span").Text())

			if wordName==""{
				return
			}

			//获得音标
			yinbiao := strings.TrimSpace(selection7.Find(".word_main_list_y strong").Text())
			//获得释义
			shiyi,exists := selection7.Find(".word_main_list_s span").Attr("title")
			if exists{
				shiyi = strings.TrimSpace(shiyi)
			}else {
				shiyi = strings.TrimSpace(selection7.Find(".word_main_list_s span").Text())
			}
			wordMusicUrl,exists:= selection7.Find(".word_main_list_y a").Attr("id")

			word.Word = wordName
			word.Phonetic = yinbiao
			word.Interpretation = shiyi
			word.WordType = wordType

			reg,err := regexp.Compile(`https?://.*\.mp3`)
			//
			if err!=nil{
				logs.Error(err)
			}
			//判断是否存在读音
			if exists && reg.MatchString(wordMusicUrl){
				/*wordMusicUrl = strings.TrimSpace(wordMusicUrl)


				//下载读音
				p := path.Join("/Users/chengbiao/music",getPath(wordType),wordName+".mp3")
				os.MkdirAll(path.Dir(p),0777)
				file,err := os.Create(p)
				if err!=nil{
					beego.Error(err)
				}

				resp,_ := http.Get(wordMusicUrl)

				data := make([]byte,1024)
				for{
					len,err := resp.Body.Read(data)
					if len<=0{
						break
					}
					if err!=nil && err!=io.EOF{
						beego.Error(err)
					}
					file.Write(data[0:len])
				}
				resp.Body.Close()
				file.Close()
				*/
				word.MusicUrl = path.Join(getPath(wordType),wordName+".mp3")
			}

			logs.Debug("============word==============")

		})
	})
}

func getPath(wordType *models.WordType) string{
	if wordType.Parent!=nil{
		return path.Join(getPath(wordType.Parent),wordType.Name)
	}else{
		return wordType.Name
	}
}