# -- coding: utf-8 --
from config.settings import MQ_CONNECT
from kafka import KafkaProducer
import logging


producer = KafkaProducer(bootstrap_servers=MQ_CONNECT)
