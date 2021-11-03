# -- coding: utf-8 --
from sqlalchemy import Column, Integer, String, DateTime, VARCHAR, JSON, TEXT
from sqlalchemy.orm import declarative_base
import datetime

Base = declarative_base()


class ExperimentRecord(Base):
    __tablename__ = 'experiment_record'

    id = Column(Integer, primary_key=True)
    username = Column(VARCHAR)
    package_id = Column(Integer)
    create_time = Column(DateTime, default=datetime.datetime.now)
    model_name_all = Column(VARCHAR)
    model_var_data = Column(JSON)
    simulate_var_data = Column(JSON)
    experiment_name = Column(VARCHAR)
    package_name = Column(VARCHAR)
