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
import io.grpc.wcity.profiles.VisitCityRequestP;
import io.grpc.wcity.profiles.VisitCityResponseP;
import io.grpc.wcity.profiles.VisitPlaceRequestP;
import io.grpc.wcity.profiles.VisitPlaceResponseP;
import io.grpc.wcity.profiles.VisitedCitysResponseP;
import io.grpc.wcity.profiles.VisitedPlacesResponseP;
import io.grpc.wcity.profiles.VisitedRequestP;


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
                    response.getCreatorEmail(), response.getDescription(), response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
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
                    response.getCreatorEmail(), response.getDescription(), response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return city;
    }

    public City updateCity(String token, String name, String country,
                           String creatorEmail, String description, GeolocationP location) {
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
            response = stub.updateCity(cityRequestP);
            city = new City(response.getValid(), response.getName(), response.getCountry(),
                    response.getCreatorEmail(), response.getDescription(),
                    response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return city;
    }

    public VisitCity visitCity(String token, String email, String cityName,
                               String cityCountry) {
        VisitCityRequestP visitCityRequestP = VisitCityRequestP
                .newBuilder()
                .setToken(token)
                .setEmail(email)
                .setCityName(cityName)
                .setCityCountry(cityCountry)
                .build();
        VisitCityResponseP response;

        VisitCity visitedCity = null;
        try {
            response = stub.visitCity(visitCityRequestP);
            visitedCity = new VisitCity(response.getValid(), response.getEmail(),
                    response.getCityName(), response.getCityCountry());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return visitedCity;
    }

//    public VisitedCities getVisitedCities(String token, String email) {
//        VisitedRequestP visitedRequestP = VisitedRequestP
//                .newBuilder()
//                .setToken(token)
//                .setEmail(email)
//                .build();
//        VisitedCitysResponseP response;
//
//        VisitedCities visitedCities = null;
//        try {
//            response = stub.getVisitedCitys(visitedRequestP);
//            visitedCities = new VisitedCities(response.getValid(), response.getEmail(),
//                    response.getCitys().getId());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return visitedCities;
//    }

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
                    response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
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
                    response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return place;
    }


    public Place updatePlace(String token, String name, String city, String country,
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
            response = stub.updatePlace(placeRequestP);
            place = new Place(response.getValid(), response.getName(), response.getCity(),
                    response.getCountry(), response.getCreatorEmail(), response.getDescription(),
                    response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return place;
    }

    public VisitPlace visitPlace(String token, String email, String placeName, String placeCity,
                                 String placeCountry) {
        VisitPlaceRequestP visitedPlaceRequestP = VisitPlaceRequestP
                .newBuilder()
                .setToken(token)
                .setEmail(email)
                .setPlaceName(placeName)
                .setPlaceCity(placeCity)
                .setPlaceCountry(placeCountry)
                .build();
        VisitPlaceResponseP response;

        VisitPlace visitPlace = null;
        try {
            response = stub.visitPlace(visitedPlaceRequestP);
            visitPlace = new VisitPlace(response.getValid(), response.getEmail(),
                    response.getPlaceName(), response.getPlaceCity(), response.getPlaceCountry());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return visitPlace;
    }
//    public VisitedPlaces getVisitedPlaces(String token, String email, PlaceResponseP places) {
//        VisitedRequestP visitedRequestP = VisitedRequestP
//                .newBuilder()
//                .setToken(token)
//                .setEmail(email)
//                .set(places)
//                .build();
//        VisitedPlacesResponseP response;
//
//        VisitedPlaces visitedPlaces = null;
//        try {
//            response = stub.getVisitedPlaces(visitedRequestP);
//            visitedPlaces = new VisitedPlaces(response.getValid(), response.getEmail(),
//                    response.getPlaces().getId());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return visitedPlaces;
//    }

//    public City getCityPlaces(String token, String name, String country, String creatorEmail,
//                           String description, GeolocationP location) {
//        CityRequestP cityRequestP = CityRequestP.newBuilder()
//                .setToken(token)
//                .setName(name)
//                .setCountry(country)
//                .setCreatorEmail(creatorEmail)
//                .setDescription(description)
//                .setLocation(location)
//                .build();
//        VisitedPlacesResponseP response;
//
//        City cityPlaces = null;
//        try {
//            response = stub.getCityPlaces(cityRequestP);
//            cityPlaces = new City(response.getValid(), response.getEmail(), response.getPlaces());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return cityPlaces;
//    }
}
