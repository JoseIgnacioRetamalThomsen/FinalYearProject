package com.wcity.grpc.clients;

import android.annotation.TargetApi;

import com.google.protobuf.ByteString;
import com.wcity.grpc.objects.CityPhoto;
import com.wcity.grpc.objects.CityPhotoResponse;
import com.wcity.grpc.objects.PlacePhoto;
import com.wcity.grpc.objects.PlacePhotoResponse;
import com.wcity.grpc.objects.PostPhoto;
import com.wcity.grpc.objects.PostPhotoResponse;
import com.wcity.grpc.objects.ProfilePhoto;
import com.wcity.grpc.objects.ProfilePhotoResponse;


import java.util.ArrayList;
import java.util.Base64;
import java.util.List;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.photo.CityPhotoRequestP;

import io.grpc.wcity.photo.CityPhotoResponseP;
import io.grpc.wcity.photo.CityUploadRequestP;
import io.grpc.wcity.photo.CityUploadResponseP;
import io.grpc.wcity.photo.PhotosServiceGrpc;
import io.grpc.wcity.photo.PlacePhotoRequestP;
import io.grpc.wcity.photo.PlacePhotoResponseP;
import io.grpc.wcity.photo.PlaceUploadRequestP;
import io.grpc.wcity.photo.PlaceUploadResponseP;
import io.grpc.wcity.photo.PostPhotoRequestP;

