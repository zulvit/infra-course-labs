#!/usr/bin/env python3
import os
from http.server import BaseHTTPRequestHandler, HTTPServer


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        target = os.environ.get("FILE", "/etc/hostname")
        try:
            with open(target, "r", encoding="utf-8") as fh:
                body = fh.read()
        except OSError as exc:
            body = f"cannot read {target}: {exc}\n"
        payload = body.encode("utf-8")
        self.send_response(200)
        self.send_header("Content-Type", "text/plain; charset=utf-8")
        self.send_header("Content-Length", str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)


port = int(os.environ.get("PORT", "8888"))
HTTPServer(("0.0.0.0", port), Handler).serve_forever()
