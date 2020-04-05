package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	// MySQL用ドライバ
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// authクラス
)

func TestMain(m *testing.M) {
	fmt.Println("before test serve_test.go")
	// 環境変数ファイルの読込
	err := godotenv.Load(fmt.Sprintf("config/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	code := m.Run()
	fmt.Println("after test serve_test.go")
	os.Exit(code)
}

func TestActuaterHealth(t *testing.T) {
	t.Log("START TestActuaterHealth")

	gin.SetMode(gin.TestMode)
	router := serve()
	// router := gin.Default()
	router.Use(gin.Logger())

	// router.GET("/api/actuaterHealth", mockHandler)
	req, _ := http.NewRequest("GET", "/api/actuaterHealth", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestActuaterHealth")
}

func TestFetchAllMembers(t *testing.T) {
	t.Log("START TestFetchAllMembers")

	gin.SetMode(gin.TestMode)
	router := serve()
	// router := gin.Default()
	router.Use(gin.Logger())

	// router.GET("/api/actuaterHealth", mockHandler)
	req, _ := http.NewRequest("GET", "/api/fetchAllMembers", nil)
	rec := httptest.NewRecorder()
	// JWTのセット
	str := os.Getenv("TEST_JWT")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", str))

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestFetchAllMembers")
}

func TestFetchAllWorker(t *testing.T) {
	t.Log("START TestFetchAllWorker")

	gin.SetMode(gin.TestMode)
	router := serve()
	// router := gin.Default()
	router.Use(gin.Logger())

	// router.GET("/api/actuaterHealth", mockHandler)
	req, _ := http.NewRequest("GET", "/api/fetchAllWorker", nil)
	rec := httptest.NewRecorder()

	// JWTのセット
	str := os.Getenv("TEST_JWT")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", str))

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestFetchAllWorker")
}

func TestFetchProfileInfo(t *testing.T) {
	t.Log("START TestFetchProfileInfo")

	gin.SetMode(gin.TestMode)
	router := serve()
	// router := gin.Default()
	router.Use(gin.Logger())

	// router.GET("/api/actuaterHealth", mockHandler)
	req, _ := http.NewRequest("GET", "/api/fetchProfileInfo", nil)
	rec := httptest.NewRecorder()

	// JWTのセット
	str := os.Getenv("TEST_JWT")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", str))

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestFetchProfileInfo")
}

func TestFetchSignUpAccountMail(t *testing.T) {
	t.Log("START TestFetchSignUpAccountMail")

	gin.SetMode(gin.TestMode)
	router := serve()
	// router := gin.Default()
	router.Use(gin.Logger())

	// router.GET("/api/actuaterHealth", mockHandler)
	req, _ := http.NewRequest("GET", "/api/fetchSignUpAccountMail", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestFetchSignUpAccountMail")
}

func TestFetchMailAdrInfo(t *testing.T) {
	t.Log("START TestFetchMailAdrInfo")

	gin.SetMode(gin.TestMode)
	router := serve()
	router.Use(gin.Logger())

	// パラメータを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	values.Set("id", "1")  // key-valueを追加

	fmt.Println(strings.NewReader(values.Encode()))
	req, _ := http.NewRequest("POST", "/api/fetchMailAdrInfo", strings.NewReader(values.Encode()))
	rec := httptest.NewRecorder()

	// JWTのセット
	str := os.Getenv("TEST_JWT")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", str))
	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestFetchMailAdrInfo")
}

func TestFetchLoginInfo(t *testing.T) {
	t.Log("START TestFetchLoginInfo")

	gin.SetMode(gin.TestMode)
	router := serve()
	router.Use(gin.Logger())

	// パラメータを組み立て
	values := url.Values{}          // url.Valuesオブジェクト生成
	values.Set("username", "test1") // key-valueを追加
	values.Add("password", "test1") // key-valueを追加

	fmt.Println(strings.NewReader(values.Encode()))
	req, _ := http.NewRequest("POST", "/api/fetchLoginInfo", strings.NewReader(values.Encode()))
	rec := httptest.NewRecorder()

	// JWTのセット
	str := os.Getenv("TEST_JWT")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", str))
	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestFetchLoginInfo")
}

func mockHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "mock call successful",
	})
}
