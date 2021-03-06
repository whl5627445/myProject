# -- coding: utf-8 --
from kafka import KafkaConsumer
from config.settings import MQ_CONNECT, USERNAME
import  time, json
import logging as log
from config.DB_config import DBSession
from app.service.simulate_func import SimulateTask
from app.model.Simulate.SimulateRecord import SimulateRecord

session = DBSession()
log.basicConfig(level=log.INFO,  # 控制台打印的日志级别
                    format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s'
                    # 日志格式
                    )



class SimulateService(object):

    def __init__(self):
        self.username = USERNAME
        self.consumer = KafkaConsumer(bootstrap_servers=MQ_CONNECT, group_id='simulate', max_in_flight_requests_per_connection=9999,api_version=(0, 10))
        self.consumer.subscribe([self.username + "_" + "SIMULATE"])

    def start(self):

        for message in self.consumer:
            data = json.loads(message.value.decode('utf-8'))
            space_id = data["space_id"]
            SRecord_id = data["SRecord_id"]
            model_name = data["model_name"]
            s_type = data["s_type"]
            file_path = data["file_path"]
            simulate_parameters_data = data["simulate_parameters_data"]
            if message:
                log.info('message: {}'.format(message))
                result, result_str = self.run(space_id, SRecord_id, model_name, s_type, file_path,simulate_parameters_data)
                log.info('result_str: {}'.format(result_str))

    def run(self, space_id, SRecord_id, model_name, s_type, file_path = None, simulate_parameters_data=None):
        s_result, s_str = SimulateTask(space_id, SRecord_id, self.username, model_name, s_type, file_path,simulate_parameters_data)
        return s_result, s_str


if __name__ == '__main__':
    # SRecord = session.query(SimulateRecord).filter(SimulateRecord.username == USERNAME).all()
    log.info('Starting consumer')
    s = SimulateService()
    log.info('consumer is ok')
    s.start()

