package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/carmel/gooxml/document"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"server/model"
	"strings"
	"time"
)

func MkMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func GetAppDir() string {
	appDir, err := os.Getwd()
	if err != nil {
		file, _ := exec.LookPath(os.Args[0])
		applicationPath, _ := filepath.Abs(file)
		appDir, _ = filepath.Split(applicationPath)
	}
	return appDir
}

func GetUser(c *gin.Context) (user model.User, err error) {
	session := sessions.Default(c)
	jsonStr := session.Get("user").(string)
	err = json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func MkEmailCode() string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}
	return code
}

func ReadTxt(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	defer file.Close()

	content := ""
	title := ""
	i := 0
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if i == 0 && isTitle(str) {
			title = strings.TrimSpace(str)
		} else {
			content += str
		}
		i++
	}
	return title, content
}

func ReadDoc(path string) (string, string) {
	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	content := ""
	title := ""
	for i, para := range doc.Paragraphs() {
		s := ""
		for _, run := range para.Runs() {
			text := run.Text()
			s += text
		}
		if i == 0 && isTitle(s) {
			title = strings.TrimSpace(s)
		} else {
			content += (s + "\n")
		}
	}
	return title, content
}

func isTitle(str string) bool {
	str = strings.TrimSpace(str)
	length := strings.Count(str, "") - 1
	if length < 24 && length > 0 {
		return true
	}
	return false
}

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
