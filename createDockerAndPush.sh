cd ./src &&
export GOOS=linux go install &&
go build -o ../learningGo &&
cd ../
docker build -t learninggo:latest . &&
docker tag learninggo:latest joncady/mycontainers:latest &&
docker push joncady/mycontainers:latest &&
ssh -t ec2-user@ec2-54-200-18-77.us-west-2.compute.amazonaws.com \
'docker pull joncady/mycontainers:latest &&
docker stop $(docker ps -a -q --filter="name=app") &&
docker container prune -f &&
docker run -d --name app -p 80:80 joncady/mycontainers:latest &&
exit'
