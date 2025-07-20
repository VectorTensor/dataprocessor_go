import asyncio
import uuid
from datetime import timedelta
from unittest.mock import AsyncMock

from temporalio import activity
from temporalio.client import Client
from temporalio.testing import WorkflowEnvironment
from temporalio.worker import Worker

from test_workflow import UserNotificationWorkflow
from activities import fetch_user_data, send_email

async def main():

    mock_fetch_user_data = AsyncMock(return_value={
        "user_id": "user-1000",
        "name": "Debug User",
        "email": "debug@example.com",
    })

    @activity.defn(name="fetch_user_data")
    async def mock_fetch_user_data_activity(user_id: str):
        return await mock_fetch_user_data(user_id)
    client = await Client.connect("localhost:7233")


    worker = Worker(
        client,
        task_queue="test-task-queue",
        workflows=[UserNotificationWorkflow],
        activities=[mock_fetch_user_data_activity,
            send_email]
    )


    await worker.run()

if __name__ == "__main__":
    asyncio.run(main())
