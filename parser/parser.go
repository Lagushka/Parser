package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"golang.org/x/net/html"
	"strconv"
	"io"
	"strings"
	// "bytes"
)

type tData struct {
	Img string
	Name string
	Price string
	Ref string
	PhoneNumber string
	Place string
}

type tElemData []tData

func createFolder(path string) {
	err := os.MkdirAll(path, 0700) 
	if err != nil {
		log.Panic(err)
	}
}

func phoneRequest(phoneToken string, id string) string {
	req, err:= http.NewRequest("GET", "https://www.olx.kz/ajax/misc/contact/phone/" + id + "/?pt=" + phoneToken, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Cookie", "mobile_default=desktop; dfp_segment_test=62; dfp_segment_test_v3=99; dfp_segment_test_v4=14; dfp_segment_test_oa=33; fingerprint=MTI1NzY4MzI5MTs0OzA7MDswOzA7MDswOzA7MDswOzE7MTsxOzE7MTsxOzE7MTsxOzE7MTsxOzE7MTswOzE7MTsxOzA7MDsxOzE7MTsxOzE7MTsxOzE7MTsxOzA7MTsxOzA7MTsxOzE7MDswOzA7MDswOzA7MTswOzE7MTswOzA7MDsxOzA7MDsxOzE7MDsxOzE7MTsxOzA7MTswOzg1MDQ2MjI0MjsyOzI7MjsyOzI7MjszOzEyMzc2Nzc1Nzk7MTY1OTU4OTY0OTsxOzE7MTsxOzE7MTsxOzE7MTsxOzE7MTsxOzE7MTsxOzE7MDswOzA7MzU2MzYzNjY3OzUzODA5ODc3ODsxNTI5NTY3MzgyOzMzMDgzODg0MTsxMDA1MzAxMjAzOzE1MzY7ODY0OzI0OzI0OzI0MDsxODA7MjQwOzE4MDsyNDA7MTgwOzI0MDsxODA7MjQwOzI0MDsyNDA7MjQwOzI0MDsyNDA7MjQwOzE4MDsxODA7MTgwOzE4MDsxODA7MDswOzA=; dfp_user_id=25c42d8b-2a26-40d8-92b4-88ddb1e1f26f-ver2; ldTd=true; _ym_uid=1610551330853320161; _ym_d=1610551330; __utmc=16996198; __utmz=16996198.1610551331.1.1.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); _ga=GA1.2.929845084.1610551331; laquesissu=; tmr_lvid=401068178fe74faeadb74d4d352309ba; tmr_lvidTS=1610551332912; user_adblock_status=false; __gads=ID=ca2b24b561f9066d:T=1610551334:S=ALNI_Mbehuf7d_LKDRqcAVEqvbSDVZI-Og; lister_lifecycle=1610551347; _gcl_au=1.1.1599253637.1610551349; cookieBarSeen=true; laquesisff=olxeu-29763#srt-627#srt-633#srt-635#srt-899; searchFavTooltip=1; observed_aui=7ea508a68b5b4df090c4106a2e87c831; last_locations=40-0-0-%D0%A2%D0%B0%D0%BB%D0%B4%D1%8B%D0%BA%D0%BE%D1%80%D0%B3%D0%B0%D0%BD-%D0%90%D0%BB%D0%BC%D0%B0%D1%82%D0%B8%D0%BD%D1%81%D0%BA%D0%B0%D1%8F+%D0%BE%D0%B1%D0%BB%D0%B0%D1%81%D1%82%D1%8C-taldykorgan_39-0-0-%D0%9A%D0%B0%D1%80%D0%B0%D0%B3%D0%B0%D0%BD%D0%B4%D0%B0-%D0%9A%D0%B0%D1%80%D0%B0%D0%B3%D0%B0%D0%BD%D0%B4%D0%B8%D0%BD%D1%81%D0%BA%D0%B0%D1%8F+%D0%BE%D0%B1%D0%BB%D0%B0%D1%81%D1%82%D1%8C-karaganda; my_city_2=40_0_0_%D0%A2%D0%B0%D0%BB%D0%B4%D1%8B%D0%BA%D0%BE%D1%80%D0%B3%D0%B0%D0%BD_0_%D0%90%D0%BB%D0%BC%D0%B0%D1%82%D0%B8%D0%BD%D1%81%D0%BA%D0%B0%D1%8F+%D0%BE%D0%B1%D0%BB%D0%B0%D1%81%D1%82%D1%8C_taldykorgan; _gid=GA1.2.875899638.1612447334; laquesis=er-603@b#er-626@b#erm-173@a#erm-201@a; _ym_isad=2; lqstatus=1612638238||||; newrelic_cdn_name=CF; PHPSESSID=0tk6h435j9mspkfsl3l728rbr6; dfp_segment=%5B%5D; pt=" + phoneToken + "; __utma=16996198.929845084.1610551331.1612621398.1612621398.24; __utmt=1; _ym_visorc=w; from_detail=1; _gat_clientNinja=1; tmr_reqNum=254; tmr_detect=0%7C1612637596975; onap=176fc570448x4c7fe010-30-17778aea5abx1585c383-5-1612639403; __utmb=16996198.8.8.1612637603531")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return string(b)
} 

func makeRequest(url string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("site.html")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, string(b))
}

func DownloadFile(path string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Panic(err)
	}
}

func Open_Parse(filename string) *html.Node {
	file, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	
	doc, err := html.Parse(file)
	if err != nil {
		log.Panic(err)
	}

	return doc
}

