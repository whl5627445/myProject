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
    model_variable_data = Column(JSON)
    model_variable_data_abscissa = Column(JSON)
    variable_description = Column(VARCHAR)
