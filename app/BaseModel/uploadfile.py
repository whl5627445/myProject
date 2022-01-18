# -- coding: utf-8 --
from typing import Optional
from pydantic import BaseModel


class UploadSaveFileModel(BaseModel):
    package_name: Optional[str]
    model_str: Optional[str] = None
    package_id: Optional[str] = None


class UploadSaveModelModel(BaseModel):
    package_name: Optional[str]
    str_type: Optional[str] = None
    model_str: Optional[str] = None
    package_id: Optional[str] = None
    vars: dict = {
        "expand": "",
        "insert_to": "",
        "partial": False,
        "encapsulated": False,
        "state": False
        }
