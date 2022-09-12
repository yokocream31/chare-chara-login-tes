package db

import (
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
)

var Bucket *gocb.Bucket

// Databaseの初期化
func InitDB() {

	// Update this to your cluster details
	bucketName := os.Getenv("COUCHBASE_BUCKET")
	username := os.Getenv("COUCHBASE_ADMINISTRATOR_USERNAME")
	password := os.Getenv("COUCHBASE_ADMINISTRATOR_PASSWORD")
	host := os.Getenv("COUCHBASE_HOST")
	scheme := os.Getenv("COUCHBSE_SCHEME")

	// エラーと警告をログとして記録
	gocb.SetLogger(gocb.DefaultStdioLogger())

	// Initialize the Connection
	// クラスターへの接続
	cluster, err := gocb.Connect(scheme+host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
		SecurityConfig: gocb.SecurityConfig{
			TLSSkipVerify: true,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected: ", *cluster)

	// バケットへの参照の取得
	Bucket = cluster.Bucket(bucketName)
	// バケットに接続され確実に利用可能になるまでwait
	err = Bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

}