package user

import (
	"crypto/md5"
	"github.com/bitrho/bit/model/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
)

var (
	TABLE          = "user"
	COOKIE_SECURET = "^faefa%FaeefaeaefaaKKKfawef"
	COOKIE_KEY     = "_bitrho"
)

type User struct {
	Uid     bson.ObjectId `bson:"_id"`
	Uname   string
	Pass    string
	Email   string
	Phone   string
	Sex     int
	Avatar  string
	AdminId int
	Salt    string
}

func New(uname string, pass string, email string, phone string, sex int, avatar string, adminId int) *User {
	uid := bson.NewObjectId()
	salt := createSalt()
	pwd := encryptPass(uid.String(), pass, salt, true)
	return &User{
		Uid:     uid,
		Uname:   uname,
		Pass:    pwd,
		Email:   email,
		Phone:   phone,
		Sex:     sex,
		Avatar:  avatar,
		AdminId: adminId,
		Salt:    salt,
	}
}

func Create(uname string, pass string, email string, phone string, sex int, avatar string, adminId int) (uid bson.ObjectId, err error) {
	u := New(uname, pass, email, phone, sex, avatar, adminId)
	uid, err = insert(u)
	return
}

func GetByUname(uname string) (u *User, err error) {
	return queryByUname(uname)
}

func GetByUid(uid string) (u *User, err error) {
	return queryByUid(uid)
}

func ceatePass(uid string, pass string, isMd5 bool) (pwd string) {
	salt := createSalt()
	pwd = encryptPass(uid, pass, salt, isMd5)
	return
}

func encryptPass(uid string, pass string, salt string, isMd5 bool) (pwd string) {
	var md5Pass string
	if isMd5 {
		md5Pass = pass
	} else {
		h := md5.New()
		io.WriteString(h, pass)
		md5Pass = string(h.Sum(nil))
	}
	ha := md5.New()
	io.WriteString(ha, salt+md5Pass+uid)
	pwd = string(ha.Sum(nil))
	return
}

func createSalt() (salt string) {
	salt = util.RandomString(10)
	log.Println(salt)
	return
}

func insert(u *User) (uid bson.ObjectId, err error) {
	uid = u.Uid
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("bit").C(TABLE)
	err = c.Insert(u)
	return
}

func queryByUname(uname string) (u *User, err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("bit").C(TABLE)
	u = &User{}
	err = c.Find(bson.M{"uname": uname}).One(u)
	return
}

func queryByUid(uid string) (u *User, err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("bit").C(TABLE)
	u = &User{}
	err = c.Find(bson.M{"_id": bson.ObjectId(uid)}).One(u)
	return
}
