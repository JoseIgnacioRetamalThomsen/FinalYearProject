package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.helloworld.GreeterGrpc;

import io.grpc.wcity.login.UserResponse;
import io.grpc.wcity.login.UserRequest;
import io.grpc.wcity.login.UserAuthenticationGrpc;


public class LoginClient {

    private final ManagedChannel channel;
    private final UserAuthenticationGrpc.UserAuthenticationBlockingStub stub;


    public LoginClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                // Channels are secure by default (via SSL/TLS). For the example we disable TLS to avoid
                // needing certificates.
                .usePlaintext()
                .build();
        stub = UserAuthenticationGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

 public String createUser(String email, String password) {
        UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
        UserResponse response = null;

        try {
            response = stub.createUser(userData);
            boolean isLoggedIn = response.getIsUser();
        } catch (StatusRuntimeException e) {
           return null;
        }
        response.getIsUser();
        return response.getToken();
    }

   /* public String updateUser(String email, String password) {
            UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
            UserResponse response = null;
            boolean isUser = false;

            try {
                response = stub.updateUser(userData);
                isUser = response.getIsUser();
            } catch (StatusRuntimeException e) {
               return e.toString();
            }
            return isUser;
        }*/
         public String loginUser(String email, String password) {
                UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
                UserResponse response = null;
                boolean isUser = false;

                try {
                    response = stub.loginUser(userData);
                    isUser = response.getIsUser();
                } catch (StatusRuntimeException e) {
                   return e.toString();
                }
                return response.getToken();
            }

       /*public String checkToken(String token, String email) {
                    UserRequest userData = UserRequest.newBuilder().setToken(token).setEmail(email).build();
                    UserResponse response = null;
                    boolean success = false;

                    try {
                        response = stub.checkToken(userData);
                        success = response.getSuccess();
                    } catch (StatusRuntimeException e) {
                       return e.toString();
                    }
                    return success;
                }
                public String logout (String token, String email) {
                                    UserRequest userData = UserRequest.newBuilder().setToken(token).setEmail(email).build();
                                    UserResponse response = null;
                                    boolean success = false;

                                    try {
                                        response = stub.logout(userData);
                                        success = response.getSuccess();
                                    } catch (StatusRuntimeException e) {
                                       return e.toString();
                                    }
                                    return success;
                                }*/
}
