# # -- coding: utf-8 --
# from fastapi import WebSocket, WebSocketDisconnect, Request
# from config.WebsocketManager import manager
# from config.DB_config import DBSession
# from config.redis_config import r
# from config.omc import omc
# from fastapi.responses import HTMLResponse
# from fastapi import APIRouter
#
# router = APIRouter(
#         prefix="/notice",
#         tags=["notice"],
#         # dependencies=[Depends(get_token_header)],
#         responses={404: {"description": "Not found"}},
# )
#
# session = DBSession()
#
#
# @router.websocket("/notice")
# async def NotificationView(websocket: WebSocket):
#     """
#     # 消息通知，目前是仿真日志
#     :return: 只被动接受消息， 服务器不接受数据
#     """
#     # await websocket.accept()
#     # while True:
#     #     data = await websocket.receive_text()
#     #     await websocket.send_text(f"Message text was: {data}")
#     await manager.connect(websocket)
#     # username = request.user.user_name
#     data = ""
#     try:
#         while True:
#             data = await websocket.receive_text()
#             data = r.rpop(username + "_" + "notification")
#             data = "ok"
#             await manager.send_personal_message(data, websocket)
#             # await manager.broadcast(f"Client #{client_id} says: {data}")
#     except WebSocketDisconnect:
#         manager.disconnect(websocket)
#         await manager.send_personal_message(data, websocket)
#         # await manager.broadcast(f"Client #{client_id} left the chat")
#
#
# html = """
# <!DOCTYPE html>
# <html>
#     <head>
#         <title>Chat</title>
#     </head>
#     <body>
#         <h1>WebSocket Chat</h1>
#         <form action="" onsubmit="sendMessage(event)">
#             <input type="text" id="messageText" autocomplete="off"/>
#             <button>Send</button>
#         </form>
#         <ul id='messages'>
#         </ul>
#         <script>
#             var ws = new WebSocket("ws://119.3.155.11:4327/notice/notice");
#             ws.onmessage = function(event) {
#                 var messages = document.getElementById('messages')
#                 var message = document.createElement('li')
#                 var content = document.createTextNode(event.data)
#                 message.appendChild(content)
#                 messages.appendChild(message)
#             };
#             function sendMessage(event) {
#                 var input = document.getElementById("messageText")
#                 ws.send(input.value)
#                 input.value = ''
#                 event.preventDefault()
#             }
#         </script>
#     </body>
# </html>
# """
#
#
# @router.get("/ws")
# async def get():
#     return HTMLResponse(html)
#
#
#
# @router.get("/test")
# async def _test (model_name: str, request: Request):
#
#     # username = request.user.username
#     # r.hdel("GetGraphicsData_" + username, model_name)
#     res = omc.sendExpression(model_name)
#     # res = request.auth
#     return {"msg": res,
#             # "user": request.user.display_name,
#             "auth": request.user.is_authenticated
#         }
