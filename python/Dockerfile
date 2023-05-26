FROM python:3.12.0a7-slim-bullseye
LABEL maintainer="rush@ru8733.com"

ADD code /code
RUN pip install -r /code/pip-requirements.txt

WORKDIR /code
ENV PYTHONPATH '/code/'
ENV PYTHONUNBUFFERED=1

CMD ["python", "-u" , "/code/collector.py"]
