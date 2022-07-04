# -- coding: utf-8 --
from config.settings import MQ_CONNECT
from kafka import KafkaProducer

producer = KafkaProducer(bootstrap_servers=MQ_CONNECT,
                         retries=10000,
                         max_in_flight_requests_per_connection=1,
                         api_version=(0, 10)
                         )
