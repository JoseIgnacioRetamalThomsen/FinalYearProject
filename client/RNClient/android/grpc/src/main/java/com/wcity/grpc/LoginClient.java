package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;

import io.grpc.wcity.login.UserResponse;
import io.grpc.wcity.login.UserRequest;
import io.grpc.wcity.login.LogResponse;
import io.grpc.wcity.login.LogRequest;
import io.grpc.wcity.login.UserAuthenticationGrpc;

public class LoginClient {

    private final ManagedChannel channel;
    private final UserAuthenticationGrpc.UserAuthenticationBlockingStub stub;

    public LoginClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        stub = UserAuthenticationGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

    public String createUser(String email, String password) {
        UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
        UserResponse response;
        String token = "";
        try {
            response = stub.createUser(userData);
            token = response.getToken();
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return token;
    }

//    public String updateUser(String email, String password) {
//        UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
//        UserResponse response ;
//        boolean isUser = false;
//        String token = "";
//        try {
//            response = stub.updateUser(userData);
//            isUser = response.ge
//            tIsUser();
//            if (isUser) {
//                token = response.getToken();
//            } else token = null;
//        } catch (StatusRuntimeException e) {
//            return e.toString();
//        }
//        return token;
//    }

    public String loginUser(String email, String password) {
        UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
        UserResponse response;
        boolean isUser;
        String token = "";
        try {
            response = stub.loginUser(userData);
            isUser = response.getIsUser();
            if (isUser) {
                token = response.getToken();
            } else token = "User is not registered";
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return token;
    }

    public boolean checkToken(String token, String email) {
        LogRequest userData = LogRequest.newBuilder().setToken(token).setEmail(email).build();
        LogResponse response;
        boolean isSuccess;

        try {
            response = stub.checkToken(userData);
            isSuccess = response.getSuccess();
        } catch (StatusRuntimeException e) {
            return false;
        }
        return isSuccess;
    }


    public boolean logout(String token, String email) {
        LogRequest userData = LogRequest.newBuilder().setToken(token).setEmail(email).build();
        LogResponse response;
        boolean isSuccess;

        try {
            response = stub.logout(userData);
            isSuccess = response.getSuccess();
        } catch (StatusRuntimeException e) {
            return false;
        }
        return isSuccess;
    }
}
