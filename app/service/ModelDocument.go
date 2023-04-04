package service

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
	"yssim-go/library/omc"

	"github.com/PuerkitoBio/goquery"
)

func GetModelDocument(modelName string) []string {
	docData := omc.OMC.GetDocumentationAnnotation(modelName)
	if docData[0] != "" {
		htmlStr := docData[0]
		htmlIo := strings.NewReader(htmlStr)
		doc, err := goquery.NewDocumentFromReader(htmlIo)
		if err != nil {
			panic(err)
		}
		doc.Find("img").Each(func(i int, selection *goquery.Selection) {
			uri, ok := selection.Attr("src")
			if !ok {
				return
			}
			imageFile := omc.OMC.UriToFilename(uri)
			file, err := ioutil.ReadFile(imageFile)
			if err != nil {
				return
			}
			fileBase64Str := base64.StdEncoding.EncodeToString(file)
			selection.SetAttr("src", "data:image/jpeg;base64,"+fileBase64Str)
			selection.SetAttr("href", "模型缩略图")
		})
		doc.Find("a").Each(func(i int, selection *goquery.Selection) {
			selection.RemoveAttr("href")
		})
		docData[0], _ = doc.Html()
		return docData
	}
	return docData
}

func SetModelDocument(modelName, docData, revisions string) bool {
	result := omc.OMC.SetDocumentationAnnotation(modelName, docData, revisions)
	omc.OMC.Save(modelName)
	return result
}
