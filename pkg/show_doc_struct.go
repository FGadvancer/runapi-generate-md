package pkg

// PageContent is a struct for page content parse from page model
type PageContent struct {
	Info struct {
		From        string `json:"from"`
		Type        string `json:"type"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Method      string `json:"method"`
		URL         string `json:"url"`
		Remark      string `json:"remark"`
		APIStatus   string `json:"apiStatus"`
	} `json:"info"`
	Request struct {
		Params struct {
			Mode       string `json:"mode"`
			Urlencoded []struct {
				Name    string `json:"name"`
				Type    string `json:"type"`
				Value   string `json:"value"`
				Require string `json:"require"`
				Remark  string `json:"remark"`
			} `json:"urlencoded"`
			Formdata []struct {
				Name    string `json:"name"`
				Type    string `json:"type"`
				Value   string `json:"value"`
				Require string `json:"require"`
				Remark  string `json:"remark"`
			} `json:"formdata"`
			JSON     string   `json:"json"`
			JSONDesc []Header `json:"jsonDesc"`
		} `json:"params"`
		Headers []Header `json:"headers"`
		Cookies []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"cookies"`
		Auth         []interface{} `json:"auth"`
		Query        []interface{} `json:"query"`
		PathVariable []interface{} `json:"pathVariable"`
	} `json:"request"`
	Response struct {
		ResponseText    string `json:"responseText"`
		ResponseExample string `json:"responseExample"`
		ResponseHeader  struct {
			AccessControlAllowCredentials string `json:"access-control-allow-credentials"`
			AccessControlAllowHeaders     string `json:"access-control-allow-headers"`
			AccessControlAllowMethods     string `json:"access-control-allow-methods"`
			AccessControlAllowOrigin      string `json:"access-control-allow-origin"`
			AccessControlExposeHeaders    string `json:"access-control-expose-headers"`
			AccessControlMaxAge           string `json:"access-control-max-age"`
			ContentLength                 string `json:"content-length"`
			ContentType                   string `json:"content-type"`
			CookieFromServer              string `json:"cookie-from-server"`
			Date                          string `json:"date"`
		} `json:"responseHeader"`
		ResponseStatus         int        `json:"responseStatus"`
		ResponseTime           int        `json:"responseTime"`
		ResponseParamsDesc     []Response `json:"responseParamsDesc"`
		ResponseFailExample    string     `json:"responseFailExample"`
		ResponseFailParamsDesc []Response `json:"responseFailParamsDesc"`
		Remark                 string     `json:"remark"`
		ResponseSize           int        `json:"responseSize"`
	} `json:"response"`
	Scripts struct {
		Pre  string `json:"pre"`
		Post string `json:"post"`
	} `json:"scripts"`
	Extend struct {
	} `json:"extend"`
}
type Header struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
	Require string      `json:"require"`
	Remark  string      `json:"remark"`
}
type Response struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Remark string `json:"remark"`
}
