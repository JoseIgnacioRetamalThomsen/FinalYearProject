FROM openjdk:8
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

EXPOSE 5701

ENTRYPOINT ["java", "-jar", "profilesdba.jar"]