package com.rnclient.grpc;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.wcity.grpc.City;
import com.wcity.grpc.Place;
import com.wcity.grpc.ProfilesClient;


import com.facebook.react.bridge.Callback;
import com.wcity.grpc.User;
import com.wcity.grpc.VisitCity;
import com.wcity.grpc.VisitPlace;
import com.wcity.grpc.VisitedCities;
import com.wcity.grpc.VisitedPlaces;

import java.util.ArrayList;

import io.grpc.wcity.profiles.GeolocationP;
import io.grpc.wcity.profiles.PlaceResponseP;

public class ProfilesModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.216.42";
    private static final int PORT_NUMBER = 60051;
    private static ProfilesClient client;

    public ProfilesModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
        client = new ProfilesClient(IP_ADDRESS, PORT_NUMBER);
    }

    @Override
    public String getName() {
        return "ProfilesModule";
    }

    @ReactMethod
    public void getUser(String token, String email, Callback errorCallback,
                        Callback successCallback) {
        User user = client.getUser(token, email);

        try {
            if (user.isValid() == true) {
                successCallback.invoke(user.getName(), user.getDescription());
            } else {
                errorCallback.invoke("Invalid user");
            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void updateUser(String token, String email, String name, String description,
                           Callback errorCallback, Callback successCallback) {
        User user = client.updateUser(token, email, name, description);
        try {
            //  if (user.isValid() == true) {
            successCallback.invoke(user.getName(), user.getDescription());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void createCity(String token, String name, String country, String creatorEmail,
                           String description, float lat, float lon, Callback errorCallback,
                           Callback successCallback) {
        City city = client.createCity(token, name, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            // if (city.isValid() == true) {
            successCallback.invoke(city.getName(), city.getCountry(),
                    city.getCreatorEmail(), city.getDescription(), city.getLat(), city.getLon(), city.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCity(String token, String name, String country, String creatorEmail,
                        String description, float lat, float lon, Callback errorCallback,
                        Callback successCallback) {
        City city = client.getCity(token, name, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            //if (city.isValid() == true) {
            successCallback.invoke(city.getName(), city.getCountry(), city.getCreatorEmail(),
                    city.getDescription(), city.getLat(), city.getLon(), city.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void updateCity(String token, String name, String country, String creatorEmail,
                           String description, float lat, float lon, Callback errorCallback,
                           Callback successCallback) {
        City city = client.updateCity(token, name, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            //if (city.isValid() == true) {
            successCallback.invoke(city.getName(), city.getCountry(), city.getCreatorEmail(),
                    city.getDescription(), city.getLat(), city.getLon(), city.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void visitCity(String token, String email, String cityName,
                          String cityCountry, Callback errorCallback,
                          Callback successCallback) {
        VisitCity visitCity = client.visitCity(token, email, cityName, cityCountry);
        try {
            //  if (visitCity.isValid() == true) {
            successCallback.invoke(visitCity.isValid(), visitCity.getEmail(),
                    visitCity.getCityName(), visitCity.getCityCountry()
            );
            //            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getVisitedCities(String token, String email,
                                 Callback errorCallback,
                                 Callback successCallback) {
        VisitedCities visitedCities = client.getVisitedCities(token, email);
        try {
            //  if (visitedCities.isValid() == true) {
            successCallback.invoke(visitedCities.isValid(), visitedCities.getEmail(),
                    visitedCities.getVisitedCities());
            //            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void createPlace(String token, String name, String city, String country,
                            String creatorEmail, String description, float lat, float lon,
                            Callback errorCallback, Callback successCallback) {
        Place place = client.createPlace(token, name, city, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            //    if (place.isValid() == true) {
            successCallback.invoke(place.getName(), place.getCountry(), place.getCreatorEmail(),
                    place.getDescription(), place.getLat(), place.getLon(), place.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlace(String token, String name, String city, String country, String creatorEmail,
                         String description, float lat, float lon, Callback errorCallback,
                         Callback successCallback) {
        Place place = client.getPlace(token, name, city, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            //  if (place.isValid() == true) {
            successCallback.invoke(place.getName(), place.getCountry(), place.getCreatorEmail(),
                    place.getDescription(), place.getLat(), place.getLon(), place.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void updatePlace(String token, String name, String city, String country, String creatorEmail,
                            String description, float lat, float lon, Callback errorCallback,
                            Callback successCallback) {
        Place place = client.updatePlace(token, name, city, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            //  if (place.isValid() == true) {
            successCallback.invoke(place.getName(), place.getCountry(), place.getCreatorEmail(),
                    place.getDescription(), place.getLat(), place.getLon(), place.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void visitPlace(String token, String email, String placeName,
                           String placeCity, String placeCountry,
                           Callback errorCallback,
                           Callback successCallback) {
        VisitPlace visitPlace = client.visitPlace(token, email, placeName, placeCity, placeCountry);
        try {
            //  if (visitPlace.isValid() == true) {
            successCallback.invoke(visitPlace.isValid(), visitPlace.getEmail(),
                    visitPlace.getPlaceName(), visitPlace.getPlaceCity(), visitPlace.getPlaceCountry());
            //            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getVisitedPlaces(String token, String email,
                                 Callback errorCallback,
                                 Callback successCallback) {
        VisitedPlaces visitedPlaces = client.getVisitedPlaces(token, email);
        try {
            //  if (visitedPlaces.isValid() == true) {
            successCallback.invoke(visitedPlaces.isValid(), visitedPlaces.getEmail(), visitedPlaces.getVisitedPlaces());
            //            } else {
            errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCityPlaces(String token, String name, String country, String creatorEmail,
                              String description, float lat, float lon, Callback errorCallback,
                              Callback successCallback) {
        VisitedPlaces cityPlaces = client.getCityPlaces(token, name, country, creatorEmail, description,
                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            // if (city.isValid() == true) {
            successCallback.invoke(cityPlaces.isValid(), cityPlaces.getEmail(), cityPlaces.getVisitedPlaces());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

}
