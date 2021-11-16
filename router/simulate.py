# -- coding: utf-8 --
from fastapi import APIRouter


router = APIRouter(
        prefix="/simulate",
        tags=["simulate"],
        # dependencies=[Depends(get_token_header)],
        responses={404: {"description": "Not found"}},
)
