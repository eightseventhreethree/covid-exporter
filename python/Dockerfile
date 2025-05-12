# syntax=docker/dockerfile:1

FROM cgr.dev/chainguard/python@sha256:3fdb6930e2b0b4631e42d99bb1d1fa9a85bc4560ce2e5bc5ce5057fe0550b9eb AS builder

WORKDIR /app

RUN python -m venv venv
ENV PATH="/app/venv/bin":$PATH
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

# chainguard/python:latest
FROM cgr.dev/chainguard/python@sha256:ed6a2d722024f3cf86347f8cc6d2b95a9f1894c7ddb84caf97a4b253a0616e5b
LABEL maintainer="rush@ru8733.com"

WORKDIR /app

COPY collector.py collector.py
COPY --from=builder /app/venv /app/venv
ENV PATH="/app/venv/bin:$PATH"

ENTRYPOINT ["python", "collector.py"]
