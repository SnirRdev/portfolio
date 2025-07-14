user for argo cd for testing
admin
password 
admin1221

this is for now

to find out password if reseted
`kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d`
