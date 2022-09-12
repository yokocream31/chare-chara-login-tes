#Dockerfile作成
FROM couchbase/server:community-aarch64

# Couchbase初期化シェルをコンテナ内にコピー
COPY ./configure.sh /opt/couchbase
# COPY ./byfield/ /opt/couchbase

# Couchbase設定実行シェル
CMD [ "/opt/couchbase/configure.sh" ]