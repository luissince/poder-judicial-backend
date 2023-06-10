mkdir logs

docker stop poder-judicial-backend && docker rm stop poder-judicial-backend

docker image rm stop poder-judicial-backend

docker build -t stop poder-judicial-backend .

docker run -d \
--restart always \
--name stop poder-judicial-backend \
--net=luis \
-p 8892:80 \
-v $(pwd)/logs:/etc/app \
stop poder-judicial-backend