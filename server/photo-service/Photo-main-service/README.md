sudo docker image build -t joseretamal/photos-service:1.3 .

docker run -d -p 30051:30051 --restart always --name photos joseretamal/photos-service:1.3

sudo docker push joseretamal/photos-service:1.3
