#环境 dev,prd
env=dev
debug=true
http_access_host=0.0.0.0
http_listen_port=9000

#数据库
db_driver=mysql
db_host=zy_mysql
db_port=3306
db_user=root
db_pass=ziyue
db_name=ziyue
db_debug=true

#es配置
es_host=zy_elasticsearch:9200

#mongo_host配置
mongo_host=mongodb://zy_mongo:27017
mongo_conn_timeout=10

#redis
redis_host=zy_redis:6379
redis_pwd=
redis_db=0
redis_cache_version=20200101

#jwt
jwt_secret=123456
jwt_expr_time=8600
log_dir=./log.log
log_level=info

#谷歌验证码秘钥
ga_secret=xxxxxxxxxxxx