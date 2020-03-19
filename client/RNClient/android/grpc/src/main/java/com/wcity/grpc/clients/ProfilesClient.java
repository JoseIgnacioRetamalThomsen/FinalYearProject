package com.wcity.grpc.clients;

import com.wcity.grpc.VisitedCities;
import com.wcity.grpc.objects.City;
import com.wcity.grpc.objects.Place;
import com.wcity.grpc.objects.User;

import com.google.gson.Gson;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;


//import io.grpc.wcity.profiles.CityRequestP;
import io.grpc.wcity.profiles.CityResponseP;
import io.grpc.wcity.profiles.CreateCityRequestP;
import io.grpc.wcity.profiles.CreatePlaceRequestP;
import io.grpc.wcity.profiles.CreateUserRequestP;
//import io.grpc.wcity.profiles.GeolocationP;

import io.grpc.wcity.profiles.GetAllRequest;
import io.grpc.wcity.profiles.GetCityRequestP;
import io.grpc.wcity.profiles.GetPlaceRequestP;
import io.grpc.wcity.profiles.GetUserRequestP;
import io.grpc.wcity.profiles.PlaceResponseP;
import io.grpc.wcity.profiles.UserResponseP;
import io.grpc.wcity.profiles.ProfilesGrpc;
import io.grpc.wcity.profiles.VisitCityRequestP;
import io.grpc.wcity.profiles.VisitCityResponseP;
import io.grpc.wcity.profiles.VisitPlaceRequestP;
import io.grpc.wcity.profiles.VisitPlaceResponseP;
import io.grpc.wcity.profiles.VisitedCitysResponseP;
import io.grpc.wcity.profiles.VisitedPlacesResponseP;
import io.grpc.wcity.profiles.VisitedRequestP;

import static io.grpc.okhttp.internal.Platform.logger;

public class ProfilesClient {
    //private final float TIME_CONSTANT = 10;
    private final ManagedChannel channel;
    private final ProfilesGrpc.ProfilesBlockingStub stub;