func ParsePage(filename string) []tData {
	doc := Open_Parse(filename)

	var data []tData
	var currentPrice, currentImg, currentName, currentRef []string
	var collect func(*html.Node)
	collect = func(node *html.Node) {
		if node == nil {
			return
		}

		if node.Data == "a" && node.Type == html.ElementNode{
			for i, att := range node.Attr {
				if att.Key == "class" && strings.Contains(att.Val, "marginright5 link linkWithHash detailsLink") {
					
					currentRef = append(currentRef, node.Attr[i-1].Val)
					
					node = node.FirstChild
					if node == nil {
						break
					}
					node = node.NextSibling
					if node == nil {
						break
					}
					node = node.FirstChild
					if node == nil {
						break
					}
					currentName = append(currentName, node.Data)
				}
			}
		}

		if node.Data == "p" && node.Type == html.ElementNode {
			for _, att := range node.Attr {
				if att.Key == "class" && att.Val == "price" {
					node = node.FirstChild
					if node == nil {
						break
					}

					node = node.NextSibling
					if node == nil {
						break
					}

					node = node.FirstChild
					if node == nil {
						break
					}

					currentPrice = append(currentPrice, node.Data)
				}
			}
		}


		if node.Type == html.ElementNode && node.Data == "img" {
			var flag bool = false
			var tmp string 
			for _, att := range node.Attr {
				if att.Key == "src" { 
					tmp = att.Val
				} else if att.Key == "class" && att.Val == "fleft" {
					flag = true
				}
			}
			if flag {
				currentImg = append(currentImg, tmp)
			}
		}

		collect(node.NextSibling)
		collect(node.FirstChild)
	}
	
	collect(doc)

	var buf tData
	for i := 0; i < len(currentImg); i++ {
		buf.Img = currentImg[i]
		buf.Name = currentName[i]
		buf.Price = currentPrice[i]
		buf.Ref = currentRef[i]
		data = append(data, buf) 
	}

	return data
	
}

func parseMainPage() []tData {
	log.Printf("Parsing main page: begun")
	makeRequest("https://www.olx.kz/dom-i-sad/karaganda/")
	data := ParsePage("site.html")
	log.Printf("Parsing main page: completed")
	return data
}

func parseProductPage(filename string, elem tData) tData {
	doc := Open_Parse(
		filename)

	var s string
	var fields []string

	var collect func(*html.Node)
	collect = func(node *html.Node) {
		if node == nil {
			return
		}

		if node.Type == html.ElementNode && node.Data == "section" {
			for _, attr := range node.Attr {
				if attr.Key == "id" && attr.Val == "body-container" {
					node = node.FirstChild
					if node == nil {
						break
					}

					node = node.NextSibling
					if node == nil {
						break
					}

					node = node.FirstChild
					if node == nil {
						break
					}

					s = node.Data
					fields = strings.Fields(s)
					if len(fields) >= 3 {
						s = fields[3]
						s = strings.Replace(s, "'", "", 1)
						
						s = strings.Replace(s, "';", "", 1)
					}

				}
			}
		}

		if node.Type == html.ElementNode && node.Data == "span" {
			for _, attr := range node.Attr {
				if attr.Key == "span" && attr.Val == "link lheight22 cpointer locationbox__show-desc" {
					log.Print("fgr")

					node = node.NextSibling
					if node == nil {
						break
					}

					node = node.FirstChild
					if node == nil {
						break
					}

					node = node.NextSibling
					if node == nil {
						break
					}

					node = node.NextSibling
					if node == nil {
						break
					}

					elem.Place = node.Data
					log.Printf(node.Data, "\n")
				}
			}
		}

		collect(node.NextSibling)
		collect(node.FirstChild)
	}
	
	collect(doc)

	fileTree, err := os.Create("tree.html")
	if err != nil {
		log.Panic(err)
	}
	defer fileTree.Close()

	var printTree func (*html.Node)
	printTree = func (node *html.Node) {
		if node.Type == html.ElementNode {

		fmt.Fprintf(fileTree, "Тег: %s ", node.Data)

		for _, attr := range node.Attr {
			fmt.Fprintf(fileTree, "Атрибут: %s = %s ", attr.Key, attr.Val)
		}

		fmt.Fprintf(fileTree, "\n")

		if node.NextSibling != nil {
			fmt.Fprintf(fileTree, "NextSibling: ")
			printTree(node.NextSibling)
		} 
		if node.FirstChild != nil {
			fmt.Fprintf(fileTree, "\n    FirstChild: ")
			printTree(node.FirstChild)
			fmt.Fprintf(fileTree, "\n")
		} 
		}
		
	}

	printTree(doc)

	file, err:= os.Open("site.html")
	if err != nil {
		log.Panic(err)
	}
	p, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}
	text := string(p)
	key := `<span class="button br3 fright"><input type="submit" class="submit cfff {id: '`

	var j, val, curval int
	for i, _ := range(text) {
		for j = 0; i+j < len(text) && j < len(key) && byte(text[i+j]) == byte(key[j]); j++ {curval = j}
		if curval == len(key)-1 {
			val = i+curval
			break
		}
	}

	id := text[val+1:val+6]

	log.Printf("%s %s", s, id)

	s = phoneRequest(s, id)

	elem.PhoneNumber = s

	return elem
}

func parsePages(data []tData) []tData {
	log.Printf("Parsing product pages: begun\n")
	for i, elem := range data {
		makeRequest(elem.Ref)
		log.Printf(elem.Ref)
		data[i] = parseProductPage("site.html", elem)
	}
	log.Printf("Parsing product pages: completed\n")

	return data
}
