# -- coding: utf-8 --
from sqlalchemy import Column, Integer, VARCHAR, JSON, TEXT, BOOLEAN, String, DateTime
from sqlalchemy.orm import declarative_base
import datetime
Base = declarative_base()


class FmuAttachment(Base):
    __tablename__ = 'fmu_attachment'

    id = Column(Integer, primary_key=True)
    file_name = Column(VARCHAR)
    obs_name = Column(VARCHAR)
    create_user = Column(VARCHAR)
