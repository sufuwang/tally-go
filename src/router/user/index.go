package user

import (
	"fmt"
	"net/http"
	"tally-go/src/database/mysql"
	"tally-go/src/tool"
)

type TypeRouters map[string]http.HandlerFunc

func RegisterUserRouter() TypeRouters {
	//http.HandleFunc("/test", handleTest)
	//http.HandleFunc("/login", handleLogin)
	//http.HandlerFunc("/register", handleRegister)
	// http.Handle("/register", Middleware.CORSMiddleWare(http.HandlerFunc(handleRegister)))
	ds := make(TypeRouters)
	ds["/isExistedNickName"] = handleIsExistedNickName
	ds["/register"] = handleRegister
	ds["/loginByToken"] = handleLoginByToken
	return ds
}

type TypeIsExisted struct {
	NickName  string `json:"nickName"`
	IsExisted bool   `json:"isExisted"`
}

// 检查昵称是否重复
func handleIsExistedNickName(w http.ResponseWriter, r *http.Request) {
	nickNames := mysql.QueryByField("nickName")
	nickName := r.URL.Query().Get("nickName")
	fmt.Printf("nickName: %v\n", nickName)
	tool.SendJson(w, TypeIsExisted{
		NickName:  nickName,
		IsExisted: nickNames[nickName],
	})
}

// 返回用户信息
func returnUserInfo(r *http.Request) tool.TypeUserInfo {
	var a = r.FormValue("avatarUrl")
	fmt.Println("a: ", a)
	if len(a) == 0 {
		a = "test_avatarUrl"
	}
	token, expireTime := tool.CreateToken()
	userInfo := tool.TypeUserInfo{
		AvatarUrl:  a,
		NickName:   r.FormValue("nickName"),
		Signature:  r.FormValue("signature"),
		Password:   r.FormValue("password"),
		Token:      token,
		ExpireTime: expireTime,
	}
	tool.ParseToken(userInfo.Token)
	return userInfo
}

// 注册
func handleRegister(w http.ResponseWriter, r *http.Request) {
	userInfo := returnUserInfo(r)
	mysql.RegisterUser(userInfo)
	tool.SetCookieToHeader(w, "token", userInfo.Token)
	tool.SendJson(w, userInfo)
}

// 登录
func handleLoginByToken(w http.ResponseWriter, r *http.Request) {
	token := tool.GetSingleCookieFromCookies(r.Header.Get("Cookie"), "token")
	ok, userInfo := mysql.QueryUserInfo("token", token)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	}
	tool.SendJson(w, userInfo)
}

// type R struct {
// 	D string
// }

// func handleTest(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(R{
// 		D: "test",
// 	})
// }

// var session_key string // 会话密钥

// func handleLogin(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	w.Header().Set("Content-Type", "application/json")
// 	code := query.Get("code")                    // 临时登录凭证
// 	appid := "wx20deb0404aa253b8"                // 小程序ID
// 	secret := "3c0d4b209e1b16edf8d105f1f4ab8976" // 小程序密钥
// 	grant_type := "authorization_code"           // 授权类型

// 	client := &http.Client{Timeout: 5 * time.Second}
// 	url := fmt.Sprintf(
// 		"https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=%v",
// 		appid, secret, code, grant_type,
// 	)
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	type Response struct {
// 		Session_key string `json:"session_key"` // 会话密钥
// 		Openid      string `json:"openid"`      // 用户唯一标识
// 	}
// 	var response = Response{}
// 	if err = json.Unmarshal(body, &response); err != nil {
// 		panic(err)
// 	}
// 	session_key = response.Session_key
// 	type Res struct {
// 		Openid string
// 	}
// 	json.NewEncoder(w).Encode(Res{
// 		Openid: response.Openid,
// 	})
// }
