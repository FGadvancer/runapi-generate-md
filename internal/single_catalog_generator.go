package internal

import (
	"encoding/json"
	"fmt"
	"github.com/ruapi-generate-md/pkg"
	"github.com/ruapi-generate-md/pkg/constant"
	"os"
	"strings"
)

var successMap = make(map[string]pkg.Response)

func codeAreaWrite(sb *strings.Builder, codeArea string) {
	sb.WriteString("```json\n" + codeArea +
		"\n```" + "\n")
}
func funcStatus(sb *strings.Builder, status string) {
	if status == "0" {
		return
	}
	titleWrite(sb, "接口状态", 4)
	switch status {
	case "1":
		valueCircleWrite(sb, "开发中", "无")
	case "2":
		valueCircleWrite(sb, "测试中", "无")
	case "3":
		valueCircleWrite(sb, "已完成", "无")
	case "4":
		valueCircleWrite(sb, "需修改", "无")
	case "5":
		valueCircleWrite(sb, "已废弃", "无")

	}

}
func titleWrite(sb *strings.Builder, title string, level int) {
	var symbol string
	switch level {
	case 1:
		symbol = "#"
	case 2:
		symbol = "##"
	case 3:
		symbol = "###"
	case 4:
		symbol = "####"
	case 5:
		symbol = "#####"
	}
	sb.WriteString(symbol + " " + title + "\n")
}
func pageBigTitleWrite(sb *strings.Builder, title string, level int) {
	sb.WriteString("\n<center>\n\n")
	titleWrite(sb, title, level)
	sb.WriteString("\n</center>\n\n")
}
func valueCircleWrite(sb *strings.Builder, value, defalut string) {
	if value == "" {
		value = defalut
	}
	sb.WriteString("- " + value + "\n")
}
func valueURLWrite(sb *strings.Builder, value, defalut string) {
	if value == "" {
		value = defalut
	}
	sb.WriteString("- " + "`" + value + "`" + " \n")
}
func headerTableWrite(sb *strings.Builder, data []pkg.Header) {
	head := pkg.Header{Type: "类型", Name: "header名", Value: "示例值", Require: "必选", Remark: "说明"}
	data = append([]pkg.Header{head}, data...)
	for i, v := range data {
		if v.Name == "" {
			continue
		}
		sb.WriteString("|" + v.Name)
		sb.WriteString("|" + v.Value)
		if v.Require == "1" {
			v.Require = "必填"
		} else {
			v.Require = "选填"
		}
		sb.WriteString("|" + v.Require)
		sb.WriteString("|" + v.Type)
		sb.WriteString("|" + v.Remark + "|\n")
		if i == 0 {
			sb.WriteString("|:----    |:-------    |:--- |---|------      |\n")
		}
	}
}
func reqTableWrite(sb *strings.Builder, data []pkg.RequestHeader) {
	head := pkg.RequestHeader{Type: "类型", Name: "字段名", Require: "必选", Remark: "说明"}
	data = append([]pkg.RequestHeader{head}, data...)
	for i, v := range data {
		sb.WriteString("|" + v.Name)
		if v.Require == "1" {
			v.Require = "必填"
		} else {
			v.Require = "选填"
		}
		sb.WriteString("|" + v.Require)
		sb.WriteString("|" + v.Type)
		sb.WriteString("|" + v.Remark + "|\n")
		if i == 0 {
			sb.WriteString("|:----    |:-------    |:--- |---|\n")
		}
	}
}

func successRespTableWrite(sb *strings.Builder, data []pkg.Response) {
	head := pkg.Response{Type: "类型", Name: "参数名", Remark: "说明"}
	for i, v := range data {
		if newValue, ok := successMap[v.Name]; ok {
			data[i] = newValue
		}
	}
	data = append([]pkg.Response{head}, data...)
	for i, v := range data {
		sb.WriteString("|" + v.Name)
		sb.WriteString("|" + v.Type)
		sb.WriteString("|" + v.Remark + "|\n")
		if i == 0 {
			sb.WriteString("|:----    |:-------    |:--- |\n")
		}
	}
}

