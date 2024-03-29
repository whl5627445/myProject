import time

from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import (
    create_engine,
    Column,
    String,
    Integer,
    Text,
    DateTime,
    JSON,
    Boolean,
    Float,
)
from sqlalchemy.orm import sessionmaker
from sqlalchemy.pool import QueuePool
from libs.function.grpc_log import log

HOST = "mysql"  # 127.0.0.1/localhost 124.70.211.127
PORT = 3306
#
# HOST = '124.70.211.127'  # 127.0.0.1/localhost 124.70.211.127
# PORT = 3307
DATA_BASE = "yssim"
USER = "root"
PWD = "root"

DB_URI = f"mysql+pymysql://{USER}:{PWD}@{HOST}:{PORT}/{DATA_BASE}"

while True:
    try:
        engine = create_engine(
            DB_URI,
            poolclass=QueuePool,
            pool_size=50,  # 最大连接数  124 3个  119 50个
            max_overflow=30,  # 连接池溢出后允许的最大连接数
            pool_timeout=15,  # 请求超时时间（秒）
            pool_pre_ping=True,  # 每次从连接池中取连接的时候，都会验证一下与数据库是否连接正常
            pool_recycle=3600,  # 主动回收mysql连接的时间
        )
        log.info("连接数据库成功！")
        break
    except Exception as e:
        log.info("连接数据库失败：" + str(e))
        time.sleep(3)

Base = declarative_base(engine)
Session = sessionmaker(engine)


class YssimSimulateRecords(Base):
    __tablename__ = "yssim_simulate_records"
    id = Column(String, primary_key=True)
    experiment_id = Column(String)
    username = Column(String)
    simulate_model_name = Column(String)
    simulate_model_result_path = Column(String)
    simulate_status = Column(String)
    simulate_start_time = Column(Integer)
    simulate_end_time = Column(Integer)
    simulate_result_str = Column(Text)
    percentage = Column(Integer)
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
    result_run_time = Column(Float)
    another_name = Column(String)
    number_intervals = Column(String)
    env_model_data = Column(JSON)
    simulate_type = Column(String)
    deleted_at = Column(DateTime)
    create_time = Column(DateTime)


class YssimModels(Base):
    __tablename__ = "yssim_models"
    id = Column(String, primary_key=True)
    userspace_id = Column(String)
    package_name = Column(String)
    version = Column(String)
    sys_or_user = Column(String)
    file_path = Column(String)
    default_version = Column(Integer)
    deleted_at = Column(DateTime)


class AppDataSources(Base):
    __tablename__ = "app_data_sources"
    id = Column(String, primary_key=True)
    username = Column(String)
    user_space_id = Column(String)
    package_id = Column(String)
    model_name = Column(String)
    compile_type = Column(String)
    compile_path = Column(String)
    compile_status = Column(Integer)
    compile_start_time = Column(Integer)
    compile_stop_time = Column(Integer)
    result_run_time = Column(Float)
    group_name = Column(String)
    data_source_name = Column(String)
    experiment_id = Column(String)
    env_model_data = Column(JSON)
    start_time = Column(String)
    stop_time = Column(String)
    method = Column(String)
    number_intervals = Column(String)
    tolerance = Column(String)
    deleted_at = Column(DateTime)
    create_time = Column(DateTime)
    zip_mo_path = Column(String)


class AppPages(Base):
    __tablename__ = "app_pages"
    id = Column(String, primary_key=True)
    mul_result_path = Column(String)
    release_state = Column(Integer)
    naming_order = Column(JSON)
    app_space_id = Column(String)
    release_time = Column(Integer)
    release_message_read = Column(Boolean)
    mul_sim_state = Column(Integer)
    mul_sim_time = Column(Integer)
    mul_sim_message_read = Column(Boolean)
    mul_sim_err = Column(String)
    release_err = Column(String)
    is_release = Column(Boolean)
    is_mul_simulate = Column(Boolean)
    is_preview = Column(Boolean)


class AppSpaces(Base):
    __tablename__ = "app_spaces"
    id = Column(String, primary_key=True)
    is_release = Column(Boolean)


