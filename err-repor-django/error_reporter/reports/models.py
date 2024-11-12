from django.db import models

class ErrorReport(models.Model):
    error_type = models.CharField(max_length=255)
    message = models.TextField()
    application = models.CharField(max_length=255)
    severity = models.CharField(max_length=50)
    timestamp = models.DateTimeField()
    context = models.JSONField()  # If you're using Django 3.1+ for JSON fields
    notify = models.JSONField()

    def __str__(self):
        return f"{self.application} - {self.error_type} - {self.severity}"
