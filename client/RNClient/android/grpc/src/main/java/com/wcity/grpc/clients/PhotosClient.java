package com.wcity.grpc.clients;

import android.annotation.TargetApi;

import com.google.gson.Gson;
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
import io.grpc.wcity.photo.GetCitysPhotoRequestP;
import io.grpc.wcity.photo.GetCitysPhotoResponseP;
import io.grpc.wcity.photo.GetPlacesPhotosPerCityRequestP;
import io.grpc.wcity.photo.GetPlacesPhotosPerCityResponseP;
import io.grpc.wcity.photo.GetPostsPhotosPerParentRequestP;
import io.grpc.wcity.photo.GetPostsPhotosPerParentResponseP;
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
import io.grpc.wcity.photoShared.CitysPhoto;
import io.grpc.wcity.photoShared.PlacesCityPhotos;
import io.grpc.wcity.photoShared.PostType;
import jdk.nashorn.internal.runtime.regexp.joni.constants.Arguments;


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
      // boolean isSuccess;
        try {
            response = stub.uploadProfilePhoto(userData);
           // isSuccess = response.getSucess();
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

    public CityPhoto uploadCityPhoto(String email, String token, int cityId, String image) {
        CityUploadRequestP userData = CityUploadRequestP.newBuilder()
                .setEmail(email)
                .setToken(token)
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
            cityPhoto = new CityPhoto();
            cityPhoto.setError(e.getMessage());
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

        try {
            response = stub.getCityImage(userData);
            List<CityPhoto> photoList = new ArrayList<>();

            for (io.grpc.wcity.photoShared.CityPhoto photo : response.getPhotosList()) {
                photoList.add(new CityPhoto(photo.getUrl(), photo.getTimestamp(), photo.getSelected()));
            }

            photoResponse = new CityPhotoResponse(response.getValid(), response.getCityID(),
                    photoList, response.getActive());
        } catch (StatusRuntimeException e) {
            photoResponse = new CityPhotoResponse();
            photoResponse.error = e.getMessage();
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

        try {
            response = stub.getPlacePhoto(userData);
            List<PlacePhoto> photoList = new ArrayList<>();
            for (io.grpc.wcity.photoShared.PlacePhoto photo : response.getPhotosList()) {
                photoList.add(new PlacePhoto(photo.getId(), photo.getPlaceId(), photo.getUrl(),
                        photo.getTimestamp(), photo.getSelected()));
            }

            photoResponse = new PlacePhotoResponse(response.getValid(), response.getPlaceId(),
                   photoList, response.getActive());

        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }

    public PlacePhoto uploadPlacePhoto(String token, String email, int placeId, String image, int placeCityId) {
        PlaceUploadRequestP userData = PlaceUploadRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setPlaceId(placeId)
                .setImage(ByteString.copyFrom(Base64.getMimeDecoder().decode(image.replaceFirst("^.*;base64,", ""))))
                .setPlaceCityId(placeCityId)
                .build();

        PlaceUploadResponseP response;
        PlacePhoto placePhoto = null;
        //boolean isSuccess;
        try {
            response = stub.uploadPlacePhoto(userData);
            //isSuccess = response.getSuccess();
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
        Gson gson = new Gson();

        PostPhotoRequestP userData = PostPhotoRequestP.newBuilder()
                .setToken(token)
                .setUserEmail(userEmail)
                .setPostId(postId)
                .build();
        PostPhotoResponseP response;
        PostPhotoResponse photoResponse = null;

        try {
            response = stub.getPostImage(userData);
            List<PostPhoto> photoList = new ArrayList<>();
            for (io.grpc.wcity.photoShared.PostPhoto photo : response.getPhotosList()) {
                photoList.add(new PostPhoto(photo.getId(), photo.getPostId(), photo.getUrl(),
                        photo.getTimestamp(), photo.getSelected()));
            }

            photoResponse = new PostPhotoResponse(response.getValid(), response.getPostId(),
                    response.getUserEmail(), gson.toJson(photoList));
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return photoResponse;
    }

    public PostPhoto uploadPostImage(String token, String userEmail, String postId, String image,
                                     int type, int parentId) {
        PostUploadRequestP userData = PostUploadRequestP.newBuilder()
                .setToken(token)
                .setUserEmail(userEmail)
                .setPostId(postId)
                .setImage(ByteString.copyFrom(Base64.getMimeDecoder()
                        .decode(image.replaceFirst("^.*;base64,", ""))))
                .setType(PostType.forNumber(type))
                .setParentId(parentId)
                .build();

        PostUploadResponseP response;
        PostPhoto postPhoto = null;
       // boolean isSuccess;
        try {
            response = stub.uploadPostImage(userData);
            //isSuccess = response.getSucess();
            // if (isSuccess == true)
            postPhoto = new PostPhoto(response.getPhoto().getId(),
                    response.getPhoto().getPostId(), response.getPhoto().getUrl(),
                    response.getPhoto().getTimestamp(), response.getPhoto().getSelected());
            // else profilePhoto = null;
        } catch (StatusRuntimeException e) {
            postPhoto = new PostPhoto();
            postPhoto.setError(e.getMessage());
            e.getMessage();
        }
        return postPhoto;
    }

    public List<CityPhoto> getCitysPhoto (String token, String email) {
        GetCitysPhotoRequestP requestP = GetCitysPhotoRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .build();

        List<CityPhoto> response = new ArrayList<>();
        GetCitysPhotoResponseP responseP;

        try{
            responseP = stub.getCitysPhotosP(requestP);
            for(CitysPhoto photo: responseP.getCityPhotosList()){
                io.grpc.wcity.photoShared.CityPhoto cityPhoto = photo.getCitysPhotosList().get(0);
                response.add(new CityPhoto(cityPhoto.getId(), cityPhoto.getCityId(), cityPhoto.getUrl(),
                cityPhoto.getTimestamp(), cityPhoto.getSelected()));
            }

        }catch(StatusRuntimeException e) {
            e.getMessage();
        }
        return response;
    }

    public List<PlacePhoto> getPlacesPerCityPhoto (String token, String email, int cityid) {
        GetPlacesPhotosPerCityRequestP requestP = GetPlacesPhotosPerCityRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setPlaceId(cityid)
                .build();

        List<PlacePhoto> response = new ArrayList<>();
        GetPlacesPhotosPerCityResponseP responseP;

        try{
            responseP = stub.getPlacesPerCityPhotoP(requestP);
            for(PlacesCityPhotos photo: responseP.getPlacePhotosList()){
                io.grpc.wcity.photoShared.PlacePhoto placePhoto = photo.getPlacePhotosList().get(0);
                response.add(new PlacePhoto(placePhoto.getId(), placePhoto.getPlaceId(), placePhoto.getUrl(),
                        placePhoto.getTimestamp(), placePhoto.getSelected()));
            }

        }catch(StatusRuntimeException e) {
            e.getMessage();
        }
        return response;
    }


    public List<PostPhoto> getPostsPhotosIdP (String token, String email, int type, int parentId) {
        GetPostsPhotosPerParentRequestP requestP = GetPostsPhotosPerParentRequestP.newBuilder()
                .setToken(token)
                .setEmail(email)
                .setType(PostType.forNumber(type) )
                .setParentId(parentId)
                .build();

        List<PostPhoto> response = new ArrayList<>();
        GetPostsPhotosPerParentResponseP responseP;

        try{
            responseP = stub.getPostsPhotosIdP(requestP);
            for(io.grpc.wcity.photoShared.PostPhoto photo: responseP.getPlacesPhotoList()){
                //PostPhoto postPhoto = photo.getPlacePhotosList().get(0);
                response.add(new PostPhoto(photo.getId(), photo.getPostId(),
                        photo.getUrl(), photo.getTimestamp(), photo.getSelected()));
            }

        }catch(StatusRuntimeException e) {
            e.getMessage();
        }
        return response;
    }
}
