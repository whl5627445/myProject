from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine, Column, String, Integer, Text, DateTime, JSON
from sqlalchemy.orm import sessionmaker
from sqlalchemy.pool import QueuePool

HOST = 'mysql'  # 127.0.0.1/localhost 124.70.211.127
PORT = 3306
#
# HOST = '124.70.211.127'  # 127.0.0.1/localhost 124.70.211.127
# PORT = 3307
DATA_BASE = 'yssim'
USER = 'root'
PWD = 'simtek_cloud_sim'

DB_URI = f'mysql+pymysql://{USER}:{PWD}@{HOST}:{PORT}/{DATA_BASE}'

engine = create_engine(DB_URI,
                       poolclass=QueuePool,
                       pool_size=50,  # 最大连接数  124 3个  119 50个
                       max_overflow=10,  # 连接池溢出后允许的最大连接数
                       pool_timeout=15  # 请求超时时间（秒）
                       )
Base = declarative_base(engine)
Session = sessionmaker(engine)


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
    env_model_data = Column(JSON)
    simulate_type = Column(String)
    deleted_at = Column(DateTime)
    create_time = Column(DateTime)


class YssimModels(Base):
    __tablename__ = 'yssim_models'
    id = Column(String, primary_key=True)
    userspace_id = Column(String)
    package_name = Column(String)
    version = Column(String)
    sys_or_user = Column(String)
    file_path = Column(String)
    default_version = Column(Integer)
    deleted_at = Column(DateTime)
