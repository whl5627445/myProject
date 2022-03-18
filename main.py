# -- coding: utf-8 --
from fastapi import FastAPI, Request, WebSocket, WebSocketDisconnect
from fastapi.responses import JSONResponse
from app.api.simulate.simulateview import router as simulate_view_router
from app.api.model.modelview import router as model_view_router
from app.api.file.upload_file import router as upload_file_router
from app.api.file.download import router as download_file_router
# from app.api.notice.notification import router as notification_router
from fastapi.responses import HTMLResponse
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from fastapi.staticfiles import StaticFiles
from starlette.authentication import (
    AuthenticationBackend, SimpleUser,
    AuthCredentials
)

# import pydevd_pycharm
# pydevd_pycharm.settrace('192.168.1.110', port=10086, stdoutToServer=True, stderrToServer=True)


USERNAME = ""


class BasicAuthBackend(AuthenticationBackend):
    async def authenticate(self, request):
        username = request.headers.get("username", "wanghailong")
        # username = request.headers.get("username", "")
        global USERNAME
        USERNAME = username
        if not username:
            return
        return AuthCredentials(["simtek"]), SimpleUser(username)

app = FastAPI()
app.include_router(simulate_view_router)
app.include_router(model_view_router)
app.include_router(upload_file_router)
app.include_router(download_file_router)
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())
app.add_middleware(SessionMiddleware, secret_key="simtek")

app.mount("/static", StaticFiles(directory="./static"), name="static")


# app.add_middleware(
#     CORSMiddleware,
#     allow_origins=["*"],
#     allow_credentials=True,
#     allow_methods=["*"],
#     allow_headers=["*"],
# )
import logging
logging.basicConfig(level=logging.DEBUG,#控制台打印的日志级别
                    filename='/home/simtek/code/Log/OM.log',
                    filemode='a',##模式，有w和a，w就是写模式，每次都会重新写日志，覆盖之前的日志
                    #a是追加模式，默认如果不写的话，就是追加模式
                    format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s'
                    #日志格式
                    )



@app.middleware("http")
async def LoginAuth(request: Request, call_next):
    response = await call_next(request)
    return response

# @app.middleware("http")
# async def add_process_time_header(request: Request, call_next):
#     response = await call_next(request)
#     if request.headers.get("Origin"):
#         response.headers["Access-Control-Allow-Origin"] = request.headers["Origin"]
#     return response


from config.WebsocketManager import manager
from config.redis_config import r
import json
import random
MANAGER = manager
@app.websocket("/notice/")
async def NotificationView(ws: WebSocket):
    """
    # 消息通知，目前是仿真日志
    :return: 暂时只被动接受消息， 服务器不接受数据
    """
    # logging.debug(ws.__dict__)
    try:
        ws.num = random.randint(5000, 9999)
        await MANAGER.connect(ws, USERNAME)
        while True:
            res = {"status": False, "msg": ""}
            req_data = await ws.receive_text()
            req_data = json.loads(req_data)
            res_data = r.rpop(req_data.get("user", "") + "_" + "notification")
            if res_data:
                res["msg"] = res_data.decode()
                res["status"] = True
            data = json.dumps(res, ensure_ascii=False)
            # if res["status"]:
            #     await manager.send_personal_message(data, ws, USERNAME)
            # else:
            await manager.send_personal_message(data, ws, USERNAME)
    except WebSocketDisconnect:
        manager.disconnect(ws, USERNAME)
        # await manager.send_personal_message(res, websocket)



html = """
<!DOCTYPE html>
<html>
    <head>
        <title>Chat</title>
    </head>
    <body>
        <h1>WebSocket Chat</h1>
        <form action="" onsubmit="sendMessage(event)">
            <input type="text" id="messageText" autocomplete="off"/>
            <button>Send</button>
        </form>
        <ul id='messages'>
        </ul>
        <script>
            var ws = new WebSocket("ws://119.3.155.11:4327/notice/");

            ws.onmessage = function(event) {
                console.log(event.data)
                var messages = document.getElementById('messages')
                var message = document.createElement('li')
                var content = document.createTextNode(event.data)
                message.appendChild(content)
                messages.appendChild(message)
            };
            function sendMessage(event) {
                var input = document.getElementById("messageText")
                ws.send(input.value)
                input.value = ''
                event.preventDefault()
            }
        </script>
    </body>
</html>
"""


@app.get("/ws")
async def get():
    return HTMLResponse(html)
