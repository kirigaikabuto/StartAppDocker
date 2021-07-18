import os
import signal
from flask import Flask, jsonify
from buzz import generator
import json

app = Flask(__name__)

signal.signal(signal.SIGINT, lambda s, f: os._exit(0))


@app.route("/")
def generate_buzz():
    page = '<html><body><h1>'
    page += generator.generate_buzz()
    page += '</h1></body></html>'
    return page


@app.route("/json")
def get_json():
    d = {
        "username": "asddd"
    }
    data = json.dumps(d)
    return data


if __name__ == "__main__":
    port = 5000
    portArgument = os.getenv('PORT')
    if portArgument != "" and portArgument is not None:
        port = int(portArgument)

    app.run(host='0.0.0.0', port=port)  # port 5000 is the default
