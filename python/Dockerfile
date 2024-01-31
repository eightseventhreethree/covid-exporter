FROM python:3.12.1-slim-bookworm
LABEL maintainer="rush@ru8733.com"

ADD . /code
RUN pip install --no-cache-dir -r /code/pip-requirements.txt && \
  apt-get autoclean && \
  rm -rf /var/lib/apt/lists/* && \
  addgroup -S nonroot && \
  adduser -S nonroot -G nonroot

WORKDIR /code
ENV PYTHONPATH '/code/'
ENV PYTHONUNBUFFERED=1
USER nonroot
CMD ["python", "-u" , "/code/collector.py"]
