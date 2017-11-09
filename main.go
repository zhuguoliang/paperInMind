package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"os"
	"./util"
	"strings"
	"bytes"
)
//This function just inserts '\n' each Nth element
func insertNth(s string,n int) string {
    var buffer bytes.Buffer
    var n_1 = n - 1
    var l_1 = len(s) - 1
    for i,rune := range s {
       buffer.WriteRune(rune)
       if i % n == n_1 && i != l_1  {
          buffer.WriteRune('\n')
       }
    }
    return buffer.String()
}

func ParsePaperInfo() (paperList []util.Paper){

	paperList=[]util.Paper{}
	doc, err := goquery.NewDocument("https://www.usenix.org/conference/atc17/technical-sessions")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Before find...")
	doc.Find(".field-item article .field article").Each(func(i int, s *goquery.Selection) {
		//trackname :=s.Find(".content .field-name-field-track-title").Text()
		curpaper:=util.Paper{}
		curpaper.Id = i
		curpaper.Title = s.Find("a").Text()
		curpaper.Title = insertNth(curpaper.Title,40)
		curpaper.AuthorList= s.Find(".content .field-group-html-element .field-items .field-item p").Text()	
		curpaper.AuthorList = insertNth(curpaper.AuthorList,40)
		tmpstr:= s.Find(".content .field-items .field-item .field-name-field-paper-description-long p").Text()
		//remove character
		curpaper.Abstract = strings.Replace(tmpstr,"\"", "", -1)
		curpaper.Abstract = strings.Replace(curpaper.Abstract,"\n", "", -1)
		curpaper.Abstract = insertNth(curpaper.Abstract,40)
		paperList = append(paperList,curpaper)
	})
	fmt.Println("After find...")
	return paperList
}

func main() {
	p:=ParsePaperInfo()
	//fmt.Printf("[Paper %d] is %s\n[Author list] is %s\n[Abstract] is %s\n\n",
	//	3,p[3].Title,p[1].AuthorList,p[3].Abstract)
	util.Write2Dot(p,os.Stdout)
	util.Write2Dotf(p,"./test.dot")

}
