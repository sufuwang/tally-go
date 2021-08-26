package tool

type TypeUserInfo struct {
	AvatarUrl  string `json:"avatarUrl"`
	NickName   string `json:"nickName"`
	Password   string `json:"password"`
	Signature  string `json:"signature"`
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
}

type TypeUserInfoMysql struct {
	Id        int    `json:"id"`
	NickName  string `json:"nickName"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatarUrl"`
	Signature string `json:"signature"`
	Token     string `json:"token"`
}
