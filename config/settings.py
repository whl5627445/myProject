# -- coding: utf-8 --
import os


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
