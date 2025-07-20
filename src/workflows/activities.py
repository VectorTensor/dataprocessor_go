from dataclasses import dataclass
from datetime import timedelta

from temporalio import activity
import asyncio

@activity.defn
async def fetch_user_data(user_id: str) -> dict:
    await asyncio.sleep(1)  # simulate I/O
    return {"user_id": user_id, "name": "Alice", "email": "alice@example.com"}

@dataclass
class SendEmailParams:
    email:str
    message:str

@activity.defn
async def send_email(params : SendEmailParams) -> str:

    return f"Email sent to {params.email} with message: {params.message}"
