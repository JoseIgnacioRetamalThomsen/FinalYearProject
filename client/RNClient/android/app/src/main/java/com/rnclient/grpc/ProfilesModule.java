package com.rnclient.grpc;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.wcity.grpc.City;
import com.wcity.grpc.Geolocation;
import com.wcity.grpc.Place;
import com.wcity.grpc.ProfilesClient;


import com.facebook.react.bridge.Callback;
import com.wcity.grpc.User;

import io.grpc.wcity.profiles.GeolocationP;

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
    public void updateUser(String token, String email, String name, String description, Callback errorCallback,
                           Callback successCallback) {
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
    public void createCity(String token, String name, String country, String creatorEmail, String description, float lat, float lon, Callback errorCallback,
                           Callback successCallback) {
        City city = client.createCity(token, name, country, creatorEmail, description, GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
           // if (city.isValid() == true) {
                successCallback.invoke(city.getName(), city.getCountry(), city.getCreatorEmail(), city.getDescription(), city.getLocation());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCity(String token, String name, String country, String creatorEmail, String description, float lat, float lon, Callback errorCallback,
                        Callback successCallback) {
        City city = client.getCity(token, name, country, creatorEmail, description, GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
            //if (city.isValid() == true) {
                successCallback.invoke(city.getName(), city.getCountry(), city.getCreatorEmail(), city.getDescription(), city.getLocation());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void createPlace(String token, String name, String city, String country, String creatorEmail, String description, float lat, float lon, Callback errorCallback,
                            Callback successCallback) {
        Place place = client.createPlace(token, name, city, country, creatorEmail, description, GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
        //    if (place.isValid() == true) {
                successCallback.invoke(place.getName(), place.getCountry(), place.getCreatorEmail(), place.getDescription(), place.getLocation());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlace(String token, String name, String city, String country, String creatorEmail, String description, float lat, float lon, Callback errorCallback,
                         Callback successCallback) {
        Place place = client.getPlace(token, name, city, country, creatorEmail, description, GeolocationP.newBuilder().setLat(lat).setLon(lon).build());
        try {
          //  if (place.isValid() == true) {
                successCallback.invoke(place.getName(), place.getCountry(), place.getCreatorEmail(), place.getDescription(), place.getLocation());
//            } else {
//                errorCallback.invoke("Invalid user");
//            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
}
