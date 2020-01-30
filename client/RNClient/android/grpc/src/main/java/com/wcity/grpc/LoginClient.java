package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.helloworld.GreeterGrpc;

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
            String token = "";
            try {
                response = stub.createUser(userData);
                token = response.getToken();
            }
            catch (StatusRuntimeException e) {
                return e.toString();
            }
            catch (Exception e) {
                return e.toString();
            }
            return token;
      }

     public boolean loginUser(String email, String password) {
            UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
            UserResponse response = null;
            boolean isUser = false;
            try {
                response = stub.loginUser(userData);
                isUser = response.getIsUser();
            }
            catch (StatusRuntimeException e) {
            }
            catch (Exception e) {
            }
            return isUser;
      }

       public boolean checkToken(String token, String email) {
            LogRequest userData = LogRequest.newBuilder().setToken(token).setEmail(email).build();
            LogResponse response = null;
            boolean isSuccess = false;

            try {
                response = stub.checkToken(userData);
                isSuccess = response.getSuccess();
            } catch (StatusRuntimeException e) {
               return false;
            }
            return isSuccess;
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
                /*
                public String logout (String token, String email) {
                                    LogRequest userData = LogRequest.newBuilder().setToken(token).setEmail(email).build();
                                    LogResponse response = null;
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
