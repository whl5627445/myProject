from fastapi import APIRouter


router = APIRouter(
        prefix="/uploadfile",
        tags=["uploadfile"],
        # dependencies=[Depends(get_token_header)],
        responses={404: {"description": "Not found"}},
)
