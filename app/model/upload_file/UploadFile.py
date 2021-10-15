from sqlalchemy import Column, Integer, DateTime, VARCHAR, TEXT
from sqlalchemy.orm import declarative_base
import datetime
Base = declarative_base()


class UploadFile(Base):
    __tablename__ = 'upload_file_record'

    id = Column(Integer, primary_key=True)
    username = Column(VARCHAR)
    create_time = Column(DateTime)
    file_path = Column(TEXT)
    file_name = Column(VARCHAR)
