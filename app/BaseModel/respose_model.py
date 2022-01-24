# -- coding: utf-8 --
from typing import Optional, List
from pydantic import BaseModel


class ResponseModel(BaseModel):
    data: List = []
    err: Optional[str] = ""
    msg: Optional[str] = ""
    status: int = 0         # 正常是0，系统级错误是1， 用戶错误是2

    class Config:
        orm_mode = True


class InitResponseModel(object):
    def __init__(self):
        self.data = []
        self.err = ""
        self.msg = ""
        self.status: int = 0






