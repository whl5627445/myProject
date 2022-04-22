# -- coding: utf-8 --
from typing import Optional, List
from pydantic import BaseModel


class UserSpaceModel(BaseModel):
    space_id: int

