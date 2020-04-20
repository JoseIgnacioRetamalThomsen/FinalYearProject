# Authentication Service DBA 

This program provide access to authentications database. Is written in go and runs in a docker container,
connects to MySQL database runnig in local host.

# compile 

go build -o main *.go


# Build docker image 

sudo docker image build -t joseretamal/auth-dba:1.4 .

# Push to docker hub

docker push joseretamal/auth-dba:tagname

# run the service

sudo docker run -d -p 7777:7777 --network="host" --restart always --name authdba joseretamal/auth-dba:1.4
