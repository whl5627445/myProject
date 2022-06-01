# -- coding: utf-8 --
# from fastapi import FastAPI, Request, WebSocket, WebSocketDisconnect, HTTPException
from fastapi import FastAPI, Request
# from fastapi.responses import JSONResponse, Response
from app.api.simulate.simulateview import router as simulate_view_router
from app.api.model.modelview import router as model_view_router
from app.api.modelfile.upload_file import router as upload_file_router
from app.api.modelfile.download import router as download_file_router
from app.api.user.user_space import router as user_space_router
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from fastapi.staticfiles import StaticFiles
from starlette.authentication import (
    AuthenticationBackend, SimpleUser,
    AuthCredentials
    )
from config.InIt import InIt


class BasicAuthBackend(AuthenticationBackend):
    async def authenticate (self, request):
        username = request.headers.get("username", None)
        user_space = request.headers.get("space_id", None)
        if not username or username == "sys":
            username = ""
        user = SimpleUser(username)
        user.user_space = user_space
        return AuthCredentials(["simtek"]), user


app = FastAPI()


@app.middleware("http")
async def LoginAuth (request: Request, call_next):
    username = request.headers.get("username", None)
    # if not username or username == "sys":
    #     return Response(content="not found", status_code=401)
    response = await call_next(request)
    return response



app.include_router(simulate_view_router)
app.include_router(model_view_router)
app.include_router(upload_file_router)
app.include_router(download_file_router)
app.include_router(user_space_router)
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())
app.add_middleware(SessionMiddleware, secret_key="simtek")

app.mount("/static", StaticFiles(directory="./static"), name="static")
InIt()

# app.add_middleware(
#     CORSMiddleware,
#     allow_origins=["*"],
#     allow_credentials=True,
#     allow_methods=["*"],
#     allow_headers=["*"],
# )





# @app.middleware("http")
# async def add_process_time_header(request: Request, call_next):
#     response = await call_next(request)
#     if request.headers.get("Origin"):
#         response.headers["Access-Control-Allow-Origin"] = request.headers["Origin"]
#     return response

