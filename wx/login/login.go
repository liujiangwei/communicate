package login

import (
	"communicate/wx"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

// 获取登陆信息
func Login(u string) (*Information, *cookiejar.Jar, error) {
	//	https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage
	//	?ticket=A7HyCtAhlvErymvY6y0X97DH@qrticket_0&uuid=odH7dPg35A==&lang=zh_&scan=1559288822&fun=new&version=v2&lang=zh_
	u += "&fun=new&version=v2&lang=zh_"
	log.Printf("get login data[%s]", u)
	req, err := http.NewRequest("Get", u, nil)
	if err != nil {
		return nil,nil, err
	}
	req.Header.Set("UserAgent", wx.UserAgent)
	req.Header.Set("Referer", "https://wx2.qq.com/?&lang=zh_")
	client := http.Client{}
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
	resp, err := client.Do(req)
	if err != nil {
		return nil,nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, errors.New(fmt.Sprintf("get login data failed[status=%s]", resp.Status))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	//<error>
	// <ret>0</ret>
	// <message></message>
	// <skey>@crypt_109e2f66_7666e69626768980a14bf47f2d6b2c37</skey>
	// <wxsid>axd0c0PTHoXypSWG</wxsid>
	// <wxuin>2451660816</wxuin>
	// <pass_ticket>UQ0etxVSTOgXsf1p0w1o3OIGCju5D%2BzWOOka6Ei19szKFpxN6R%2F6MPBaSZB2SYED</pass_ticket>
	// <isgrayscale>1</isgrayscale>
	// </error
	var res Result
	if err := xml.Unmarshal(data, &res); err != nil {
		return nil, nil, err
	}

	if res.Ret != 0{
		return nil, nil, errors.New(fmt.Sprintf("login error [%d] [%s]", res.Ret, res.Message))
	}

	return &Information{
		Skey:        res.Skey,
		Wxsid:       res.Wxsid,
		Wxuin:       res.Wxuin,
		PassTicket:   res.PassTicket,
	}, jar, nil
}

type Result struct {
	Error       xml.Name `xml:"error"`
	Ret         int   `xml:"ret"`
	Message     string   `xml:"message"`
	Skey        string   `xml:"skey"`
	Wxsid       string   `xml:"wxsid"`
	Wxuin       string   `xml:"wxuin"`
	PassTicket  string   `xml:"pass_ticket"`
	Isgrayscale string   `xml:"isgrayscale"`
}

type Information struct {
	Skey        string
	Wxsid       string
	Wxuin       string
	PassTicket  string
}