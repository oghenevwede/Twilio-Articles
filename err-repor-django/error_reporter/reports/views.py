import json
import logging
import os
from django.shortcuts import render
from django.http import JsonResponse
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail
from dotenv import load_dotenv
from django.views.decorators.csrf import csrf_exempt

# Load environment variables
load_dotenv()

# Load SendGrid API key
SENDGRID_API_KEY = os.getenv('SENDGRID_API_KEY')

# Configure logging
logger = logging.getLogger(__name__)

# Load team email configurations
def load_team_emails():
    with open('config.json') as f:
        config = json.load(f)
    return config["team_emails"]

# Send error report to the selected team
def send_email(to_email, subject, content):
    message = Mail(
        from_email='vwede@mymoyo.com.ng',
        to_emails=to_email,
        subject=subject,
        html_content=content)
    
    try:
        sg = SendGridAPIClient(SENDGRID_API_KEY)
        response = sg.send(message)
        return response
    except Exception as e:
        logger.error(f"Failed to send email: {e}")
        return None

# Report error view
@csrf_exempt
def report_error(request):
    if request.method == 'POST':
        # Get the payload from the request
        payload = json.loads(request.body)

        # Extract details from the payload
        error_type = payload.get("error_type")
        message = payload.get("message")
        application = payload.get("application")
        severity = payload.get("severity")
        timestamp = payload.get("timestamp")
        context = payload.get("context")
        notify_teams = payload.get("notify", [])

        # Log the error
        logger.error(f"Error Type: {error_type}, Message: {message}, Severity: {severity}, Timestamp: {timestamp}, Context: {context}")

        # Load team emails
        team_emails = load_team_emails()

        # Send emails to selected teams
        for team in notify_teams:
            if team in team_emails:
                email = team_emails[team]
                subject = f"Error Report: {error_type} - {severity}"
                content = f"<h3>Application: {application}</h3><p><b>Error:</b> {message}</p><p><b>Severity:</b> {severity}</p><p><b>Timestamp:</b> {timestamp}</p>"
                response = send_email(email, subject, content)

                if response:
                    return JsonResponse({"message": "Error report sent successfully."}, status=200)
                else:
                    return JsonResponse({"message": "Failed to send error report."}, status=500)

        return JsonResponse({"message": "No valid teams to notify."}, status=400)

    return JsonResponse({"message": "Invalid request method."}, status=400)
