mkdir logs

docker stop porder-judicial && docker rm stop porder-judicial

docker image rm stop porder-judicial

docker build -t stop porder-judicial .

docker run -d \
--restart always \
--name stop porder-judicial \
--net=upla \
-p 8892:80 \
-v $(pwd)/logs:/etc/app-movil \
stop porder-judicial