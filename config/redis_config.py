# -- coding: utf-8 --
import redis
# 本地开发环境redis
# pool = redis.ConnectionPool(host='localhost', port=6379)
# 测试环境redis
pool = redis.ConnectionPool(host='119.3.155.11', port=56799)

r = redis.Redis(connection_pool=pool, charset='UTF-8')
# Openmodelica2022

