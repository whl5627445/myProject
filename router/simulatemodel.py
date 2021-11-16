# -- coding: utf-8 --
from fastapi import APIRouter


router = APIRouter(
        prefix="/simulatemodel",
        tags=["simulatemodel"],
        # dependencies=[Depends(get_token_header)],
        responses={404: {"description": "Not found"}},
)
