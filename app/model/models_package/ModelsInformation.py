# -- coding: utf-8 --
from sqlalchemy import Column, Integer, VARCHAR, JSON, TEXT, BOOLEAN, String, DateTime
from sqlalchemy.orm import declarative_base
import datetime
Base = declarative_base()


class ModelsInformation(Base):
    __tablename__ = 'models'

    id = Column(Integer, primary_key=True)
    package_name = Column(VARCHAR)
    create_time = Column(DateTime, default=datetime.datetime.now)
    model_name = Column(VARCHAR)
    child_name = Column(JSON)
    haschild = Column(BOOLEAN)
    sys_or_user = Column(String)
    image = Column(TEXT)
    file_path = Column(VARCHAR)


class ModelsInformationAll(Base):
    __tablename__ = 'models_all'

    id = Column(Integer, primary_key=True)
    package_id = Column(Integer)
    package_name = Column(VARCHAR)
    model_name = Column(VARCHAR)
    model_name_all = Column(VARCHAR)
    parent_name = Column(VARCHAR)
    sys_or_user = Column(VARCHAR)
    haschild = Column(BOOLEAN)
    child_name = Column(JSON)
    image = Column(TEXT)
