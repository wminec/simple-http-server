# simple-http-server
simple http server with test app for kubernetes

# Useage

# dns check
curl :8080/dnscheck?host=google.com

# reverse dns check
curl :8080/rdnscheck?ip=0.0.0.0

# print all http Request Header
curl :8080/allreqheader
