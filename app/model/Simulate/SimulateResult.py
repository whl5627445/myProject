# -- coding: utf-8 --
from sqlalchemy import Column, Integer, String, DateTime, VARCHAR, JSON
from sqlalchemy.orm import declarative_base
import datetime
Base = declarative_base()


class SimulateResult(Base):
    __tablename__ = 'simulate_result'

    id = Column(Integer, primary_key=True)
    username = Column(VARCHAR)
    simulate_model_name = Column(VARCHAR)
    simulate_record_id = Column(VARCHAR)
    model_variable_name = Column(VARCHAR)
    model_variable_parent = Column(VARCHAR)
    model_variable_data = Column(JSON)
    model_variable_data_abscissa = Column(JSON)
    variable_description = Column(VARCHAR)
    value_reference = Column(VARCHAR)
    description = Column(VARCHAR)
    variability = Column(VARCHAR)
    is_discrete = Column(VARCHAR)
    causality = Column(VARCHAR)
    is_value_changeable = Column(VARCHAR)
    alias = Column(VARCHAR)
    class_index = Column(VARCHAR)
    class_type = Column(VARCHAR)
    is_protected = Column(VARCHAR)
    hide_result = Column(VARCHAR)
    file_name = Column(VARCHAR)
    start_line = Column(VARCHAR)
    start_column = Column(VARCHAR)
    end_line = Column(VARCHAR)
    end_column = Column(VARCHAR)
    file_writable = Column(VARCHAR)
    var_type = Column(VARCHAR)
    start = Column(VARCHAR)
    fixed = Column(VARCHAR)
    use_nominal = Column(VARCHAR)
    unit = Column(VARCHAR)
