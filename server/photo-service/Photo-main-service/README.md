sudo docker image build -t joseretamal/photos-service:1.6 .

docker run -d -p 30051:30051 --restart always --name photos joseretamal/photos-service:1.6

sudo docker push joseretamal/photos-service:1.6
