# -- coding: utf-8 --
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, close_all_sessions

# 本地
engine = create_engine('mysql+pymysql://root:root@localhost:3306/test?charset=utf8mb4',
                       pool_size=200,
                       max_overflow=50,
                       pool_recycle=3600,
                       # pool_timeout=15,
                       pool_pre_ping=True,
                       echo_pool=True,
                       )
# 测试环境
# engine = create_engine('mysql+pymysql://root:simtek_cloud_sim@121.36.222.22:3307/simtek-cloud?charset=utf8mb4',
#                        pool_size=200,
#                        max_overflow=500,
#                        pool_recycle=3600,
#                        pool_pre_ping=True,
#                        )
DBSession = sessionmaker(bind=engine, autoflush=False, autocommit=True)
session = DBSession()