class AppPagesComponent(Base):
    __tablename__ = "app_page_components"
    id = Column(String, primary_key=True)
    page_id = Column(String)
    type = Column(String)
    width = Column(Integer)
    height = Column(Integer)
    position_x = Column(Integer)
    position_y = Column(Integer)
    angle = Column(Integer)
    horizontal_flip = Column(Boolean)
    vertical_flip = Column(Boolean)
    opacity = Column(Integer)
    other_configuration = Column(JSON)
    z_index = Column(Integer)
    styles = Column(JSON)
    events = Column(JSON)
    chart_config = Column(JSON)
    option = Column(JSON)
    component_path = Column(String)
    hide = Column(Boolean)
    lock = Column(Boolean)
    is_group = Column(Boolean)
    create_time = Column(DateTime)
    update_time = Column(DateTime)
    deleted_at = Column(DateTime)
    input_name = Column(String)
    output = Column(JSON)
    max = Column(Float)
    min = Column(Float)
    interval = Column(Float)


class AppPagesComponentRelease(Base):
    __tablename__ = "app_page_components_releases"
    id = Column(String, primary_key=True)
    page_id = Column(String)
    type = Column(String)
    width = Column(Integer)
    height = Column(Integer)
    position_x = Column(Integer)
    position_y = Column(Integer)
    angle = Column(Integer)
    horizontal_flip = Column(Boolean)
    vertical_flip = Column(Boolean)
    opacity = Column(Integer)
    other_configuration = Column(JSON)
    z_index = Column(Integer)
    styles = Column(JSON)
    events = Column(JSON)
    chart_config = Column(JSON)
    option = Column(JSON)
    component_path = Column(String)
    hide = Column(Boolean)
    lock = Column(Boolean)
    is_group = Column(Boolean)
    create_time = Column(DateTime)
    deleted_at = Column(DateTime)
    input_name = Column(String)
    output = Column(JSON)
    max = Column(Float)
    min = Column(Float)
    interval = Column(Float)


class AppPagesComponentPreview(Base):
    __tablename__ = "app_page_components_previews"
    id = Column(String, primary_key=True)
    page_id = Column(String)
    type = Column(String)
    width = Column(Integer)
    height = Column(Integer)
    position_x = Column(Integer)
    position_y = Column(Integer)
    angle = Column(Integer)
    horizontal_flip = Column(Boolean)
    vertical_flip = Column(Boolean)
    opacity = Column(Integer)
    other_configuration = Column(JSON)
    z_index = Column(Integer)
    styles = Column(JSON)
    events = Column(JSON)
    chart_config = Column(JSON)
    option = Column(JSON)
    component_path = Column(String)
    hide = Column(Boolean)
    lock = Column(Boolean)
    is_group = Column(Boolean)
    create_time = Column(DateTime)
    deleted_at = Column(DateTime)
    input_name = Column(String)
    output = Column(JSON)
    max = Column(Float)
    min = Column(Float)
    interval = Column(Float)


class ParameterCalibrationRecord(Base):
    __tablename__ = "parameter_calibration_records"
    id = Column(String, primary_key=True)
    package_id = Column(String)
    userspace_id = Column(String)
    version = Column(String)
    username = Column(String)
    model_name = Column(String)
    compile_path = Column(String)
    package_path = Column(String)
    compile_Dependencies = Column(JSON)
    compile_status = Column(String)
    compile_start_time = Column(Integer)
    compile_stop_time = Column(Integer)

    simulate_model_result_path = Column(String)
    simulate_status = Column(String)
    start_time = Column(String)
    stop_time = Column(String)
    interval = Column(String)
    simulate_result_str = Column(Text)
    percentage = Column(JSON)

    rated_condition = Column(JSON)
    condition_parameters = Column(JSON)
    formula = Column(JSON)
    actual_data = Column(JSON)
    simulate_result = Column(JSON)
    associated_parameters = Column(JSON)
    result_parameters = Column(JSON)
    deleted_at = Column(DateTime)
