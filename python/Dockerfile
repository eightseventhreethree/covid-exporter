FROM python:3.12.4-slim-bookworm
LABEL maintainer="rush@ru8733.com"

ADD . /code
RUN pip install --no-cache-dir -r /code/pip-requirements.txt && \
  apt-get autoclean && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /code
ENV PYTHONPATH '/code/'
ENV PYTHONUNBUFFERED=1

CMD ["python", "-u" , "/code/collector.py"]
