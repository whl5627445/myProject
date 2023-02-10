from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine, Column, String, Integer, Text, DateTime
from sqlalchemy.orm import sessionmaker

HOST = '124.70.211.127'  # 127.0.0.1/localhost  '124.70.211.127':3307   'mysql':3306
PORT = 3307   # 3307
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


class ProcessState(Base):
    __tablename__ = 'process_state'
    uuid = Column(String, primary_key=True)
    progress = Column(Integer)
    exception = Column(Integer)
    log = Column(Text)
    state = Column(String)
    processStartTime = Column(DateTime)
    processRunTime = Column(DateTime)
    resPath = Column(String)

    def toDict(self):
        return {
            "uuid": self.uuid,
            "progress": self.progress,
            "exception": self.exception,
            "log": self.log,
            "state": self.state,
            "processStartTime": self.processStartTime,
            "processRunTime": self.processRunTime,
            "resPath": self.resPath
        }
