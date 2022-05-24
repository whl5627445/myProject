# -- coding: utf-8 --
from config.settings import MQ_CONNECT
from kafka import KafkaProducer

producer = KafkaProducer(bootstrap_servers=MQ_CONNECT)
