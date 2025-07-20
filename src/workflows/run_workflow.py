import asyncio
import uuid

from temporalio.client import Client

from src.workflows.test_workflow import UserNotificationWorkflow


async def start_workflow():
    client = await Client.connect("localhost:7233")

    result = await client.execute_workflow(
        UserNotificationWorkflow.run,
        "user-42",  # argument
        id=str(uuid.uuid4()),
        task_queue="test-task-queue",  # must match worker
    )

    print("Workflow result:", result)

if __name__ == "__main__":
    asyncio.run(start_workflow())