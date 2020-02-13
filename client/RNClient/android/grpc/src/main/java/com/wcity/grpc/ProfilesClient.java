package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;

import io.grpc.wcity.profiles.CityRequestP;
import io.grpc.wcity.profiles.CityResponseP;
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
            //e.getMessage();
        }
        return user;
    }

    public User updateUser(String token, String email, String name, String description) {
        UserRequestP userRequest = UserRequestP.newBuilder().setToken(token).setEmail(email).setName(name).setDescription(description).build();
        UserResponseP response;
        User user = null;
        try {
            response = stub.updateUser(userRequest);
            user = new User(response.getValid(), response.getName(), response.getDescription());
        } catch (StatusRuntimeException e) {
            //e.getMessage();
        }
        return user;
    }

//    public City updateCity(String token, String email, String name, String description) {
//        CityRequestP cityRequestP = CityRequestP.newBuilder().setToken(token).setEmail(email).setName(name).setDescription(description).build();
//        UserResponseP response;
//        City city = null;
//        try {
//            response = stub.updateUser(CityRequestP);
//            city = new City(response.getValid(), response.getName(), response.getDescription());
//        } catch (StatusRuntimeException e) {
//            //e.getMessage();
//        }
//        return city;
//    }
}
