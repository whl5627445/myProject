# -- coding: utf-8 --
import redis
from config.settings import REDIS_CONNECT

pool = redis.ConnectionPool(host='yssim-redis', port=6379)
# pool = redis.ConnectionPool(**REDIS_CONNECT)

r = redis.Redis(connection_pool=pool, charset='UTF-8')
# Openmodelica2022

