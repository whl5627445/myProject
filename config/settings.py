# -- coding: utf-8 --
import os
import logging

logging.basicConfig(level=logging.INFO,  # 控制台打印的日志级别
                    filename='/home/simtek/code/Log/OM.log',
                    filemode='a',  ##模式，有w和a，w就是写模式，每次都会重新写日志，覆盖之前的日志
                    # a是追加模式，默认如果不写的话，就是追加模式
                    format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s'
                    # 日志格式
                    )

modelica_keywords = ["der", "and", "or", "not", "constant"]

USERNAME = os.getenv("USERNAME", "")

JMODELICA_CONNECT = ("119.3.155.11", 56789)

EXAMPLES = [
    {
        "id": 1,
        "name": "Modelica.Blocks.Examples.TotalHarmonicDistortion",
        "image": "https://yssim-static.obs.cn-east-3.myhuaweicloud.com/example/thumbnail/TotalHarmonicDistortion.jpg",
    },
    {
        "id": 1,
        "name": "Modelica.Blocks.Examples.Filter",
        "image": "https://yssim-static.obs.cn-east-3.myhuaweicloud.com/example/thumbnail/Filter.jpg",
    },
    {
        "id": 1,
        "name": "Modelica.Blocks.Examples.NoiseExamples.Densities",
        "image": "https://yssim-static.obs.cn-east-3.myhuaweicloud.com/example/thumbnail/Densities.jpg",
    },
]

MQ_CONNECT = ["119.3.155.11:9092"]


IMAGE_FORMAT = ["bmp","jpg","jpeg","png","tif","gif","pcx","tga","exif","fpx","svg","cdr","pcd","dxf","ufo","eps","ai","raw","WMF","webp","avif","apng"]