import io.grpc.wcity.photo.PostPhotoResponseP;
import io.grpc.wcity.photo.PostUploadRequestP;
import io.grpc.wcity.photo.PostUploadResponseP;
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
        ProfilePhotoRequestP userData = ProfilePhotoRequestP.newBuilder()
                .setEmail(email)
                .setToken(token)
                .build();
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
                        photoList.get(0));
            else photoResponse = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }
    @TargetApi(11)
    public ProfilePhoto uploadProfilePhoto(String email, String token, String image) {
        ProfileUploadRequestP userData = ProfileUploadRequestP.newBuilder()
                .setEmail(email)
                .setToken(token)
                .setImage(ByteString.copyFrom(Base64.getMimeDecoder().decode(image.replaceFirst("^.*;base64,", ""))))
                .build();

        ProfileUploadResponseP response;
        ProfilePhoto profilePhoto;
        boolean isSuccess;
        try {
            response = stub.uploadProfilePhoto(userData);
            isSuccess = response.getSucess();
            // if (isSuccess == true)
            profilePhoto = new ProfilePhoto(response.getPhoto().getId(),
                    response.getPhoto().getUserEmail(), response.getPhoto().getUrl(),
                    response.getPhoto().getTimestamp(), response.getPhoto().getSelected());
            // else profilePhoto = null;
        } catch (StatusRuntimeException e) {
            profilePhoto = null;
            e.getMessage();
        }
        return profilePhoto;
    }

    public CityPhoto uploadCityPhoto(String token, String email, int cityId, String image) {
        CityUploadRequestP userData = CityUploadRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setCityId(cityId)
                .setImage(ByteString.copyFrom(Base64.getMimeDecoder().decode(image.replaceFirst("^.*;base64,", ""))))
                .build();

        CityUploadResponseP response;
        CityPhoto cityPhoto = null;
        boolean isSuccess;
        try {
            response = stub.uploadCityPhoto(userData);
            isSuccess = response.getSucess();
            // if (isSuccess == true)
            cityPhoto = new CityPhoto(response.getPhoto().getId(),
                    response.getPhoto().getCityId(), response.getPhoto().getUrl(),
                    response.getPhoto().getTimestamp(), response.getPhoto().getSelected());
            // else profilePhoto = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return cityPhoto;
    }

    public CityPhotoResponse getCityImage(String token, String email, int cityId) {
        CityPhotoRequestP userData = CityPhotoRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setCityId(cityId)
                .build();
        CityPhotoResponseP response;
        CityPhotoResponse photoResponse = null;
        boolean isValid;

        try {
            response = stub.getCityImage(userData);
            isValid = response.getValid();
            List<CityPhoto> photoList = new ArrayList<>();
            for (io.grpc.wcity.photoShared.CityPhoto photo : response.getPhotosList()) {
                photoList.add(new CityPhoto(photo.getId(), photo.getCityId(), photo.getUrl(),
                        photo.getTimestamp(), photo.getSelected()));
            }

            if (isValid == true)
                photoResponse = new CityPhotoResponse(response.getValid(), response.getCityID(),
                        photoList, response.getActive());
            else photoResponse = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }

    public PlacePhotoResponse getPlacePhoto(String token, String email, int placeId) {
        PlacePhotoRequestP userData = PlacePhotoRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setPlaceId(placeId)
                .build();
        PlacePhotoResponseP response;
        PlacePhotoResponse photoResponse = null;
        boolean isValid;

        try {
            response = stub.getPlacePhoto(userData);
            isValid = response.getValid();
            List<PlacePhoto> photoList = new ArrayList<>();
            for (io.grpc.wcity.photoShared.PlacePhoto photo : response.getPhotosList()) {
                photoList.add(new PlacePhoto(photo.getId(), photo.getPlaceId(), photo.getUrl(),
                        photo.getTimestamp(), photo.getSelected()));
            }

            if (isValid == true)
                photoResponse = new PlacePhotoResponse(response.getValid(), response.getPlaceId(),
                        photoList, response.getActive());
            else photoResponse = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }

    public PlacePhoto uploadPlacePhoto(String token, String email, int placeId, String image) {
        PlaceUploadRequestP userData = PlaceUploadRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setPlaceId(placeId)
                .setImage(ByteString.copyFrom(Base64.getMimeDecoder().decode(image.replaceFirst("^.*;base64,", ""))))
                .build();

        PlaceUploadResponseP response;
        PlacePhoto placePhoto = null;
        boolean isSuccess;
        try {
            response = stub.uploadPlacePhoto(userData);
            isSuccess = response.getSuccess();
            // if (isSuccess == true)
            placePhoto = new PlacePhoto(response.getPhoto().getId(),
                    response.getPhoto().getPlaceId(), response.getPhoto().getUrl(),
                    response.getPhoto().getTimestamp(), response.getPhoto().getSelected());
            // else profilePhoto = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return placePhoto;
    }

    public PostPhotoResponse getPostImage(String token, String userEmail, String postId) {
        PostPhotoRequestP userData = PostPhotoRequestP.newBuilder()
                .setToken(token)
                .setUserEmail(userEmail)
                .setPostId(postId)
                .build();
        PostPhotoResponseP response;
        PostPhotoResponse photoResponse = null;
        boolean isValid;

        try {
            response = stub.getPostImage(userData);
            isValid = response.getValid();
            List<PostPhoto> photoList = new ArrayList<>();
            for (io.grpc.wcity.photoShared.PostPhoto photo : response.getPhotosList()) {
                photoList.add(new PostPhoto(photo.getId(), photo.getPostId(), photo.getUrl(),
                        photo.getTimestamp(), photo.getSelected()));
            }

            if (isValid == true)
                photoResponse = new PostPhotoResponse(response.getValid(), response.getPostId(),
                        response.getUserEmail(), photoList);
            else photoResponse = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }

    public PostPhoto uploadPostImage(String token, String postId, String userEmail, String image) {
        PostUploadRequestP userData = PostUploadRequestP.newBuilder()
                .setToken(token)
                .setPostId(postId)
                .setUserEmail(userEmail)
                .setImage(ByteString.copyFrom(Base64.getMimeDecoder()
                        .decode(image.replaceFirst("^.*;base64,", ""))))
                .build();

        PostUploadResponseP response;
        PostPhoto postPhoto = null;
        boolean isSuccess;
        try {
            response = stub.uploadPostImage(userData);
            isSuccess = response.getSucess();
            // if (isSuccess == true)
            postPhoto = new PostPhoto(response.getPhoto().getId(),
                    response.getPhoto().getPostId(), response.getPhoto().getUrl(),
                    response.getPhoto().getTimestamp(), response.getPhoto().getSelected());
            // else profilePhoto = null;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return postPhoto;
    }
}
