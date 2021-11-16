# -- coding: utf-8 --
from typing import Optional, List
from pydantic import BaseModel


class ResponseModel(BaseModel):
    data: List = []
    err: Optional[str] = ""  # 正常是0，omc错误是1， 用戶错误是2
    msg: Optional[str] = ""
    status: int = 0

    class Config:
        orm_mode = True


class InitResponseModel(object):
    data = []
    err = ""
    msg = ""
    status: int = 0




