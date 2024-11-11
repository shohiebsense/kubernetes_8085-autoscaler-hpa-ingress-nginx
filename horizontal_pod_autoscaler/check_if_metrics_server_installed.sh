kubectl get deployment metrics-server -n kube-system

# user@usernoMacBook-Pro kubernetes_8085 % kubectl get deployment metrics-server -n kube-system
# NAME             READY   UP-TO-DATE   AVAILABLE   AGE
# metrics-server   0/1     1            0           57s


# if not normal, run the edit_deployment_metrics_server.sh script
# kubectl get deployment metrics-server -n kube-system
# NAME             READY   UP-TO-DATE   AVAILABLE   AGE
# metrics-server   1/1     1            1           12m