func failedRespTableWrite(sb *strings.Builder, _ []pkg.Response) {
	head := pkg.Response{Type: "类型", Name: "参数名", Remark: "说明"}
	data := []pkg.Response{
		{Name: "errCode", Type: "int", Remark: "错误码,具体查看全局错误码文档"},
		{Name: "errMsg", Type: "string", Remark: "错误简要信息"},
		{Name: "errDlt", Type: "errDlt", Remark: "错误详细信息"},
	}
	data = append([]pkg.Response{head}, data...)
	for i, v := range data {
		sb.WriteString("|" + v.Name)
		sb.WriteString("|" + v.Type)
		sb.WriteString("|" + v.Remark + "|\n")
		if i == 0 {
			sb.WriteString("|:----    |:-------    |:--- |\n")
		}
	}
}
func IsContain(target string, List []string) bool {

	for _, element := range List {

		if target == element {
			return true
		}
	}
	return false

}
func generateOnePageMarkDown(jsonStr string, globalHeader []pkg.Header, bigTile string, catalogPath string, x bool) {
	//jsonStr := "{&quot;info&quot;:{&quot;from&quot;:&quot;runapi&quot;,&quot;type&quot;:&quot;api&quot;,&quot;title&quot;:&quot;&quot;,&quot;description&quot;:&quot;获取用户token&quot;,&quot;method&quot;:&quot;post&quot;,&quot;url&quot;:&quot;{{host}}/auth/user_token&quot;,&quot;remark&quot;:&quot;&quot;,&quot;apiStatus&quot;:&quot;3&quot;},&quot;request&quot;:{&quot;params&quot;:{&quot;mode&quot;:&quot;json&quot;,&quot;urlencoded&quot;:[{&quot;name&quot;:&quot;&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;&quot;}],&quot;formdata&quot;:[{&quot;name&quot;:&quot;&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;&quot;}],&quot;json&quot;:&quot;{\\n  \\&quot;userID\\&quot;: \\&quot;635204331\\&quot;,\\n  \\&quot;platformID\\&quot;: 1,\\n  \\&quot;secret\\&quot;:\\&quot;openIM123\\&quot;\\n}&quot;,&quot;jsonDesc&quot;:[{&quot;name&quot;:&quot;userID&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;,&quot;value&quot;:&quot;&quot;,&quot;require&quot;:&quot;1&quot;},{&quot;name&quot;:&quot;platformID&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;,&quot;value&quot;:&quot;&quot;,&quot;require&quot;:&quot;1&quot;},{&quot;name&quot;:&quot;secret&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;,&quot;value&quot;:&quot;&quot;,&quot;require&quot;:&quot;1&quot;}]},&quot;headers&quot;:[{&quot;name&quot;:&quot;operationID&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;123123123&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;operationID用于全局链路追踪&quot;}],&quot;cookies&quot;:[{&quot;name&quot;:&quot;&quot;,&quot;value&quot;:&quot;&quot;}],&quot;auth&quot;:[],&quot;query&quot;:[{&quot;name&quot;:&quot;&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;&quot;}],&quot;pathVariable&quot;:[]},&quot;response&quot;:{&quot;responseText&quot;:&quot;{\\n  \\&quot;errCode\\&quot;: 0,\\n  \\&quot;errMsg\\&quot;: \\&quot;\\&quot;,\\n  \\&quot;errDlt\\&quot;: \\&quot;\\&quot;,\\n  \\&quot;data\\&quot;: {\\n    \\&quot;token\\&quot;: \\&quot;eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2MzUyMDQzMzEiLCJQbGF0Zm9ybSI6IklPUyIsImV4cCI6MTY4NjU2MTkxMSwibmJmIjoxNjc4Nzg1NjExLCJpYXQiOjE2Nzg3ODU5MTF9.uRSAZE_Z-FlBFjzuLp4Usy2utT-BQLR1eIzeaS9nCJk\\&quot;,\\n    \\&quot;expireTimeSeconds\\&quot;: 90\\n  }\\n}&quot;,&quot;responseOriginal&quot;:{&quot;errCode&quot;:0,&quot;errMsg&quot;:&quot;&quot;,&quot;errDlt&quot;:&quot;&quot;,&quot;data&quot;:{&quot;token&quot;:&quot;eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2MzUyMDQzMzEiLCJQbGF0Zm9ybSI6IklPUyIsImV4cCI6MTY4NjU2MTkxMSwibmJmIjoxNjc4Nzg1NjExLCJpYXQiOjE2Nzg3ODU5MTF9.uRSAZE_Z-FlBFjzuLp4Usy2utT-BQLR1eIzeaS9nCJk&quot;,&quot;expireTimeSeconds&quot;:90}},&quot;responseExample&quot;:&quot;{\\n  \\&quot;errCode\\&quot;: 0,\\n  \\&quot;errMsg\\&quot;: \\&quot;\\&quot;,\\n  \\&quot;errDlt\\&quot;: \\&quot;\\&quot;,\\n  \\&quot;data\\&quot;: {\\n    \\&quot;token\\&quot;: \\&quot;eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2MzUyMDQzMzEiLCJQbGF0Zm9ybSI6IklPUyIsImV4cCI6MTY4NjU2MTI2OCwibmJmIjoxNjc4Nzg0OTY4LCJpYXQiOjE2Nzg3ODUyNjh9.bypHmiNJtIcpxxCOcA-cWNe0ymyp4-Sa80pmtpF5G0s\\&quot;,\\n    \\&quot;expireTimeSeconds\\&quot;: 90\\n  }\\n}&quot;,&quot;responseHeader&quot;:{&quot;access-control-allow-credentials&quot;:&quot;false&quot;,&quot;access-control-allow-headers&quot;:&quot;*&quot;,&quot;access-control-allow-methods&quot;:&quot;*&quot;,&quot;access-control-allow-origin&quot;:&quot;*&quot;,&quot;access-control-expose-headers&quot;:&quot;Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar&quot;,&quot;access-control-max-age&quot;:&quot;172800&quot;,&quot;content-length&quot;:&quot;277&quot;,&quot;content-type&quot;:&quot;application/json&quot;,&quot;cookie-from-server&quot;:&quot;&quot;,&quot;date&quot;:&quot;Tue, 14 Mar 2023 09:25:11 GMT&quot;},&quot;responseStatus&quot;:200,&quot;responseTime&quot;:162,&quot;responseParamsDesc&quot;:[{&quot;name&quot;:&quot;errCode&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;errMsg&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;errDlt&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;data&quot;,&quot;type&quot;:&quot;object&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;token&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;expireTimeSeconds&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;}],&quot;responseFailExample&quot;:&quot;{\\n  \\&quot;errCode\\&quot;: 10000,\\n  \\&quot;errMsg\\&quot;: \\&quot;: [12]unknown service OpenIMServer.pbAuth.Auth\\&quot;,\\n  \\&quot;errDlt\\&quot;: \\&quot;\\&quot;,\\n  \\&quot;data\\&quot;: null\\n}&quot;,&quot;responseFailParamsDesc&quot;:[{&quot;name&quot;:&quot;errCode&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;errMsg&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;errDlt&quot;,&quot;type&quot;:&quot;string&quot;,&quot;remark&quot;:&quot;&quot;},{&quot;name&quot;:&quot;data&quot;,&quot;type&quot;:&quot;object&quot;,&quot;remark&quot;:&quot;&quot;}],&quot;remark&quot;:&quot;&quot;,&quot;responseSize&quot;:0},&quot;scripts&quot;:{&quot;pre&quot;:&quot;&quot;,&quot;post&quot;:&quot;&quot;},&quot;extend&quot;:{}}"
	jsonStr = strings.ReplaceAll(jsonStr, "&quot;", "\"")
	//fmt.Println("11111", jsonStr)
	if _, ok := constant.Exclude[bigTile]; ok {
		fmt.Println("exclude:", bigTile)
		return
	}
	data := pkg.PageContent{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Unmarshal err", err, "catlogPath", catalogPath, "bigTile", bigTile)
		panic(err)
	}
	var sb strings.Builder
	if !x {
		sb.WriteString("<!-- 使用表格样式 -->\n" +
			"<style>\n" +
			"th {\n    background-color: #1E90FF; /* 设置表头背景颜色 */\n}\n.highlight {\n    background-color: black;\n    color: white;\n    font-family: Consolas, Monaco, 'Andale Mono', 'Ubuntu Mono', monospace;\n}\n</style>" + "\n\n")
	}
	sb.WriteString("---\ntitle: " + bigTile + "\nhide_title: true\n---\n\n")
	pageBigTitleWrite(&sb, bigTile, 2)
	//fmt.Println("1", data.Request.Params.JSON, data.Request.Headers)
	titleWrite(&sb, "简要描述", 3)
	valueCircleWrite(&sb, data.Info.Description, "无")
	funcStatus(&sb, data.Info.APIStatus)
	titleWrite(&sb, "请求方式", 3)
	valueURLWrite(&sb, data.Info.Method, "无")

	titleWrite(&sb, "请求URL", 3)
	valueURLWrite(&sb, data.Info.URL, "无")
	sb.WriteString("\n\n")
	titleWrite(&sb, "Header", 3)
	headerTableWrite(&sb, append(globalHeader, data.Request.Headers...))
	sb.WriteString("\n\n")
	titleWrite(&sb, "请求参数示例\n\n", 3)
	codeAreaWrite(&sb, data.Request.Params.JSON)
	reqTableWrite(&sb, data.Request.Params.JSONDesc)
	titleWrite(&sb, "成功返回示例\n\n", 3)
	codeAreaWrite(&sb, data.Response.ResponseExample)
	titleWrite(&sb, "成功返回示例的参数说明\n\n", 3)

	successRespTableWrite(&sb, data.Response.ResponseParamsDesc)

	titleWrite(&sb, "失败返回示例\n\n", 3)
	codeAreaWrite(&sb, data.Response.ResponseFailExample)
	titleWrite(&sb, "失败返回示例的参数说明\n\n", 3)
	failedRespTableWrite(&sb, data.Response.ResponseFailParamsDesc)
	err = os.WriteFile(catalogPath+"/"+tileChanged(bigTile)+".mdx", []byte(sb.String()), 0644)
	//err = ioutil.WriteFile(bigTile+".md", []byte(sb.String()), 0644)
	if err != nil {
		fmt.Println("WriteFile err", err, "catlogPath", catalogPath, "bigTile", bigTile)
	}
}
func tileChanged(bigTile string) string {
	if v, ok := constant.FileName[bigTile]; ok {
		return v
	} else {
		return bigTile
	}
}
