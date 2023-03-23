from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine, Column, String, Integer, Text, DateTime
from sqlalchemy.orm import sessionmaker

HOST = 'mysql'  # 127.0.0.1/localhost 124.70.211.127
PORT = 3306
DATA_BASE = 'yssim'
USER = 'root'
PWD = 'simtek_cloud_sim'

DB_URI = f'mysql+pymysql://{USER}:{PWD}@{HOST}:{PORT}/{DATA_BASE}'

# with engine.connect() as conn:
#     # 执行SQL即可
#     conn.execute(sql)

engine = create_engine(DB_URI)
Base = declarative_base(engine)
Session = sessionmaker(engine)


# class ProcessState(Base):
#     __tablename__ = 'process_state'
#     uuid = Column(String, primary_key=True)
#     progress = Column(Integer)
#     exception = Column(Integer)
#     log = Column(Text)
#     state = Column(String)
#     processStartTime = Column(DateTime)
#     processRunTime = Column(DateTime)
#     resPath = Column(String)
#     def toDict(self):
#         return {
#             "uuid": self.uuid,
#             "progress": self.progress,
#             "exception": self.exception,
#             "log": self.log,
#             "state": self.state,
#             "processStartTime": self.processStartTime,
#             "processRunTime": self.processRunTime,
#             "resPath": self.resPath
#         }


class YssimSimulateRecords(Base):
    __tablename__ = 'yssim_simulate_records'
    id = Column(String, primary_key=True)
    experiment_id = Column(String)
    username = Column(String)
    simulate_model_name = Column(String)
    simulate_model_result_path = Column(String)
    simulate_status = Column(String)
    simulate_start_time = Column(Integer)
    simulate_end_time = Column(Integer)
    simulate_result_str = Column(Text)
    fmi_version = Column(String)
    description = Column(String)
    start_time = Column(String)
    stop_time = Column(String)
    step_size = Column(String)
    tolerance = Column(String)
    solver = Column(String)
    output_format = Column(String)
    variable_filter = Column(String)
    package_id = Column(String)
    userspace_id = Column(String)
    simulate_start = Column(String)
    method = Column(String)
    another_name = Column(String)
    number_intervals = Column(String)
    simulate_type = Column(String)
    deleted_at = Column(DateTime)
    create_time = Column(DateTime)
