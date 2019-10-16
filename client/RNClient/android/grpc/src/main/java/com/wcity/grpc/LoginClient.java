package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.helloworld.GreeterGrpc;

import io.grpc.wcity.login.LoginResponse;
import io.grpc.wcity.login.UserData;
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

    public LoginResponse Check(String email, String password) {
        UserData userData = UserData.newBuilder().setEmail(email).setHashPassword(password).build();
        LoginResponse response = null;
        try {
            response = stub.check(userData);
        } catch (StatusRuntimeException e) {

        }
        return response;
    }


}
