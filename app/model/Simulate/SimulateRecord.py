# -- coding: utf-8 --
from sqlalchemy import Column, Integer, String, DateTime, VARCHAR, JSON, TEXT
from sqlalchemy.orm import declarative_base
import datetime

Base = declarative_base()


class SimulateRecord(Base):
    __tablename__ = 'py_simulate_record'

    id = Column(Integer, primary_key=True)
    username = Column(VARCHAR)
    simulate_model_name = Column(VARCHAR)
    simulate_model_result_path = Column(VARCHAR)
    simulate_status = Column(VARCHAR)
    simulate_nametree = Column(JSON)
    simulate_start_time = Column(DateTime, default=datetime.datetime.now)
    simulate_end_time = Column(DateTime)
    simulate_result_str = Column(VARCHAR)
    fmi_version = Column(VARCHAR)
    description = Column(VARCHAR)
    start_time = Column(VARCHAR)
    stop_time = Column(VARCHAR)
    step_size = Column(VARCHAR)
    tolerance = Column(VARCHAR)
    solver = Column(VARCHAR)
    output_format = Column(VARCHAR)
    variable_filter = Column(VARCHAR)
