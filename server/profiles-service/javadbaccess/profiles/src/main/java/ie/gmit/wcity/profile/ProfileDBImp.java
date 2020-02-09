package ie.gmit.wcity.profile;

import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.logging.Logger;

import io.grpc.Context;
import io.grpc.stub.StreamObserver;
import io.grpc.wcity.profilesDB.CityPDB;
import io.grpc.wcity.profilesDB.CityRequestPDB;
import io.grpc.wcity.profilesDB.CityResponsePDB;
import io.grpc.wcity.profilesDB.CreateUserRequestPDB;
import io.grpc.wcity.profilesDB.CreateUserResponsePDB;
import io.grpc.wcity.profilesDB.GeolocationPDB;
import io.grpc.wcity.profilesDB.GetUserRequestPDB;
import io.grpc.wcity.profilesDB.PlacePDB;
import io.grpc.wcity.profilesDB.PlaceRequestPDB;
import io.grpc.wcity.profilesDB.PlaceResponsePDB;
import io.grpc.wcity.profilesDB.ProfilesDBGrpc.ProfilesDBImplBase;
import io.grpc.wcity.profilesDB.UserResponsePDB;
import io.grpc.wcity.profilesDB.VisitCityRequestPDB;
import io.grpc.wcity.profilesDB.VisitCityResponsePDB;
import io.grpc.wcity.profilesDB.VisitPlaceRequestPDB;
import io.grpc.wcity.profilesDB.VisitPlaceResponsePDB;
import io.grpc.wcity.profilesDB.VisitedCitysRequestPDB;
import io.grpc.wcity.profilesDB.VisitedCitysResponsePDB;
import io.grpc.wcity.profilesDB.VisitedPlacesRequestPDB;
import io.grpc.wcity.profilesDB.VisitedPlacesResponsePDB;

public class ProfileDBImp extends ProfilesDBImplBase {

  // private final static String URL = "bolt://10.154.0.6:7687";
  private final static String URL = "bolt://0.0.0.0:7687";
  private final static String USER_NAME = "neo4j";
  private final static String PASSWORD = "test";

  private static final ExecutorService CANCELLATION_EXECUTOR = Executors
      .newCachedThreadPool();

  private static final Logger logger = Logger
      .getLogger(ProfileDBImp.class.getName());

