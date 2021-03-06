# -- coding: utf-8 --

from config.HW_OBS import obsClient
# 使用访问OBS


class HWOBS(object):
    def __init__(self):
        self.bucketName = "yssim-static"

    def putFile(self, new_filename, local_file):
        res = obsClient.putFile(self.bucketName, new_filename, local_file)
        return res

    @staticmethod
    def createSignedUrl(bucketName, objectKey, method="GET"):
        res = obsClient.createSignedUrl(method, bucketName, objectKey)
        return res


if __name__ == '__main__':
    obs = HWOBS()
    res = obs.putFile("example/thumbnail/Filter.jpg", "/home/simtek/code/public/UserFiles/Filter.jpg")
    print(res)
