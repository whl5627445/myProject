# -- coding: utf-8 --
from typing import Optional
from pydantic import BaseModel


class UploadSaveFile(BaseModel):
    package_name: str
    model_str: str = None
    package_id: int = None
