/**
 * Jose I. Retamal
 * GMIT - 2020
 */

package ie.gmit.wcity.profile;

import java.util.List;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.logging.Logger;

import io.grpc.Context;
import io.grpc.stub.StreamObserver;
import io.grpc.wcity.profiles.ProfilesDBGrpc.ProfilesDBImplBase;
import io.grpc.wcity.profiles.*;

/**
 * Provide end points for access neo4j database.
 * 
 * @author Jose Retamal
 *
 */
public class ProfileDBImp extends ProfilesDBImplBase {

	// private final static String URL = "bolt://10.154.0.6:7687";
	private final static String URL = "bolt://0.0.0.0:7687";
	private final static String USER_NAME = "neo4j";
	private final static String PASSWORD = "test";

	private static final ExecutorService CANCELLATION_EXECUTOR = Executors.newCachedThreadPool();

	private static final Logger logger = Logger.getLogger(ProfileDBImp.class.getName());

	/**
	 * Return all city's async
	 * @param request limit
	 * @param response stream observer
	 */
	public void GetAllCitysDBA(GetAllRequest request, StreamObserver<City> response) {
		logger.info("Get All Citys.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			BlockingQueue<CityOrBuilder> queue = new ArrayBlockingQueue<CityOrBuilder>(1000);
			dao.getAllCitys(request, queue);
			while (true) {
				CityOrBuilder temp = queue.take();
				if (temp instanceof CityPoison) {
					response.onCompleted();
					break;
				}
				response.onNext((City) temp);
			}
		} catch (Exception e) {
			response.onError(e);
		}
	}

	public void GetAllPlacesDBA(GetAllRequest request, StreamObserver<City> response) {
		
	}
	
	
	/**
	 * Create a new user
	 */
	public void createUser(User request, StreamObserver<CreateUserResponsePDB> response) {
		logger.info("Create user request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			int userId = dao.createUser(request);
			response.onNext(CreateUserResponsePDB.newBuilder().setValid(true)
					.setUser(User.newBuilder().setUserId(userId).setDescripiton(request.getDescripiton())
							.setEmail(request.getEmail()).setName(request.getName()))
					.build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Get user from database
	 */
	public void getUser(GetUserRequestPDB request, StreamObserver<UserResponsePDB> response) {
		logger.info("Get user request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			User u = dao.getUser(request.getEmail());

			response.onNext(UserResponsePDB.newBuilder().setValid(true).setUser(u).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Create a new city
	 */
	public void createCity(City request, StreamObserver<CityResponsePDB> response) {
		logger.info("Create city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			int res = dao.createCity(request);

			City cityRes = request.toBuilder().setCityId(res).build();
			response.onNext(CityResponsePDB.newBuilder().setValid(true).setCity(cityRes).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Get city data
	 */
	public void getCity(CityRequestPDB request, StreamObserver<CityResponsePDB> response) {
		logger.info("Get city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			City res = dao.getCity(request.getName(), request.getCountry());
			response.onNext(CityResponsePDB.newBuilder().setValid(true).setCity(res).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Get place data
	 */
	public void getPlace(PlaceRequestPDB request, StreamObserver<Place> response) {
		logger.info("Get place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			Place u = dao.getPlace(request.getName(), request.getCity(), request.getCountry());

			response.onNext(u);
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Create a new place
	 */
	public void createPlaceRequest(Place request, StreamObserver<PlaceResponsePDB> response) {
		logger.info("Create place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			int id = dao.createPlace(request);
			Place placeRes = request.toBuilder().setPlaceId(id).build();

			response.onNext(PlaceResponsePDB.newBuilder().setValid(true).setPlace(placeRes).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);

		}
	}

	/**
	 * Mark a place as visit for a user
	 */
	public void visitPlace(VisitPlaceRequestPDB request, StreamObserver<VisitPlaceResponsePDB> response) {
		logger.info("Visit place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			String timeStamp = dao.visitPlace(request);
			response.onNext(VisitPlaceResponsePDB.newBuilder().setValid(true).setTimeStamp(timeStamp).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Get all visited places for a user
	 */
	public void getVisitedPlaces(VisitedPlacesRequestPDB request, StreamObserver<VisitedPlacesResponsePDB> response) {
		logger.info("Get Visited place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			List<Place> res = dao.getVisitedPlaces(request.getEmail());

			response.onNext(
					VisitedPlacesResponsePDB.newBuilder().addAllPlaces(res).setEmail(request.getEmail()).build());

			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);

		}
	}

	/**
	 * Mark a city as visited for a user
	 */
	public void visitCity(VisitCityRequestPDB request, StreamObserver<VisitCityResponsePDB> response) {
		logger.info("Visit city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			boolean res = dao.visitCity(request);

			response.onNext(VisitCityResponsePDB.newBuilder().setValid(true).build());

			response.onCompleted();
		} catch (Exception e) {
			// response.onError(e);
			response.onNext(VisitCityResponsePDB.newBuilder().setValid(false).build());

		}
	}

	/**
	 * Get all visited places for a user
	 */
	public void getVisitedCitys(VisitedCitysRequestPDB request, StreamObserver<VisitedCitysResponsePDB> response) {
		logger.info("Get Visited city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			List<City> res = dao.getVisitedCitys(request.getEmail());

			response.onNext(VisitedCitysResponsePDB.newBuilder().addAllCitys(res).setEmail(request.getEmail()).build());

			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);

		}
	}

	/**
	 * Update user
	 */
	public void updateUserRequest(User request, StreamObserver<CreateUserResponsePDB> response) {
		logger.info("Update user request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			boolean valid = dao.updateUser(request);
			response.onNext(CreateUserResponsePDB.newBuilder().setValid(valid).setUser(request).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}

	/**
	 * Update city
	 */
	public void updateCityRequest(City request, StreamObserver<CityResponsePDB> response) {
		logger.info("Update city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);

			boolean valid = dao.updateCity(request);

			response.onNext(CityResponsePDB.newBuilder().setCity(request).setValid(valid).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);

		}

	}

	/**
	 * Update place
	 */
	public void updatePlaceRequest(Place request, StreamObserver<PlaceResponsePDB> response) {
		logger.info("Update place request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			boolean valid = dao.updatePlace(request);
			response.onNext(PlaceResponsePDB.newBuilder().setPlace(request).setValid(valid).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);

		}
	}

	/**
	 * Get all places from a city
	 */
	public void getCityPlaces(CityRequestPDB request, StreamObserver<VisitedPlacesResponsePDB> response) {
		logger.info("Get places city request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			List<Place> res = dao.getPlacesInCity(request);
			response.onNext(VisitedPlacesResponsePDB.newBuilder().addAllPlaces(res).build());
			response.onCompleted();
		} catch (Exception e) {
			response.onError(e);
		}
	}
}
