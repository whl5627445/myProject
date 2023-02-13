FROM python:3.9
COPY requirements.txt /yssim-go/grpc/PythonGrpc/
WORKDIR /yssim-go/grpc/PythonGrpc
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
	pip install -r requirements.txt
EXPOSE 50051
CMD python app.py