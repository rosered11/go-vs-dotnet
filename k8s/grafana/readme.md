## Read

helm install -f values/grafana.yaml grafana grafana/grafana 

user: admin
pass: kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo