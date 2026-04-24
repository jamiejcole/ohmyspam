#!/usr/bin/env python3

import smtplib
import os
from email.mime.text import MIMEText

port = int(os.getenv("MAIL_PORT", "2525"))

try:
    with smtplib.SMTP("localhost", port) as server:
        msg = MIMEText("this is a test email from send.py ")
        msg["Subject"] = "cool subject"
        msg["From"] = "me@example.com"
        msg["To"] = "you@example.com"
        
        server.send_message(msg)
        print("email sent")
except Exception as e:
    print(f"Error: {e}")
