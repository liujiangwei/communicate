package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)
var client = http.Client{}
var codeFile = "code.png"
var UserAgent = "User-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"

func init() {
	cookieJar, _ := cookiejar.New(nil)
	client.Jar = cookieJar
}

type WX struct {
	client http.Client
	jar cookiejar.Jar
}

func New(){
	wx := &WX{}
	wx.client = http.Client{}
	cookieJar, _ := cookiejar.New(nil)
	wx.client.Jar = cookieJar
}

func main() {
	for ; ; {
		uuid, err := getUUid()
		if err != nil{
			log.Fatal(err)
		}
		if err := getCode(codeFile, uuid); err != nil{
			log.Fatal(err)
		}

		if err = login(uuid, "1"); err != nil{
			log.Println(err)
		}else{
			break
		}
	}

	//_ = initWx()
}

//tip: 1 未扫描 0 已扫描
func login(uuid string, tip string)  error{
	var u ="https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid="+ uuid +"&tip="+tip+"&r=-25471707&_=1559098319558"
	log.Println("get login status", u)
	resp, err := client.Get(u)
	if err != nil{
		return err
	}
	if resp.StatusCode != http.StatusOK{
		return errors.New(fmt.Sprintf("get login status failed[status=%s]", resp.Status))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	str := strings.TrimSpace(string(data))
	pos := strings.Index(str, ";")
	code := str[0:pos]

	switch code {
	case "window.code=408":
		log.Println("清扫描二维码")
		return login(uuid, "0")
	case "window.code=400":
		log.Println("二维码已过期， 请刷新")
		return errors.New("code is expires")
	case "window.code=201":
		//window.code = 201;
		//window.userAvatar = 'data:img/jpg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAoHBwgHBgoICAgLCgoLDhgQDg0NDh0VFhEYIx8lJCIfIiEmKzcvJik0KSEiMEExNDk7Pj4+JS5ESUM8SDc9Pjv/2wBDAQoLCw4NDhwQEBw7KCIoOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozv/wAARCACEAIQDASIAAhEBAxEB/8QAHAAAAgIDAQEAAAAAAAAAAAAAAAECBQMEBgcI/8QAPBAAAQQBAgQEBAMGAwkAAAAAAQACAwQRBSEGEjFBE1FhcRQygZEiUqEHFSNCcsEXQ6JigpKxstLh8PH/xAAYAQEBAQEBAAAAAAAAAAAAAAAAAQIEA//EABwRAQEBAAIDAQAAAAAAAAAAAAABEQIhAxIxcf/aAAwDAQACEQMRAD8A8dCB1QEKIaEJhAJ4TBTwgjhCtuHKda/xBQqXA7wJ7DI5OU4OC4A4K1dTqClqNmqDkQyuYCfQ4QaKFIjdJAkk0IEkmhAkJ4QUzAkIQgaaE/JAAJ4TAUsIpBSASwptblBv6PIYNUpTdPDna7PsQtjimMM4o1RvlblH+orHUquMcMoBIEjt/YNP91tcVt5uKNUd525D/qKq6oC1Y8LYLDjosZaojEkpkKOUQkJpIEhNJAITQgApBRCkFRMBTaCTgdSosBJAHVWuj6TY1aSSOq3mliYZOXuQNzj6o1Fc1hcdgrSppU1+xBFXicXTtHKGjO/Qn7glYpaUley+N7HZZIWE+oJH9l6Lw4791UtNttghgl8SSWKwx+RIMNa5rm7kdN8bdUpiWm8HWK3B8ti1AYZcuzzD5WjGSP8AhA9iVxXErc8Q6gR0Nh//AFFe4X45+IeGJaUbXeNy84e5oId9PXdecWeHY7PF3wll3K2y4lzi35O56dcYKrUmx589mB0Wu5u66HXaNWtcfHRdJJA04a97cE47qlfGQsmNUtWMtwthzVjcNkZYiFFTKiUZRQhCBIQhUZQxMMKiCphxQMNOV2XBEkjdTikhZz3WP/gHBJDiOpAwSNtx69u/Is5iMjt+ittE1B2n347DcMwfn5Q4tPY7/dGo9q1fRresaSyM6e2tZnPJI3Ygn87T2x+LbbPMey1NI4IMDtPhuzPa2AudI0DA5sjABPbp9lfcFX9Q1KiZdSfJ4jDytJOz8bE+u+fv7Lo3fDmZrnCPxCcNJxn6Km2dMTa7Wzh0bAzGG7bDlGdlqWNDhmEg8KF3M1wY5zN252O47eitQhQlseY8SaJRqvZVhoOJYMeNLkjGejQdgFwXE3D40wiQM5PNoOR7hfQdyvVtwGKy1jmO/MvM+O+DLngufRBkgJzg78o8luZa7OFnk4+k+vH3M5n4HdYZG4JGQfZb1ulYrOc2WJwIOM9louBWHLy43jcrAQokKbhuolRhDCSkkUQkIQgmFsVq5suMcZzKccjAPnPkPX/mtYFbFSEzztaHtjbkc0jzhrB5kqq39H0TUNZvfCUa75Jg0u5cYxjqrvTOEbNkOdanr0msOC+R+W58nEZDT7kK0vce1KVKKvo0LZ7rWBk2pTR8rpTjc8vQ+7sk4yFRw8UGW2LWpVG33tGGCWV4aw+Ya0gY9Feklj17hDTPhZq9iGy+xL4Qhle35Q1oJHfqcAHHdqsJL8d6874KYHkkd4skQ5Q1zBk5ee2MD1wVRcN/tLoX6ppuosrztjxhuzHDo0Z7bkDfplYLnEmlw8I3YdKryw+NJyyRsdyvgdn8QJJPcbdBvtujXxvs45Yx7Gx3Y3PfMQ4vBDcefoM9PZdRd1hrZuSFjpgx3LK1oyOQgZd9M/oV8/xWOa9GOdxJkGQ7r17rpptbf8Vqz3WZo2RPDh4R33GMYyO4bn0BUbyOh1LXTV065LPYmEzJQz4eMnl74PN2HfIyDjsqj/EvVZa744Z4XPk+eGyOUD+h2R9c/Rcvb4lls6ZBWL+V1fmLXEZMhdgYd54GeqqCxtzeuOWXvD+b+nz9uvv2ur7O8g4g0a9EKuv6RYqzn/Phbloz0JB7LndU0nTbLHT6XM2dpOA1ow8e7eqqafEmp6bTlpQT/wACUgujeA5vfsenUqskuTGUSh/I8HILAGkfdN16Tz8rLOfbYn0q1CZPEgcwx/MHDlI+hVe5rVsTajbsPY+ed8r2dHPPMf1WZ7zYjfatwsaDsHsaGcx8gBt+n/nLnuX40BFzNLsgAdSoFim9wcdth2CgjKOEJ5KECUgTjGdj2UQmEGRpWQOWEFSBRXQcPRGWLUXRukMzKjnMYxwGcEEk774xnHotePU5qd2SWB5DZOrDuC09iD122wq+rdlpyF8TiC5vKfVp6hYi8nB8tldXcXtdzbNqF9bDQJWudCdy3fctJ3I9Oo9eq29PlFjXrsD3bTh0e/bJAJ+gyfoufozFlyFwJBa8HI7bq84dqT65xMIq8TiZnPY9zRhoDgW5J6DqkW1Ry80cjmPBa5pwQexWPnIOQcFdNx9oMuhauH2HAy3Gictb0bn5vf8AFn6rky5SJKtGTQaliO28RWD8tjGzvR//AHDfzz1OPU9C1LSZxFcqyRuLQ5pxkOae4I2IVcHkHIO5XccAXtRtarWgntu+BjeGFr3Egh38gHke4O3oSqjn9M0ytHTdq2qOLarHcsUIOH2X/lHk0dz9OpVbfvS3p/EeGsa0YZGwYaweQXrGvnhXULwj0+jTdPylkEc0hjbIASOVjiSxhyD+EtGc5zuvONcIiuOrvoSUJIzh0MkUZIPuGtS9QUaMrKOYBxEvL2xuM/dYiogyhJCATCSaBgp5Ucp5RTypAnCgnnZEZYnFsjT5ELNBfs1nNdDM9nI7maGnAB81rNO6WUVa6pxFqWtQxR6lZdYMORG9+5AJzj/3+wVVhIlJA+ytNN1V1W/ULSWxRytcWjyyM598foPNVWdkNzzbdRsiNr4p8rHtkkOeYvbnz2z9wB9gun0XW4dYbBpHENd9uDIjgtsH8av5DP8AM3/ZP0wuSgi8eQRtc1rj05jgH69l6Pwiyxpb6bLtSKj4JdM+V8eJJ4huckD5AB/vEgZ8kHWWv2OaUNLObRbYbDgykfh5gR+LHsN15fqfBv7smlZPq9LnZu1jS4lw7YyMb+69vscX0HarT061L4PjweK6KZnKQCMjcHYgE57bBcH+0GzotC7YgqUxE4xN8J1eTw2SHHRwHUYO3/1W9M/ryR+A4gHIHdCHEE9CEKNEhJNA0JZQgaEkIJAoykEsop5QkhECEIygkwgPBPReh2eOK0fBMFEl1q69pgdloDYY+wHrv19PRecqRdn2RLNbtmzYcY5nSukAaWseSdvT09lCS/YsV2wzSl7YxhnMc4HXH3yfqfNYq1p9Zzi1rHtcMOY9vM0/RFidlh5k8JkRwMNjGG/ZTFYScnJQooVAE0IQCEIQCEIQMJBCEAhCEAgoQgEkIQNJCEAhCEH/2Q==';
		str = strings.TrimSpace(str[pos:])
		log.Println("userAvatar", str[strings.Index(str, "'")+1: len(str)-2])
		log.Println("请确认登陆")
		return login(uuid, "0")
	case "window.code=200":
		log.Println("登陆成功")
		str = strings.TrimSpace(str[pos:])
		//window.redirect_uri="https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage?ticket=xxx&uuid=xxx&lang=xxx&scan=xxx";
		_  = ioutil.WriteFile("login", data, 0666)
		if wd, err := getWxData(str[strings.Index(str, "\"")+1: len(str)-2]);err != nil{
			return err
		}else{

		}
	}

	return errors.New(fmt.Sprintf("error response[%s]", string(data)))
}

func initWx() error{
	//var url = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxinit?pass_ticket=xxx&skey=xxx&r=xxx"
	var u = "https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxinit?r=-212339180"
	resp, err := client.Post(u, "text/plain", nil)
	if err != nil{
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK{
		return errors.New(fmt.Sprintf("get login status failed[status=%s]", resp.Status))
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	log.Println(string(data))
	return nil
}

func getWxData(u string) (*WxData, error){
	//	https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage
	//	?ticket=A7HyCtAhlvErymvY6y0X97DH@qrticket_0&uuid=odH7dPg35A==&lang=zh_&scan=1559288822&fun=new&version=v2&lang=zh_
	u += "&fun=new&version=v2&lang=zh_"
	log.Println("get login ", fmt.Sprintf("[%s]", u))
	req, err := http.NewRequest("Get", u, nil)
	if err != nil{
		return nil, err
	}
	req.Header.Set("UserAgent", UserAgent)
	req.Header.Set("Referer", "https://wx2.qq.com/?&lang=zh_")
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}

	if resp.StatusCode != http.StatusOK{
		return nil, errors.New(fmt.Sprintf("get login status failed[status=%s]", resp.Status))
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	//<error>
	// <ret>0</ret>
	// <message></message>
	// <skey>@crypt_109e2f66_7666e69626768980a14bf47f2d6b2c37</skey>
	// <wxsid>axd0c0PTHoXypSWG</wxsid>
	// <wxuin>2451660816</wxuin>
	// <pass_ticket>UQ0etxVSTOgXsf1p0w1o3OIGCju5D%2BzWOOka6Ei19szKFpxN6R%2F6MPBaSZB2SYED</pass_ticket>
	// <isgrayscale>1</isgrayscale>
	// </error>
	_  = ioutil.WriteFile("login_page", data, 0666)

	var wd WxData
	if err := xml.Unmarshal(data, &wd); err != nil{
		return nil, err
	}

	return &wd, nil
}

type WxData struct {
	Error xml.Name `xml:"error"`
	Ret string `xml:"ret"`
	Message string `xml:"message"`
	Skey string `xml:"skey"`
	Wxsid string `xml:"wxsid"`
	Wxuin string `xml:"wxuin"`
	PassTicket string `xml:"pass_ticket"`
	Isgrayscale string `xml:"isgrayscale"`
}


func getCode(codeFile string, uuid string) error{
	var u= "https://login.weixin.qq.com/qrcode/" + uuid
	log.Println("get code", u)
	resp, err := client.Get(u)
	if err != nil{
		return err
	}

	if resp.StatusCode != http.StatusOK{
		return errors.New("http response error")
	}

	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	return ioutil.WriteFile(codeFile,data,0666)
}

func getUUid() (string, error){
	var u = "https://login.wx.qq.com/jslogin?" +
		"appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_&_="
	log.Println("get uuid", u)
	resp, err := client.Get(u)
	if err != nil{
		return "", err
	}

	if resp.StatusCode != http.StatusOK{
		return "", errors.New("http response error")
	}

	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}
	fields := strings.Split(string(data), "\"")
	if len(fields) < 2{
		return "", errors.New("get uuid error")
	}

	return fields[1], nil
}
