from django.urls import path
from . import views

urlpatterns = [
    path("errors", views.report_error, name="report_error"),
]
