# -- coding: utf-8 --

from config.HW_OBS import obsClient
# 使用访问OBS
bucketName = "yssim-static"

class OBSClient(object):
    def __init__(self):
        self.bucketName = "yssim-static"

    def putFile(self, new_filename, local_file):
        res = obsClient.putFile(self.bucketName, new_filename, local_file)
        return res
