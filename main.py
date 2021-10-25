from fastapi import FastAPI, Request, Response
from app.api.simulateview import router as simulate_view_router
from app.api.modelview import router as model_view_router
from app.api.upload_file import router as upload_file_router
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from fastapi.middleware.cors import CORSMiddleware
from starlette.authentication import (
    AuthenticationBackend, AuthenticationError, SimpleUser,
    AuthCredentials
)
# import pydevd_pycharm
# pydevd_pycharm.settrace('192.168.1.110', port=10086, stdoutToServer=True, stderrToServer=True)

class BasicAuthBackend(AuthenticationBackend):
    async def authenticate(self, request):
        # if "Authorization" not in request.headers:
        #     return
        #
        # auth = request.headers["Authorization"]
        # try:
        #     scheme, credentials = auth.split()
        #     if scheme.lower() != 'basic':
        #         return
        #     decoded = base64.b64decode(credentials).decode("ascii")
        # except (ValueError, UnicodeDecodeError, binascii.Error) as exc:
        #     raise AuthenticationError('Invalid basic auth credentials')
        #
        # username, _, password = decoded.partition(":")
        username = {
            "username": "tom"
        }
        # TODO: You'd want to verify the username and password here.
        # return AuthCredentials(["authenticated"]), SimpleUser(username)
        return True, username


app = FastAPI()
app.include_router(simulate_view_router)
app.include_router(model_view_router)
app.include_router(upload_file_router)
app.add_middleware(SessionMiddleware, secret_key="simtek")
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    response = await call_next(request)
    response.headers["Access-Control-Allow-Origin"] = request.headers["Origin"]
    return response




