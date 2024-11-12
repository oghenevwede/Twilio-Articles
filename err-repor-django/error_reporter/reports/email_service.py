import json
import os
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail
from dotenv import load_dotenv

load_dotenv()  # Load environment variables

def send_alert_email(report):
    # Load email addresses from config.json
    with open("config.json") as config_file:
        config = json.load(config_file)
        notification_emails = config["notification_emails"]

    # Send an email to each configured address
    for email in notification_emails:
        message = Mail(
            from_email="noreply@yourapp.com",
            to_emails=email,
            subject=f"Error Alert: {report['application']} - {report['severity']}",
            html_content=f"<strong>Error Message:</strong> {report['message']}"
        )
        
        try:
            sg = SendGridAPIClient(os.getenv("SENDGRID_API_KEY"))
            response = sg.send(message)
            print(f"Email sent to {email}, status: {response.status_code}")
        except Exception as e:
            print(f"Error sending email: {e}")
