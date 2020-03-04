package com.wcity.grpc.clients;

import com.google.protobuf.ByteString;
import com.wcity.grpc.objects.ProfilePhoto;
import com.wcity.grpc.objects.ProfilePhotoResponse;

import java.util.ArrayList;
import java.util.Base64;
import java.util.List;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.photo.PhotosServiceGrpc;

import io.grpc.wcity.photo.ProfilePhotoRequestP;
import io.grpc.wcity.photo.ProfilePhotoResponseP;
import io.grpc.wcity.photo.ProfileUploadRequestP;
import io.grpc.wcity.photo.ProfileUploadResponseP;


public class PhotosClient {

    private final ManagedChannel channel;
    private final PhotosServiceGrpc.PhotosServiceBlockingStub stub;

    public PhotosClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        stub = PhotosServiceGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

    public ProfilePhotoResponse getProfilePhoto(String email, String token) {
        ProfilePhotoRequestP userData = ProfilePhotoRequestP.newBuilder().setEmail(email).setToken(token).build();
        ProfilePhotoResponseP response;
        ProfilePhotoResponse photoResponse = null;
        boolean isValid;

        try {
            response = stub.getProfilePhoto(userData);
            isValid = response.getValid();
            List<ProfilePhoto> photoList = new ArrayList<>();
            for (io.grpc.wcity.photoShared.ProfilePhoto photo : response.getPhotosList()) {
                photoList.add(new ProfilePhoto(photo.getId(), photo.getUserEmail(), photo.getUrl(),
                        photo.getTimestamp(), photo.getSelected()));
            }

            if (isValid == true)
                photoResponse = new ProfilePhotoResponse(response.getEmail(), response.getValid(),
                        photoList);
            else photoResponse = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }

    public ProfilePhoto uploadProfilePhoto(String email, String token, String image) {
        ProfileUploadRequestP userData = ProfileUploadRequestP.newBuilder().setEmail(email).setToken(token).setImage(ByteString.copyFrom(Base64.getMimeDecoder().decode(image.replaceFirst("^.*;base64,", "")))).build();

        ProfileUploadResponseP response;
        ProfilePhoto profilePhoto = null;
        boolean isSuccess;

        try {
            response = stub.uploadProfilePhoto(userData);
            isSuccess = response.getSucess();
            if (isSuccess == true)
                profilePhoto = new ProfilePhoto(response.getPhoto().getId(),
                        response.getPhoto().getUserEmail(), response.getPhoto().getUrl(),
                        response.getPhoto().getTimestamp(), response.getPhoto().getSelected());
            else profilePhoto = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return profilePhoto;
    }
}
