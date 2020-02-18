package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;

import io.grpc.wcity.profiles.CityRequestP;
import io.grpc.wcity.profiles.CityResponseP;
import io.grpc.wcity.profiles.GeolocationP;

import io.grpc.wcity.profiles.PlaceRequestP;
import io.grpc.wcity.profiles.PlaceResponseP;
import io.grpc.wcity.profiles.UserResponseP;
import io.grpc.wcity.profiles.UserRequestP;
import io.grpc.wcity.profiles.ProfilesGrpc;


public class ProfilesClient {

    private final ManagedChannel channel;
    private final ProfilesGrpc.ProfilesBlockingStub stub;


    public ProfilesClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        stub = ProfilesGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

    public User getUser(String token, String email) {
        UserRequestP userRequest = UserRequestP.newBuilder().setToken(token).setEmail(email).build();
        UserResponseP response;
        User user = null;
        try {
            response = stub.getUser(userRequest);
            user = new User(response.getValid(), response.getName(), response.getDescription());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return user;
    }

    public User updateUser(String token, String email, String name, String description) {
        UserRequestP userRequest = UserRequestP
                .newBuilder()
                .setToken(token)
                .setEmail(email)
                .setName(name)
                .setDescription(description)
                .build();
        UserResponseP response;
        User user = null;
        try {
            response = stub.updateUser(userRequest);
            user = new User(response.getValid(), response.getName(), response.getDescription());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return user;
    }

    public City createCity(String token, String name, String country, String creatorEmail,
                           String description, GeolocationP location) {
        CityRequestP cityRequestP = CityRequestP.newBuilder()
                .setToken(token)
                .setName(name)
                .setCountry(country)
                .setCreatorEmail(creatorEmail)
                .setDescription(description)
                .setLocation(location)
                .build();
        CityResponseP response;

        City city = null;
        try {
            response = stub.createCity(cityRequestP);
            city = new City(response.getValid(), response.getName(), response.getCountry(),
                    response.getCreatorEmail(), response.getDescription(), response.getLocation());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return city;
    }

    public City getCity(String token, String name, String country, String creatorEmail,
                        String description, GeolocationP location) {
        CityRequestP cityRequestP = CityRequestP
                .newBuilder()
                .setToken(token)
                .setName(name)
                .setCountry(country)
                .setCreatorEmail(creatorEmail)
                .setDescription(description)
                .setLocation(location)
                .build();
        CityResponseP response;

        City city = null;
        try {
            response = stub.createCity(cityRequestP);
            city = new City(response.getValid(), response.getName(), response.getCountry(),
                    response.getCreatorEmail(), response.getDescription(), response.getLocation());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return city;
    }

    public Place createPlace(String token, String name, String city, String country,
                             String creatorEmail, String description, GeolocationP location) {
        PlaceRequestP placeRequestP = PlaceRequestP
                .newBuilder()
                .setToken(token)
                .setCity(city)
                .setCountry(country)
                .setName(name)
                .setCreatorEmail(creatorEmail)
                .setDescription(description)
                .setLocation(location)
                .build();
        PlaceResponseP response;

        Place place = null;
        try {
            response = stub.createPlace(placeRequestP);
            place = new Place(response.getValid(), response.getName(), response.getCity(),
                    response.getCountry(), response.getCreatorEmail(), response.getDescription(),
                    response.getLocation());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return place;
    }


    public Place getPlace(String token, String name, String city, String country,
                          String creatorEmail, String description, GeolocationP location) {
        PlaceRequestP placeRequestP = PlaceRequestP
                .newBuilder()
                .setToken(token)
                .setName(name)
                .setCity(city)
                .setCountry(country)
                .setCreatorEmail(creatorEmail)
                .setDescription(description)
                .setLocation(location)
                .build();
        PlaceResponseP response;

        Place place = null;
        try {
            response = stub.getPlace(placeRequestP);
            place = new Place(response.getValid(), response.getName(), response.getCity(),
                    response.getCountry(), response.getCreatorEmail(), response.getDescription(),
                    response.getLocation());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return place;
    }

//    public City updateCity(String token, String email, String name, String description) {
//        CityRequestP cityRequestP = CityRequestP.newBuilder().setToken(token).setEmail(email).setName(name).setDescription(description).build();
//        UserResponseP response;
//        City city = null;
//        try {
//            response = stub.updateUser(CityRequestP);
//            city = new City(response.getValid(), response.getName(), response.getDescription());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return city;
//    }
}