  public void createUser(CreateUserRequestPDB request,
      StreamObserver<CreateUserResponsePDB> response) {
    logger.info("Create user request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      dao.createUser(request.getEmail(), request.getName(),
          request.getDescription());

      response.onNext(CreateUserResponsePDB.newBuilder()
          .setEmail(request.getEmail()).setValied("true").build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void getUser(GetUserRequestPDB request,
      StreamObserver<UserResponsePDB> response) {
    logger.info("Get user request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      User u = dao.getUser(request.getEmail());

      response.onNext(UserResponsePDB.newBuilder().setEmail(request.getEmail())
          .setName(u.getName()).setDescription(u.getDescription())

          .build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void createCity(CityPDB request,
      StreamObserver<CityResponsePDB> response) {
    logger.info("Create city request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      String res = dao.createCity(
          new City().setName(request.getName()).setCountry(request.getCountry())
              .setCreatorEmail(request.getCreatorEmail())
              .setGeolocation(new Geolocation(request.getLocation().getLon(),
                  request.getLocation().getLat()))
              .setDescription(request.getDescription()));

      response.onNext(CityResponsePDB.newBuilder().setName(request.getName())
          .setCountry(request.getCountry()).build());

      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void getCity(CityRequestPDB request,
      StreamObserver<CityPDB> response) {
    logger.info("Get city request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      City u = dao.getCity(request.getName(), request.getCountry());

      response
          .onNext(CityPDB.newBuilder().setName(u.getName())
              .setCountry(u.getCountry()).setDescription(u.getDescription())
              .setCreatorEmail(u.getCreatorEmail())
              .setLocation(GeolocationPDB.newBuilder()
                  .setLat(u.getGeolocation().getLat())
                  .setLon(u.getGeolocation().getLon()))
              .build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void getPlace(PlaceRequestPDB request,
      StreamObserver<PlacePDB> response) {
    logger.info("Get place request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      Place u = dao.getPlace(request.getName(), request.getCity(),
          request.getCountry());

      response
          .onNext(PlacePDB.newBuilder().setName(u.getName())
              .setCountry(u.getCityCountry()).setDescription(u.getDescription())
              .setCreatorEmail(u.getCreatorEmail())
              .setLocation(GeolocationPDB.newBuilder()
                  .setLat(u.getGeolocation().getLat())
                  .setLon(u.getGeolocation().getLon()))
              .build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void createPlaceRequest(PlacePDB request,
      StreamObserver<PlaceResponsePDB> response) {
    logger.info("Create place request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      String res = dao.createPlace(new Place().setName(request.getName())
          .setCityName(request.getCity()).setCityCountry(request.getCountry())
          .setCreatorEmail(request.getCreatorEmail())
          .setDescription(request.getDescription())
          .setGeolocation(new Geolocation(request.getLocation().getLat(),
              request.getLocation().getLon())));

      response.onNext(PlaceResponsePDB.newBuilder().setName(request.getName())

          .build());

      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void visitPlace(VisitPlaceRequestPDB request,
      StreamObserver<VisitPlaceResponsePDB> response) {
    logger.info("Visit place request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      String res = dao.visitPlace(request.getEmail(), request.getPlaceName(),
          request.getPlaceCity(), request.getPlaceCountry()).toString();

      response.onNext(VisitPlaceResponsePDB.newBuilder()
          .setEmail(request.getEmail()).setPlaceName(request.getPlaceName())
          .setPlaceCity(request.getPlaceCity())
          .setPlaceCountry(request.getPlaceCountry()).setValid(true).build());

      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void getVisitedPlaces(VisitedPlacesRequestPDB request,
      StreamObserver<VisitedPlacesResponsePDB> response) {
    logger.info("Get Visited place request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      List<PlacePDB> res = dao.getVisitedPlaces(request.getEmail());

      response.onNext(VisitedPlacesResponsePDB.newBuilder().addAllPlaces(res)
          .setEmail(request.getEmail()).build());

      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void visitCity(VisitCityRequestPDB request,
      StreamObserver<VisitCityResponsePDB> response) {
    logger.info("Visit city request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      boolean res = dao.visitCity(request);

      response.onNext(VisitCityResponsePDB.newBuilder()
          .setEmail(request.getEmail()).setCityName(request.getCityName())
          .setCityCountry(request.getCityCountry()).setValid(true).build());

      response.onCompleted();
    } catch (Exception e) {
      // response.onError(e);
      response.onNext(VisitCityResponsePDB.newBuilder()
          .setEmail(request.getEmail()).setValid(false).build());

    }
  }

  public void getVisitedCitys(VisitedCitysRequestPDB request,
      StreamObserver<VisitedCitysResponsePDB> response) {
    logger.info("Get Visited city request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      List<CityPDB> res = dao.getVisitedCitys(request.getEmail());

      response.onNext(VisitedCitysResponsePDB.newBuilder().addAllCitys(res)
          .setEmail(request.getEmail()).build());

      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void updateUserRequest(CreateUserRequestPDB request,
      StreamObserver<CreateUserResponsePDB> response) {
    logger.info("Update user request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      boolean valid = dao.updateUser(request);

      response.onNext(
          CreateUserResponsePDB.newBuilder().setEmail(request.getEmail())
              .setValied(Boolean.toString(valid)).build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void updateCityRequest(CityPDB request,
      StreamObserver<CityResponsePDB> response) {
    logger.info("Update city request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      boolean valid = dao.updateCity(request);

      response.onNext(CityResponsePDB.newBuilder().setName(request.getName())
          .setCountry(request.getCountry()).setValid(valid).build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }

  }

  public void updatePlaceRequest(PlacePDB request,
      StreamObserver<PlaceResponsePDB> response) {
    logger.info("Update place request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      boolean valid = dao.updatePlace(request);

      response.onNext(PlaceResponsePDB.newBuilder().setName(request.getName())
          .setCountry(request.getCountry()).setCity(request.getCity())
          .setValid(valid).build());
      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

  public void getCityPlaces(CityRequestPDB request,
      StreamObserver<VisitedPlacesResponsePDB> response) {
    logger.info("Get places city request.");
    try {
      Context context = Context.current();
      DAO dao = new DAO(URL, USER_NAME, PASSWORD);

      List<PlacePDB> res = dao.getPlacesInCity(request);

      response.onNext(VisitedPlacesResponsePDB.newBuilder().addAllPlaces(res)

          .build());

      response.onCompleted();
    } catch (Exception e) {
      response.onError(e);

    }
  }

}
