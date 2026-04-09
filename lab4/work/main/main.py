import os

from fastapi import FastAPI
from fastapi.responses import PlainTextResponse

app = FastAPI()


@app.get("/", response_class=PlainTextResponse)
def root():
    target = os.environ.get("FILE", "/etc/hostname")
    try:
        with open(target, "r", encoding="utf-8") as fh:
            return fh.read()
    except OSError as exc:
        return f"cannot read {target}: {exc}\n"
