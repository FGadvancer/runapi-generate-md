package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ruapi-generate-md/pkg"
	"github.com/ruapi-generate-md/pkg/db"
	"github.com/ruapi-generate-md/pkg/db/model"
	"os"
	"strings"
	"sync"
)

var data *db.DataBase

func GeneratePageByItemID(outPath string, projectName string) {
	var wg sync.WaitGroup
	dbFileName := "/showdoc_data/html/Sqlite" + "/showdoc.db.php"
	//dbFileName := "./showdoc.db.php"
	data = db.NewDataBase(dbFileName)
	data.Init()
	dir := outPath + "/" + projectName
	item, err := data.Item.TakeItem(projectName)
	if err != nil {
		panic("not found this project:" + projectName)
	}
	headerArgs, err := data.RunapiGlobalParam.TakeRunapiGlobalHeaderParam(item.ItemId)
	if err != nil {
		panic("not found this header")
	}
	globalHeader := getGlobalHeader(headerArgs.ContentJsonStr)
	catalogs, _ := data.Catalog.TakeCatalogs(item.ItemId)
	var ni sql.NullInt32
	ni.Int32 = 0
	ni.Valid = true
	pages, _ := data.Page.TakePages(ni, item.ItemId)
	for _, page := range pages {
		//fmt.Print(i, page.PageContent.String, globalHeader, page.PageTitle.String)
		generateOnePageMarkDown(page.PageContent.String, globalHeader, page.PageTitle.String, dir)
	}
	for _, catalog := range catalogs {
		newCataLog := *catalog
		fmt.Println(newCataLog.CatName.String, newCataLog.CatId.Int32)
		wg.Add(1)
		go func() {
			//tempPath := dir + "/" + newCataLog.CatName.String
			//if _, err := os.Stat(tempPath); os.IsNotExist(err) {
			//	if err := os.MkdirAll(tempPath, 0755); err != nil {
			//		fmt.Printf("mkdir %s failed：%s\n", dir, err)
			//		return
			//	}
			//}
			//pages, _ := data.Page.TakePages(newCataLog.CatId, item.ItemId)
			//for _, page := range pages {
			//	//fmt.Print(i, page.PageContent.String, globalHeader, page.PageTitle.String)
			//	generateOnePageMarkDown(page.PageContent.String, globalHeader, page.PageTitle.String, tempPath)
			//}
			recursionGen(dir, newCataLog, globalHeader)

			wg.Done()
		}()
	}
	wg.Wait()
}
func recursionGen(dir string, newCataLog model.Catalog, globalHeader []pkg.Header) {
	tempPath := dir + "/" + newCataLog.CatName.String
	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		if err := os.MkdirAll(tempPath, 0755); err != nil {
			fmt.Printf("mkdir %s failed：%s\n", dir, err)
			return
		}
	}
	pages, _ := data.Page.TakePages(newCataLog.CatId, newCataLog.ItemId)
	for _, page := range pages {
		//fmt.Print(i, page.PageContent.String, globalHeader, page.PageTitle.String)
		generateOnePageMarkDown(page.PageContent.String, globalHeader, page.PageTitle.String, tempPath)
	}
	catalogs, _ := data.Catalog.TakeSubCatalogs(newCataLog.ItemId, newCataLog.CatId)
	if len(catalogs) == 0 {
		return
	} else {
		for _, catalog := range catalogs {
			recursionGen(tempPath, *catalog, globalHeader)
		}
	}

}

func getGlobalHeader(jsonStr string) []pkg.Header {
	//jsonStr := "[{&quot;name&quot;:&quot;operationID&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;testGordon1111&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;operationID用于全局链路追踪&quot;},{&quot;name&quot;:&quot;token&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJvcGVuSU1BZG1pbiIsIlBsYXRmb3JtIjoiIiwiZXhwIjoxNjg2NzEyMzkwLCJuYmYiOjE2Nzg5MzYwOTAsImlhdCI6MTY3ODkzNjM5MH0.Pr4NqWCVF54oSzFePAU0gFWDHrBGzyEnb_K1BymnC88&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;管理员或者用户token&quot;}]"
	jsonStr = strings.ReplaceAll(jsonStr, "&quot;", "\"")
	data := []pkg.Header{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}
	return data

}
