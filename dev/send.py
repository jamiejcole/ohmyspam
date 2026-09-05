#!/usr/bin/env python3

import smtplib
import os
from email.mime.text import MIMEText
import random

port = int(os.getenv("MAIL_PORT", "2525"))
colours = ["red", "black", "green", "yellow", "purple", "pink", "orange"]
animals = ["monkey", "cat", "dog", "wolf", "fish", "whale", "tiger"]

pick1 = random.choice(colours)
pick2 = random.choice(animals)
fromAddr = f"{pick1}.{pick2}@example.com"

try:
    with smtplib.SMTP("localhost", port) as server:
        msg = MIMEText(f"[{pick1} {pick2}] - this is a test email from send.py ")
        msg["Subject"] = f"{pick1} !"
        msg["From"] = fromAddr
        msg["To"] = "you@example.com"
        
        server.send_message(msg)
        print(f"email sent from: {fromAddr}")
except Exception as e:
    print(f"Error: {e}")
