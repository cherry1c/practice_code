https://www.cnblogs.com/xiao987334176/p/9930517.html

docker run -d -p 9100:9100 \
  -v "/proc:/host/proc:ro" \
  -v "/sys:/host/sys:ro" \
  -v "/:/rootfs:ro" \
  prom/node-exporter

docker run  -d \
  -p 9090:9090 \
  -v /root/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yml  \
  prom/prometheus


  docker run -d \
  -p 3000:3000 \
  --name=grafana \
  -v /root/prometheus/grafana-storage:/var/lib/grafana \
  grafana/grafana

curl -XPUT http://192.168.229.133:8500/v1/agent/service/register -d@./registry.json


{
    "ID":"prometheus-143",
    "Name":"monitor",
    "Address":"192.168.229.133",
    "Port": 9090,
    "Tags": ["prometheus"],
    "Check":{
        "HTTP":"http://192.168.229.133:9090/-/healthy",
        "Interval":"10s"
    }
}