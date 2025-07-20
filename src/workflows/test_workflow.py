from datetime import timedelta

from temporalio import workflow
from activities import fetch_user_data, send_email, SendEmailParams


@workflow.defn
class UserNotificationWorkflow:
    @workflow.run
    async def run(self, user_id: str) -> str:
        user = await workflow.execute_activity(
            fetch_user_data,
            user_id,
            schedule_to_close_timeout=timedelta(hours=10),
        )
        email = "prat@gmail.com"
        message = "Hello"
        params = SendEmailParams(email, message)

        email_status = await workflow.execute_activity(
            send_email,
            params,
            schedule_to_close_timeout=timedelta(hours=10),
        )

        return email_status
