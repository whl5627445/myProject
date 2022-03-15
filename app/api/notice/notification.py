# -- coding: utf-8 --
from fastapi import WebSocket, WebSocketDisconnect, Request
from config.WebsocketManager import manager
from config.DB_config import DBSession
from config.redis_config import r
from router.notice_router import router

session = DBSession()


@router.websocket("/")
async def NotificationView(websocket: WebSocket, request: Request):
    """
    # 消息通知，目前是仿真日志
    :return: 只被动接受消息， 服务器不接受数据
    """
    await manager.connect(websocket)
    username = request.user.user_name
    data = ""
    try:
        while True:
            # data = await websocket.receive_text()
            data = r.rpop(username + "_" + "notification")
            await manager.send_personal_message(data, websocket)
            # await manager.broadcast(f"Client #{client_id} says: {data}")
    except WebSocketDisconnect:
        manager.disconnect(websocket)
        await manager.send_personal_message(data, websocket)
        # await manager.broadcast(f"Client #{client_id} left the chat")
