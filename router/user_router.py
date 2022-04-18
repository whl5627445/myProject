# -- coding: utf-8 --
from fastapi import APIRouter


router = APIRouter(
        prefix="/user",
        tags=["user"],
        # dependencies=[Depends(get_token_header)],
        responses={404: {"description": "Not found"}},
)
