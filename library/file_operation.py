# -- coding: utf-8 --
import os


class FileOperation(object):
    @staticmethod
    def make_dir(file_path):
        if not os.path.exists(file_path):
            os.makedirs(file_path)

    @staticmethod
    def touth_file(file_path, filename):
        FileOperation.make_dir(file_path)
        if not os.path.exists(file_path + "/" + filename):
            os.mknod(file_path + "/" + filename)

    @staticmethod
    def write_file(file_path, filename, file_data):
        FileOperation.touth_file(file_path, filename)
        if type(file_data) is bytes:
            data = file_data.decode()
        else:
            data = file_data
        with open(file_path + "/" + filename, "w") as f:
            f.write(data)

if __name__ == '__main__':
    a = FileOperation()
    a.make_dir('./a/b/c')
