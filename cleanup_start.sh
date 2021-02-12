sudo docker stop covid-prometheus-exporter && sudo docker rm covid-prometheus-exporter
sudo docker build . -t covid-prometheus-exporter
sudo docker run --detach --publish 8000:8000 --name covid-prometheus-exporter -it --restart always covid-prometheus-exporter:latest
