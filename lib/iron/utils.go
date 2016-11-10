package iron

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	r "math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func SafePage(page int) int {
	if page <= 0 {
		page = 1
	}
	return page
}

func CalculatePages(count, perPageCount int) int {
	ret := count / perPageCount
	if count%perPageCount > 0 {
		ret += 1
	}
	return ret
}

func IsEmail(email string) bool {
	regEmail := regexp.MustCompile("^[a-zA-Z0-9-_.]+@.*\\..*$")
	return regEmail.MatchString(email)
}

type Sizer interface {
	Size() int64
}

type ImageCouldCrop interface {
	image.Image
	SubImage(r image.Rectangle) image.Image
}

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func HtmlSpecialchars(res *string) {
	*res = strings.Replace(*res, ">", "&gt;", -1)
	*res = strings.Replace(*res, "<", "&lt;", -1)
	*res = strings.Replace(*res, "\r\n", "<br />", -1)
	*res = strings.Replace(*res, "\n", "<br />", -1)
}

func HtmlSpecialcharsSafeDecode(res *string) {
	*res = strings.Replace(*res, "<br />", "\n", -1)
}

func SqlEscape(res interface{}) string {
	var ret string
	switch res.(type) {
	case string:
		ret = strings.Replace(res.(string), "\\", "\\\\", -1)
		ret = strings.Replace(ret, "'", "\\'", -1)
		ret = strings.Replace(ret, "\"", "\\\"", -1)
		ret = strings.Replace(ret, "/", "\\/", -1)
	default:
		ret = fmt.Sprintf("%v", res)
	}
	return "'" + ret + "'"
}

func Int64IsIn(value int64, arr []int64) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}

func StringIsIn(value string, arr []string) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}

func initEncoder() {
	gob.Register(map[string]interface{}{})
}

func Encode(data interface{}) ([]byte, error) {
	var err error
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)

	switch data.(type) {
	case map[string]interface{}:
		nd := map[string]interface{}(data.(map[string]interface{}))
		for _, v := range nd {
			gob.Register(v)
		}
		err = enc.Encode(nd)
	case map[interface{}]interface{}:
		nd := map[string]interface{}(data.(map[string]interface{}))
		for _, v := range nd {
			gob.Register(v)
		}
		err = enc.Encode(nd)
	case []interface{}:
		nd := []interface{}(data.([]interface{}))
		for _, v := range nd {
			gob.Register(v)
		}
		err = enc.Encode(nd)
	default:
	}
	err = enc.Encode(data)

	if err != nil {
		return []byte(""), err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(to)
	return err
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func FileExt(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index < 0 {
		return ""
	}
	return filename[index+1:]
}

func InitId() string {
	now := time.Now()
	r := r.New(r.NewSource(now.UnixNano()))
	return fmt.Sprintf("%d%d", time.Now().UnixNano(), r.Intn(100000))
}

func imageDecode(fileext string, file io.Reader) (image.Image, error) {
	switch fileext {
	case "jpg", "jpeg":
		return jpeg.Decode(file)
	case "png":
		return png.Decode(file)
	case "gif":
		return gif.Decode(file)
	default:
		return nil, nil
	}
}

func imageEncode(fileext string, file io.Writer, img image.Image) {
	switch fileext {
	case "jpg", "jpeg":
		jpeg.Encode(file, img, nil)
	case "png":
		png.Encode(file, img)
	case "gif":
		gif.Encode(file, img, nil)
	}
}

func ApiRet(data interface{}, errno int) []byte {
	ret := map[string]interface{}{"data": data, "errno": errno}
	res, _ := json.Marshal(ret)
	return res
}

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[r.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[r.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return bytes
}
