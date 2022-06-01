# -- coding: utf-8 --
from sqlalchemy import Column, Integer, VARCHAR, JSON, TEXT, BOOLEAN, String, DateTime
from sqlalchemy.orm import declarative_base
import datetime
Base = declarative_base()


class ModelsInformation(Base):
    __tablename__ = 'py_models'

    id = Column(Integer, primary_key=True)
    package_name = Column(VARCHAR)
    create_time = Column(DateTime, default=datetime.datetime.now)
    update_time = Column(DateTime)
    model_name = Column(VARCHAR)
    sys_or_user = Column(String)
    file_path = Column(VARCHAR)
    userspace_id = Column(Integer, default=0)



