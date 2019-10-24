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

    public String Check(String email, String password) {
        UserRequest userData = UserRequest.newBuilder().setEmail(email).setHashPassword(password).build();
        UserResponse response = null;
        String cookig = "";
        try {
            response = stub.checkUser(userData);
            cookig = response.getCookie();
        } catch (StatusRuntimeException e) {
           return e.toString();
        }
        return cookig;
    }


}
