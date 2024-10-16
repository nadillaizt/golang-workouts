// // membaca dan parsing file xml
// package main

// import (
// 	"encoding/json"
// 	"github.com/beevik/etree"
// 	"log"
// )

// type M map[string]interface{}

// func main() {
// 	doc := etree.NewDocument()
// 	if err := doc.ReadFromFile("./data.xml"); err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	root := doc.SelectElement("website")
// 	rows := make([]M, 0)

// 	for _, article := range root.FindElements("//article") {
// 		row := make(M)
// 		row["title"] = article.SelectElement("title").Text()
// 		row["url"] = article.SelectElement("url").Text()

// 		categories := make([]string, 0)
// 		for _, category := range article.SelectElements("category") {
// 			categories = append(categories, category.Text())
// 		}
// 		row["categories"] = categories

// 		if info := article.SelectAttr("info"); info != nil {
// 			row["info"] = info.Value
// 		}

// 		rows = append(rows, row)
// 	}

// 	bts, err := json.MarshalIndent(rows, "", "  ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(string(bts))

// }

// membuat xml dari golang
package main

import (
	"encoding/json"
	"github.com/beevik/etree"
	"log"
)

type Document struct {
	Title   string
	URL     string
	Content struct {
		Articles []struct {
			Title      string
			URL        string
			Categories []string
			Info       string
		}
	}
}

func main() {
	const jsonString = `{
		"Title": "Noval Agung",
		"URL": "https://novalagung.com",
		"Content": {
			"Articles": [{
				"Categories": [ "Server" ],
				"Title": "Connect to Oracle Server using Golang and Go-OCI8 on Ubuntu",
				"URL": "/go-oci8-oracle-linux/"
			}, {
				"Categories": [ "Server", "VPN" ],
				"Title": "Easy Setup OpenVPN Using Docker DockVPN",
				"URL": "/easy-setup-openvpn-docker/"
			}, {
				"Categories": [ "Server" ],
				"Info": "popular article",
				"Title": "Setup Ghost v0.11-LTS, Ubuntu, Nginx, Custom Domain, and SSL",
				"URL": "/ghost-v011-lts-ubuntu-nginx-custom-domain-ssl/"
			}]
		}
	}`

	data := Document{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err.Error())
	}

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	website := doc.CreateElement("website")

	website.CreateElement("title").SetText(data.Title)
	website.CreateElement("url").SetText(data.URL)

	content := website.CreateElement("contents")

	for _, each := range data.Content.Articles {
		article := content.CreateElement("article")
		article.CreateElement("title").SetText(each.Title)
		article.CreateElement("url").SetText(each.URL)

		for _, category := range each.Categories {
			article.CreateElement("category").SetText(category)
		}

		if each.Info != "" {
			article.CreateAttr("info", each.Info)
		}
	}

	doc.Indent(2)

	err = doc.WriteToFile("output.xml")
	if err != nil {
		log.Println(err.Error())
	}
}
