# create image

  sudo docker image build -t joseretamal/post-service:1.1 .

  # push 

  sudo docker push joseretamal/post-service:1.1

  # run

  sudo docker run -d -p 10051:10051 --restart always --name post joseretamal/post-service:1.0 
