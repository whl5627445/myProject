# -- coding: utf-8 --
from fastapi import APIRouter


router = APIRouter(
        prefix="/download",
        tags=["download"],
        # dependencies=[Depends(get_token_header)],
        responses={404: {"description": "Not found"}},
)
