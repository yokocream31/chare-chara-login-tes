#!/bin/bash
# Couchbase設定シェル

# docker-compose up -d 前に chmod +x configure.shを実行する

set -m

# サーバ起動
/entrypoint.sh couchbase-server &

# APIサーバ起動までの待ち時間をざっくり設定 TODO:サーバ起動を検知して処理を進めるようにしたい。
sleep 30

# クラスタのメモリ割り当て
curl -v -X POST http://127.0.0.1:8091/pools/default -d memoryQuota=4096 -d indexMemoryQuota=4096

# サービス設定
curl -v http://127.0.0.1:8091/node/controller/setupServices -d services=kv%2cn1ql%2Cindex

# Couchbaseサーバーの認証情報設定
curl -v http://127.0.0.1:8091/settings/web -d port=8091 -d username=$COUCHBASE_ADMINISTRATOR_USERNAME -d password=$COUCHBASE_ADMINISTRATOR_PASSWORD

# バケット上限解除(バケットの上限数はデフォルトで10個,,それ以上作成する場合は以下設定で変更が必要)
curl -v -u $COUCHBASE_ADMINISTRATOR_USERNAME:$COUCHBASE_ADMINISTRATOR_PASSWORD -X POST http://127.0.0.1:8091/internalSettings -d maxBucketCount=20

# バケット作成
curl -v -u $COUCHBASE_ADMINISTRATOR_USERNAME:$COUCHBASE_ADMINISTRATOR_PASSWORD -X POST http://127.0.0.1:8091/pools/default/buckets -d name=$COUCHBASE_BUCKET -d bucketType=couchbase -d ramQuotaMB=100 -d authType=sasl -d saslPassword=ladd_personal_info1

echo "please wait 30 sec..."
sleep 30
# View設定
# curl -v -u $COUCHBASE_ADMINISTRATOR_USERNAME:$COUCHBASE_ADMINISTRATOR_PASSWORD -X PUT -H "Content-Type: application/json" http://127.0.0.1:8092/$COUCHBASE_BUCKET/_design/api%2Ftest -d @/opt/couchbase/all.ddoc


# sleep 15

# フォアグランド実行
fg 1