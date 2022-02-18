# -- coding: utf-8 --
import os
import rarfile
import zipfile
import py7zr


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
            data = file_data
        else:
            data = file_data.encode()
        with open(file_path + "/" + filename, "wb") as f:
            f.write(data)

    @staticmethod
    def write(filename, file_data):
        if type(file_data) is bytes:
            data = file_data
        else:
            data = file_data.encode()
        with open(filename, "wb") as f:
            f.write(data)

    def un_zip(self, file_name, un_zip_path):
        zip_file = zipfile.ZipFile(file_name)
        zip_file.extractall(un_zip_path)
        zip_file.close()

    def make_zip (self, source_dir, output_filename):
        zipf = zipfile.ZipFile(output_filename, 'w',zipfile.ZIP_DEFLATED)
        pre_len = len(os.path.dirname(source_dir))
        for parent, dirnames, filenames in os.walk(source_dir):
            for filename in filenames:
                pathfile = os.path.join(parent, filename)
                arcname = pathfile[pre_len:].strip(os.path.sep)  # 相对路径
                zipf.write(pathfile, arcname)
        zipf.close()

    def un_rar(self, file_name, un_rar_path):
        rar = rarfile.RarFile(file_name)
        rar.extractall(un_rar_path)
        rar.close()

    def un_7z(self, file_name, path):
        with py7zr.SevenZipFile(file_name, mode='r') as z:
            z.extractall(path)

    def un_file(self, file_name, file_path):
        try:
            if file_name.endswith(".rar"):
                self.un_rar(file_name, file_path)
            elif file_name.endswith(".zip"):
                self.un_zip(file_name, file_path)
            elif file_name.endswith(".7z"):
                self.un_7z(file_name, file_path)
            else:
                raise "暂不支持此文件后缀"
            list_dir = os.listdir(file_path)
            data = []
            for i in list_dir:
                if os.path.isdir(file_path + "/" + i):
                    data_dict = {
                                "type": "package",
                                "package_name": i,
                                "file_path": file_path + "/" + i + "/" + "package.mo",
                                }
                    data.append(data_dict)
                else:
                    pass
            return data, ""
        except Exception as e:
            print(e)
            return False, "解压失败"
if __name__ == '__main__':
    a = FileOperation()
