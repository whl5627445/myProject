# -- coding: utf-8 --
from typing import Optional
from pydantic import BaseModel


class UploadSaveFileModel(BaseModel):
    package_name: Optional[str]
    model_str: Optional[str]
    package_id: Optional[str]

    # class Config:
    #     orm_mode = True

class UploadSaveModelModel(BaseModel):
    package_name: Optional[str]
    str_type: Optional[str] = None
    model_str: Optional[str] = None
    package_id: Optional[str] = None
    comment: Optional[str] = None
    vars: dict = {
        "expand": "",
        "insert_to": "",
        "partial": False,
        "encapsulated": False,
        "state": False
        }


class UploadIconModel(BaseModel):
    package_name: Optional[str]
    model_name: Optional[str]
    package_id: Optional[str]

