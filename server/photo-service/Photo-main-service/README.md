sudo docker image build -t joseretamal/photos-service:1.2 .

docker run -d -p 30051:30051 --restart always --name photos joseretamal/photos-service:1.2

sudo docker push joseretamal/photos-service:1.2