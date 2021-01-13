/**
 * @Author: li.zhang
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2020/12/11 下午4:49
 */
package model

import (
	"crypto/rsa"
	"github.com/zlsmd/zchat/server/utils"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
	"time"
)

type User struct {
	Id       int64
	Account  string
	Password string
	NickName string
	AddTime  string
	IsDel    byte
	IsOnline byte
	GroupId  int32
}

type LoginJwt struct {
	*jwt.StandardClaims
	UserId  int64 `json:"userId"`
	GroupId int32 `json:"groupId"`
}

const PriKey = ``

const PubKey = ``

func (u *User) TableName() string {
	return TablePre + "user"
}

const PasswordKey = ""

func (u *User) CheckLogin(userName, password string) (ok bool, token string, user *User) {
	db.Select("id, group_id, password, nick_name, account").Where("account = ?", userName).Limit(1).Find(u)
	if u.Password == utils.Md5(password+PasswordKey) {
		ok = true

		var err error
		t := jwt.New(jwt.GetSigningMethod("RS256"))
		t.Claims = &LoginJwt{
			StandardClaims: &jwt.StandardClaims{
				Audience:  "",
				ExpiresAt: time.Now().Unix() + 3600,
				Id:        "",
				IssuedAt:  0,
				Issuer:    "",
				NotBefore: 0,
				Subject:   "",
			},
			UserId:  u.Id,
			GroupId: u.GroupId,
		}
		var key *rsa.PrivateKey
		key, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(PriKey))
		if err != nil {
			token = ""
			return
		}
		token, err = t.SignedString(key)
		user = u
	}
	return
}