    public ProfilesClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        stub = ProfilesGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);//*10.0?
    }

    public String getAllCities(int max) {
        Gson gson = new Gson();
        GetAllRequest userRequest = GetAllRequest.newBuilder().setMax(max).build();
        Iterator<io.grpc.wcity.profiles.City> cityList;
        List<City> myc = new ArrayList<>();
        try {
            cityList = stub.getAllCitys(userRequest);
            while (cityList.hasNext()) {
                io.grpc.wcity.profiles.City city = cityList.next();
                myc.add(
                        new City(
                                city.getName(), city.getCountry(), city.getCreatorEmail(),
                                city.getLocation().getLat(), city.getLocation().getLon(),
                                city.getDescription(), city.getCityId()
                        ));
            }
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return gson.toJson(myc);
    }

    public String getAllPlaces(int max) {
        Gson gson = new Gson();
        GetAllRequest userRequest = GetAllRequest.newBuilder().setMax(max).build();
        Iterator<io.grpc.wcity.profiles.Place> placeList;
        List<Place> myp = new ArrayList<>();
        try {
            placeList = stub.getAllPlaces(userRequest);
            while (placeList.hasNext()) {
                io.grpc.wcity.profiles.Place place = placeList.next();
                myp.add(
                        new Place(
                                place.getName(), place.getCity(), place.getCountry(), place.getCreatorEmail(),
                                place.getDescription(), place.getLocation().getLat(), place.getLocation().getLon(),
                                place.getPlaceId()
                        ));
            }
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return gson.toJson(myp);
    }

    public User getUser(String token, String email) {
        GetUserRequestP userRequest = GetUserRequestP.newBuilder().setToken(token).setEmail(email).build();
        UserResponseP response;
        User user;
        //boolean isValid;
        try {
            //isValid = response.getValid();
            //if (isValid){
            response = stub.getUser(userRequest);
            user = new User(response.getUser().getEmail(), response.getUser().getName(),
                    response.getUser().getDescripiton(), response.getUser().getUserId());
            // }
        } catch (StatusRuntimeException e) {
            e.getMessage();
            user = new User("email", "name", "desc", 1);
        }
        return user;
    }

    public User updateUser(String token, String email, String userEmail, String name, String description, int userId) {
        CreateUserRequestP userRequest = CreateUserRequestP
                .newBuilder()
                .setToken(token)
                .setEmail(email)
                .setUser(io.grpc.wcity.profiles.User.newBuilder()
                        .setEmail(userEmail)
                        .setName(name)
                        .setDescripiton(description)
                        .setUserId(userId))
                .build();
        UserResponseP response;
        User user = null;
        try {
            response = stub.updateUser(userRequest);
            user = new User(response.getUser().getEmail(), response.getUser().getName(),
                    response.getUser().getDescripiton(), response.getUser().getUserId());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return user;
    }

    public City createCity(String token, String email, String name, String country, String creatorEmail,
                           float lat, float lon, String description) {
        CreateCityRequestP cityRequestP = CreateCityRequestP.newBuilder()
                .setToken(token)
                .setName(email)
                .setCity(io.grpc.wcity.profiles.City.newBuilder()
                        .setName(name)
                        .setCountry(country)
                        .setCreatorEmail(creatorEmail)
                        .setLocation(io.grpc.wcity.profiles.Geolocation.newBuilder()
                                .setLat(lat)
                                .setLon(lon)
                                .build())
                        .setDescription(description)
                        .build())
                .build();
        CityResponseP response;

        City city;
        //boolean isValid;
        try {
            response = stub.createCity(cityRequestP);
            //isValid = response.getValid();
            city = new City(response.getCity().getName(), response.getCity().getCountry(),
                    response.getCity().getCreatorEmail(), response.getCity().getLocation().getLat(),
                    response.getCity().getLocation().getLon(), response.getCity().getDescription(),
                    response.getCity().getCityId());

        } catch (StatusRuntimeException e) {
            if (e.getStatus().getCode() == Status.Code.PERMISSION_DENIED) {
                logger.log(Level.SEVERE, "Invalid token", e);

            }
            city = new City();
            city.error = e.getMessage();
            e.getMessage();
        }
        return city;
    }

    public Place createPlace(String token, String email, String name, String city, String country,
                             float lat, float lon, String description) {
        CreatePlaceRequestP placeRequestP = CreatePlaceRequestP.newBuilder()
                .setToken(token)
                .setName(email)
                .setPlace(io.grpc.wcity.profiles.Place.newBuilder()
                        .setName(name)
                        .setCity(city)
                        .setCountry(country)
                        .setLocation(io.grpc.wcity.profiles.Geolocation.newBuilder()
                                .setLat(lat)
                                .setLon(lon)
                                .build())
                        .setDescription(description)
                        .build())
                .build();
        PlaceResponseP response;
        Place place = null;
        try {
            response = stub.createPlace(placeRequestP);
            place = new Place(response.getPlace().getName(), response.getPlace().getCity(),
                    response.getPlace().getCountry(), response.getPlace().getCreatorEmail(),
                    response.getPlace().getDescription(), response.getPlace().getLocation().getLat(),
                    response.getPlace().getLocation().getLon(),
                    response.getPlace().getPlaceId());
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return place;
    }

    //    public City getCity(String token, String name, String cityName, String cityCountry) {
//        GetCityRequestP cityRequestP = GetCityRequestP
//                .newBuilder()
//                .setToken(token)
//                .setName(name)
//                .setCityName(cityName)
//                .setCityCountry(cityCountry)
//                .build();
//        CityResponseP response;
//
//        City city = null;
//        try {
//            response = stub.getCity(cityRequestP);
//            city = new City(response.getValid(), response.getName(), response.getCountry(),
//                    response.getCreatorEmail(), response.getDescription(), response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return city;
//    }
//
//    public City updateCity(String token, String name, String country,
//                           String creatorEmail, String description, GeolocationP location) {
//        CityRequestP cityRequestP = CityRequestP
//                .newBuilder()
//                .setToken(token)
//                .setName(name)
//                .setCountry(country)
//                .setCreatorEmail(creatorEmail)
//                .setDescription(description)
//                .setLocation(location)
//                .build();
//        CityResponseP response;
//
//        City city = null;
//        try {
//            response = stub.updateCity(cityRequestP);
//            city = new City(response.getValid(), response.getName(), response.getCountry(),
//                    response.getCreatorEmail(), response.getDescription(),
//                    response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return city;
//    }
//
    public boolean visitCity(String token, String email, int id) {
        VisitCityRequestP visitCityRequestP = VisitCityRequestP
                .newBuilder()
                .setToken(token)
                .setEmail(email)
                .setId(id)
                .build();
        VisitCityResponseP response;

        boolean isValid = false;
        try {
            response = stub.visitCity(visitCityRequestP);
            isValid = response.getValid();
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return isValid;
    }

    public List<City> getVisitedCities(String token, String email) {
        VisitedRequestP visitedRequestP = VisitedRequestP
                .newBuilder()
                .setToken(token)
                .setEmail(email)
                .build();
        VisitedCitysResponseP response;
        List<City> cityList = null;
        try {
            response = stub.getVisitedCitys(visitedRequestP);

            cityList = new ArrayList<>();
            for (io.grpc.wcity.profiles.City city : response.getCitysList()) {
                cityList.add(new City(city.getName(), city.getCountry(),
                        city.getCreatorEmail(), city.getLocation().getLon(),
                        city.getLocation().getLat(), city.getDescription(), city.getCityId()));
            }

        } catch (StatusRuntimeException e) {
            cityList = new ArrayList<>();
            City city = new City();
            city.error = e.getMessage();
            cityList.add(city);
            //e.getMessage();
        }
        return cityList;
    }


//    public Place getPlace(String token, String email, String placeName, String placeCity,
//                          String placeCountry) {
//        GetPlaceRequestP placeRequestP = GetPlaceRequestP
//                .newBuilder()
//                .setToken(token)
//                .setEmail(email)
//                .setPlaceName(placeName)
//                .setPlaceCity(placeCity)
//                .setPlaceCountry(placeCountry)
//                .build();
//        PlaceResponseP response;
//
//        Place place = null;
//        try {
//            response = stub.getPlace(placeRequestP);
//            place = new Place(response.getValid(), response.getName(), response.getCity(),
//                    response.getCountry(), response.getCreatorEmail(), response.getDescription(),
//                   );
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return place;
//    }


//    public Place updatePlace(String token, String name, String city, String country,
//                             String creatorEmail, String description, GeolocationP location) {
//        PlaceRequestP placeRequestP = PlaceRequestP
//                .newBuilder()
//                .setToken(token)
//                .setName(name)
//                .setCity(city)
//                .setCountry(country)
//                .setCreatorEmail(creatorEmail)
//                .setDescription(description)
//                .setLocation(location)
//                .build();
//        PlaceResponseP response;
//
//        Place place = null;
//        try {
//            response = stub.updatePlace(placeRequestP);
//            place = new Place(response.getValid(), response.getName(), response.getCity(),
//                    response.getCountry(), response.getCreatorEmail(), response.getDescription(),
//                    response.getLocation().getLat(), response.getLocation().getLon(), response.getId());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return place;
//    }
//
//    public VisitPlace visitPlace(String token, String email, String placeName, String placeCity,
//                                 String placeCountry) {
//        VisitPlaceRequestP visitedPlaceRequestP = VisitPlaceRequestP
//                .newBuilder()
//                .setToken(token)
//                .setEmail(email)
//                .setPlaceName(placeName)
//                .setPlaceCity(placeCity)
//                .setPlaceCountry(placeCountry)
//                .build();
//        VisitPlaceResponseP response;
//
//        VisitPlace visitPlace = null;
//        try {
//            response = stub.visitPlace(visitedPlaceRequestP);
//            visitPlace = new VisitPlace(response.getValid(), response.getEmail(),
//                    response.getPlaceName(), response.getPlaceCity(), response.getPlaceCountry());
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return visitPlace;
//    }
//
//    public VisitedPlaces getVisitedPlaces(String token, String email) {
//        VisitedRequestP visitedRequestP = VisitedRequestP
//                .newBuilder()
//                .setToken(token)
//                .setEmail(email)
//                .build();
//        VisitedPlacesResponseP response;
//        VisitedPlaces visitedPlaces = null;
//
//        try {
//            response = stub.getVisitedPlaces(visitedRequestP);
//
//            ArrayList<Place> placeList = new ArrayList<>();
//            for (PlaceResponseP place : response.getPlacesList()) {
//                placeList.add(new Place(place.getValid(), place.getName(), place.getCity(),
//                        place.getCountry(), place.getCreatorEmail(), place.getDescription(),
//                        place.getLocation().getLat(), place.getLocation().getLon(), place.getId()));
//            }
//            visitedPlaces = new VisitedPlaces(response.getValid(), response.getEmail(),
//                    placeList);
//        } catch (StatusRuntimeException e) {
//            e.getMessage();
//        }
//        return visitedPlaces;
//    }

    public List<Place> getCityPlaces(String token, String email, String name,
                                     String country) {
        CreateCityRequestP userRequest = CreateCityRequestP.newBuilder()
                .setToken(token)
                .setName(email)
                .setCity(io.grpc.wcity.profiles.City.newBuilder()
                        .setName(name)
                        .setCountry(country)
                        .build())
                .build();
        VisitedPlacesResponseP response;
        List<Place> placeList = new ArrayList<>();
        Place placeError;
        try {
            response = stub.getCityPlaces(userRequest);
            for (io.grpc.wcity.profiles.Place place : response.getPlacesList()) {
                placeList.add(
                        new Place(place.getName(), place.getCity(), place.getCountry(),
                                place.getCreatorEmail(), place.getDescription(),
                                place.getLocation().getLat(), place.getLocation().getLon(),
                                place.getPlaceId()
                        ));
            }
        } catch (StatusRuntimeException e) {
            placeError = new Place();
            placeError.error = e.getMessage();
            placeList.add(placeError);
            e.getMessage();
        }
        return placeList;
    }
}