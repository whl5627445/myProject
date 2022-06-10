# -- coding: utf-8 --
from sqlalchemy import Column, Integer, VARCHAR, JSON, TEXT, BOOLEAN, String, DateTime
from sqlalchemy.orm import declarative_base
import datetime
Base = declarative_base()


class UserSpace(Base):
    __tablename__ = 'py_user_space'

    id = Column(Integer, primary_key=True)
    spacename = Column(VARCHAR)
    username = Column(VARCHAR)
    update_time = Column(DateTime)
    create_time = Column(DateTime, default=datetime.datetime.now)
    last_login_time = Column(DateTime)
