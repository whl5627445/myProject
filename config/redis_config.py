# -- coding: utf-8 --
import redis
# 本地开发环境redis
pool = redis.ConnectionPool(host='localhost', port=6379)
# 测试环境redis
# pool = redis.ConnectionPool(host='127.0.0.1', port=6379, password="Openmodelica2022")

r = redis.Redis(connection_pool=pool)
# Openmodelica2022
