package com.rnclient.grpc;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import com.google.gson.Gson;
import com.wcity.grpc.objects.City;
import com.wcity.grpc.objects.Place;
import com.wcity.grpc.clients.ProfilesClient;


import com.facebook.react.bridge.Callback;
import com.wcity.grpc.objects.User;
import com.wcity.grpc.VisitCity;
import com.wcity.grpc.VisitPlace;
import com.wcity.grpc.VisitedCities;
import com.wcity.grpc.VisitedPlaces;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import io.grpc.wcity.profiles.Geolocation;


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
    public void getAllCities(int max, Callback errorCallback,
                             Callback successCallback) {
        String cityList;
        try {
            cityList = client.getAllCities(max);
            successCallback.invoke(cityList);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getAllPlaces(int max, Callback errorCallback,
                             Callback successCallback) {
        String placeList;
        try {
            placeList = client.getAllPlaces(max);
            successCallback.invoke(placeList);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getUser(String token, String email, Callback errorCallback,
                        Callback successCallback) {
        User user = client.getUser(token, email);
        if (user == null) errorCallback.invoke("error8");
        try {
            successCallback.invoke(user.getEmail(), user.getName(), user.getDescription(),
                    user.getUserId());

        } catch (Exception e) {
//            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void updateUser(String token, String email, String userEmail, String name,
                           String description, int userId,
                           Callback errorCallback, Callback successCallback) {
        User user = client.updateUser(token, email, userEmail, name, description, userId);
        try {
            //  if (user.isValid() == true) {
            successCallback.invoke(user.getEmail(), user.getName(), user.getDescription(),
                    user.getUserId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void createCity(String token, String email, String name, String country,
                           String creatorEmail, float lat, float lon, String description,
                           Callback errorCallback, Callback successCallback) {
        City city = client.createCity(token, email, name, country, creatorEmail,
                lat, lon, description);
        if (city.error != null) {
            errorCallback.invoke(city.error);
            return;
        }
        try {
            // if (city.isValid() == true) {
            successCallback.invoke(city.getCityId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }


//    @ReactMethod
//    public void getCity(String token, String name, String country, String creatorEmail,
//                        String description, float lat, float lon, Callback errorCallback,
//                        Callback successCallback) {
//        City city = client.getCity(token, name, country, creatorEmail, description,
//                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
//        try {
//            //if (city.isValid() == true) {
//            successCallback.invoke(city.getName(), city.getCountry(), city.getCreatorEmail(),
//                    city.getDescription(), city.getLat(), city.getLon(), city.getId());
////            } else {
////                errorCallback.invoke("Invalid user");
////            }
//        } catch (Exception e) {
//            errorCallback.invoke(e.getMessage());
//        }
//    }
//
//    @ReactMethod
//    public void updateCity(String token, String name, String country, String creatorEmail,
//                           String description, float lat, float lon, Callback errorCallback,
//                           Callback successCallback) {
//        City city = client.updateCity(token, name, country, creatorEmail, description,
//                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
//        try {
//            //if (city.isValid() == true) {
//            successCallback.invoke(city.getName(), city.getCountry(), city.getCreatorEmail(),
//                    city.getDescription(), city.getLat(), city.getLon(), city.getId());
////            } else {
////                errorCallback.invoke("Invalid user");
////            }
//        } catch (Exception e) {
//            errorCallback.invoke(e.getMessage());
//        }
//    }

    @ReactMethod
    public void visitCity(String token, String email, int id, Callback errorCallback,
                          Callback successCallback) {
        boolean isValid = client.visitCity(token, email, id);
        try {
            successCallback.invoke(isValid);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getVisitedCities(String token, String email, Callback errorCallback,
                                 Callback successCallback) {
        Gson gson = new Gson();
        List<City> visitedCities = client.getVisitedCities(token, email);

        if (visitedCities.size() == 1) {
            errorCallback.invoke(visitedCities.get(0).error);
            return;
        }

        try {
            successCallback.invoke(gson.toJson(visitedCities));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void createPlace(String token, String email, String name, String city, String country,
                            float lat, float lon, String description,
                            Callback errorCallback, Callback successCallback) {
        Place place = client.createPlace(token, email, name, city, country,
                lat, lon, description);
        try {
            //    if (place.isValid() == true) {
            successCallback.invoke(place.getId());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    //
//    @ReactMethod
//    public void getPlace(String token, String name, String city, String country, String creatorEmail,
//                         String description, float lat, float lon, Callback errorCallback,
//                         Callback successCallback) {
//        Place place = client.getPlace(token, name, city, country, creatorEmail, description,
//                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
//        try {
//            //  if (place.isValid() == true) {
//            successCallback.invoke(place.getName(), place.getCity(), place.getCountry(), place.getCreatorEmail(),
//                    place.getDescription(), place.getLat(), place.getLon(), place.getId());
////            } else {
////                errorCallback.invoke("Invalid user");
////            }
//        } catch (Exception e) {
//            errorCallback.invoke(e.getMessage());
//        }
//    }
//
//    @ReactMethod
//    public void updatePlace(String token, String name, String city, String country, String creatorEmail,
//                            String description, float lat, float lon, Callback errorCallback,
//                            Callback successCallback) {
//        Place place = client.updatePlace(token, name, city, country, creatorEmail, description,
//                GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
//        try {
//            //  if (place.isValid() == true) {
//            successCallback.invoke(place.getName(), place.getCountry(), place.getCreatorEmail(),
//                    place.getDescription(), place.getLat(), place.getLon(), place.getId());
////            } else {
////                errorCallback.invoke("Invalid user");
////            }
//        } catch (Exception e) {
//            errorCallback.invoke(e.getMessage());
//        }
//    }
//
    @ReactMethod
    public void visitPlace(String token, String email, int placeId, Callback errorCallback,
                           Callback successCallback) {
        boolean isValid = client.visitPlace(token, email, placeId);
        try {
            successCallback.invoke(isValid);

        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getVisitedPlaces(String token, String email, Callback errorCallback,
                                 Callback successCallback) {
        Gson gson = new Gson();
        List<Place> visitedPlaces = client.getVisitedPlaces(token, email);
        try {
            successCallback.invoke(gson.toJson(visitedPlaces));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCityPlaces(String token, String email, String name, String country,
                              Callback errorCallback, Callback successCallback) {
        Gson gson = new Gson();
        List<Place> placeList = client.getCityPlaces(token, email, name, country);
//        if(placeList.get(0).error != null){
//            errorCallback.invoke(placeList.get(0).error); return;
//        }
        try {
            successCallback.invoke(gson.toJson(placeList));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

}
