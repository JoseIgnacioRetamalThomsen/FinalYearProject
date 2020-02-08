package ie.gmit.wcity.profile;

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

public class ProfileDBImp extends ProfilesDBImplBase {

	//private final static String URL = "bolt://10.154.0.6:7687";
	private final static String URL = "bolt://0.0.0.0:7687";
	private final static String USER_NAME = "neo4j";
	private final static String PASSWORD = "test";

	private static final ExecutorService CANCELLATION_EXECUTOR = Executors.newCachedThreadPool();

	private static final Logger logger = Logger.getLogger(ProfileDBImp.class.getName());
	
	public void createUser(CreateUserRequestPDB request, StreamObserver<CreateUserResponsePDB> response) {
		logger.info("Create user request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			dao.createUser(request.getEmail(), request.getName(), request.getDescription());
			
			response.onNext( CreateUserResponsePDB.newBuilder().setEmail(request.getEmail()).
					setValied("true")
					.build());
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}
	
	public void getUser(GetUserRequestPDB request, StreamObserver<UserResponsePDB> response) {
		logger.info("Get user request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			User u = dao.getUser(request.getEmail());
			
			response.onNext( UserResponsePDB.newBuilder().setEmail(request.getEmail())
					.setName(u.getName())
					.setDescription(u.getDescription())
					
					.build());
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}
	
	public void createCity(CityPDB request, StreamObserver<CityResponsePDB> response) {
		logger.info("Create city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			String res = dao.createCity(new City()
					.setName(request.getName())
					.setCountry(request.getCountry())
					.setCreatorEmail(request.getCreatorEmail())
					.setGeolocation(new Geolocation(
							request.getLocation().getLon(),
							request.getLocation().getLat()
							))
					.setDescription(request.getDescription()));
			
			response.onNext( CityResponsePDB.newBuilder().setName(request.getName())
					.setCountry(request.getCountry())		
					.build());
					
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}
	
	public void getCity(CityRequestPDB request, StreamObserver<CityPDB> response) {
		logger.info("Get city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			City u = dao.getCity(request.getName(),request.getCountry());
			
			response.onNext( CityPDB.newBuilder()
					.setName(u.getName())
					.setCountry(u.getCountry())
					.setDescription(u.getDescription())
					.setCreatorEmail(u.getCreatorEmail())
					.setLocation(GeolocationPDB.newBuilder()
							.setLat(u.getGeolocation().getLat())
							.setLon(u.getGeolocation().getLon()))
					.build());
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}
	
	public void getPlace(PlaceRequestPDB request, StreamObserver<PlacePDB> response) {
		logger.info("Get place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			Place u = dao.getPlace(request.getName(),request.getCity(),request.getCountry());
			
			response.onNext( PlacePDB.newBuilder()
					.setName(u.getName())
					.setCountry(u.getCityCountry())
					.setDescription(u.getDescription())
					.setCreatorEmail(u.getCreatorEmail())
					.setLocation(GeolocationPDB.newBuilder()
							.setLat(u.getGeolocation().getLat())
							.setLon(u.getGeolocation().getLon()))
					.build());
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}

	public void createPlaceRequest(PlacePDB request, StreamObserver<PlaceResponsePDB> response) {
		logger.info("Create place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			String res = dao.createPlace(new Place()
					.setName(request.getName())
					.setCityName(request.getCity())
					.setCityCountry(request.getCountry())
					.setCreatorEmail(request.getCreatorEmail())
					.setDescription(request.getDescription())
					.setGeolocation(new Geolocation(
							request.getLocation().getLat(),request.getLocation().getLon())));
			
			response.onNext( PlaceResponsePDB.newBuilder().setName(request.getName())
						
					.build());
					
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}
	

}
