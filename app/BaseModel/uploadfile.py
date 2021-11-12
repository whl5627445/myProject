# -- coding: utf-8 --
from typing import Optional
from pydantic import BaseModel


class UploadSaveFile(BaseModel):
    package_name: str
    str_type: str = None
    model_str: str = None
    package_id: Optional[str] = None
    vars: dict = {
        "expand": "",
        "insert_to": "",
        "partial": False,
        "encapsulated": False,
        "state": False
        }
