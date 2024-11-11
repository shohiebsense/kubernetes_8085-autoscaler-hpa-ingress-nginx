kubectl autoscale deployment network-traffic-monitor-deployment-8085 --cpu-percent=3 --min=2 --max=5

# this is either use this, or use the hpa.yaml