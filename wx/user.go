package wx

//"Uin": 2451660816,
//"UserName": "@acbedbfbdce2f4f6a6ee8d42d497b4355f3e8fe14411605c0d27d4f533416724",
//"NickName": "风萧萧，易水寒",
//"HeadImgUrl": "/cgi-bin/mmwebwx-bin/webwxgeticon?seq=1709249277&username=@acbedbfbdce2f4f6a6ee8d42d497b4355f3e8fe14411605c0d27d4f533416724&skey=@crypt_109e2f66_249ccb0a9646bfa0f0f8e1a28283b463",
//"RemarkName": "",
//"PYInitial": "",
//"PYQuanPin": "",
//"RemarkPYInitial": "",
//"RemarkPYQuanPin": "",
//"HideInputBarFlag": 0,
//"StarFriend": 0,
//"Sex": 1,
//"Signature": "",
//"AppAccountFlag": 0,
//"VerifyFlag": 0,
//"ContactFlag": 0,
//"WebWxPluginSwitch": 0,
//"HeadImgFlag": 1,
//"SnsFlag": 1
type User struct {
	Uin               int    `json:"Uin"`
	UserName          string `json:"UserName"`
	NickName          string `json:"NickName"`
	HeadImgUrl        string `json:"HeadImgUrl"`
	RemarkName        string `json:"RemarkName"`
	PYInitial         string `json:"PYInitial"`
	PYQuanPin         string `json:"PYQuanPin"`
	RemarkPYInitial   string `json:"RemarkPYInitial"`
	RemarkPYQuanPin   string `json:"RemarkPYQuanPin"`
	HideInputBarFlag  int    `json:"HideInputBarFlag"`
	StarFriend        int    `json:"StarFriend"`
	Sex               int    `json:"Sex"`
	Signature         string `json:"Signature"`
	AppAccountFlag    int    `json:"AppAccountFlag"`
	VerifyFlag        int    `json:"VerifyFlag"`
	ContactFlag       int    `json:"ContactFlag"`
	WebWxPluginSwitch int    `json:"WebWxPluginSwitch"`
	HeadImgFlag       int    `json:"HeadImgFlag"`
	SnsFlag           int    `json:"SnsFlag"`
}

/*
Alias: ""
AppAccountFlag: 0
AttrStatus: 0
ChatRoomId: 0
City: ""
ContactFlag: 0
DisplayName: ""
EncryChatRoomId: ""
HeadImgUrl: "/cgi-bin/mmwebwx-bin/webwxgetheadimg?seq=0&username=@@c34f85604efc35cbcf0d09b382652f65667f7246bcca8977a0b8cb261c385d8a&skey=@crypt_109e2f66_a9f4da30dec3e6ac7e236038540b831f"
HideInputBarFlag: 0
IsOwner: 0
KeyWord: ""
MemberCount: 228
MemberList: [,…]
NickName: "æ’é€šå•†åŠ¡å›­ä¹°èœç¾¤-æ¯æ—¥ä¼˜é²œDSZ"
OwnerUin: 0
PYInitial: ""
PYQuanPin: ""
Province: ""
RemarkName: ""
RemarkPYInitial: ""
RemarkPYQuanPin: ""
Sex: 0
Signature: ""
SnsFlag: 0
StarFriend: 0
Statues: 0
Uin: 0
UniFriend: 0
UserName: "@@c34f85604efc35cbcf0d09b382652f65667f7246bcca8977a0b8cb261c385d8a"
VerifyFlag: 0
*/
type Member struct {
	Alias            string   `json:"Alias"`
	AppAccountFlag   int      `json:"AppAccountFlag"`
	AttrStatus       int      `json:"AttrStatus"`
	ChatRoomId       int      `json:"ChatRoomId"`
	City             string   `json:"City"`
	ContactFlag      int      `json:"ContactFlag"`
	DisplayName      string   `json:"DisplayName"`
	EncryChatRoomId  string   `json:"EncryChatRoomId"`
	HeadImgUrl       string   `json:"HeadImgUrl"`
	HideInputBarFlag int      `json:"HideInputBarFlag"`
	IsOwner          int      `json:"IsOwner"`
	KeyWord          string   `json:"KeyWord"`
	MemberCount      int      `json:"MemberCount"`
	MemberList       []Member `json:"MemberList"`
	NickName         string   `json:"NickName"`
	OwnerUin         int      `json:"OwnerUin"`
	PYInitial        string   `json:"PYInitial"`
	PYQuanPin        string   `json:"PYQuanPin"`
	Province         string   `json:"Province"`
	RemarkName       string   `json:"RemarkName"`
	RemarkPYInitial  string   `json:"RemarkPYInitial"`
	RemarkPYQuanPin  string   `json:"RemarkPYQuanPin"`
	Sex              int      `json:"Sex"`
	Signature        string   `json:"Signature"`
	SnsFlag          int      `json:"SnsFlag"`
	StarFriend       int      `json:"StarFriend"`
	Statues          int      `json:"Statues"`
	Uin              int      `json:"Uin"`
	UniFriend        int      `json:"UniFriend"`
	UserName         string   `json:"UserName"`
	VerifyFlag       int      `json:"VerifyFlag"`
}
