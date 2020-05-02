package ie.gmit.wcity.profile;

import java.util.List;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

import io.grpc.Context;
import io.grpc.stub.StreamObserver;
import io.grpc.wcity.profiles.City;
import io.grpc.wcity.profiles.CityRequestPDB;
import io.grpc.wcity.profiles.CityResponsePDB;
import io.grpc.wcity.profiles.CreateUserResponsePDB;
import io.grpc.wcity.profiles.GetAllRequest;
import io.grpc.wcity.profiles.GetUserRequestPDB;
import io.grpc.wcity.profiles.Place;
import io.grpc.wcity.profiles.PlaceRequestPDB;
import io.grpc.wcity.profiles.PlaceResponsePDB;
import io.grpc.wcity.profiles.User;
import io.grpc.wcity.profiles.UserResponsePDB;
import io.grpc.wcity.profiles.VisitCityRequestPDB;
import io.grpc.wcity.profiles.VisitCityResponsePDB;
import io.grpc.wcity.profiles.VisitPlaceRequestPDB;
import io.grpc.wcity.profiles.VisitPlaceResponsePDB;
import io.grpc.wcity.profiles.VisitedCitysRequestPDB;
import io.grpc.wcity.profiles.VisitedCitysResponsePDB;
import io.grpc.wcity.profiles.VisitedPlacesRequestPDB;
import io.grpc.wcity.profiles.VisitedPlacesResponsePDB;

public interface DataAccess {
	
	void GetAllCitysDBA(GetAllRequest request, StreamObserver<City> response);

	/**
	 * Return all places async
	 * @param request limit
	 * @param response stream observer
	 */
	 void GetAllPlacesDBA(GetAllRequest request, StreamObserver<Place> response);
	
	
	/**
	 * Create a new user
	 */
	public void createUser(User request, StreamObserver<CreateUserResponsePDB> response) ;

	/**
	 * Get user from database
	 */
	public void getUser(GetUserRequestPDB request, StreamObserver<UserResponsePDB> response) ;

	/**
	 * Create a new city
	 */
	public void createCity(City request, StreamObserver<CityResponsePDB> response) ;

	/**
	 * Get city data
	 */
	public void getCity(CityRequestPDB request, StreamObserver<CityResponsePDB> response);

	/**
	 * Get place data
	 */
	public void getPlace(PlaceRequestPDB request, StreamObserver<Place> response) ;

	/**
	 * Create a new place
	 */
	public void createPlaceRequest(Place request, StreamObserver<PlaceResponsePDB> response);

	/**
	 * Mark a place as visit for a user
	 */
	public void visitPlace(VisitPlaceRequestPDB request, StreamObserver<VisitPlaceResponsePDB> response) ;
	/**
	 * Get all visited places for a user
	 */
	public void getVisitedPlaces(VisitedPlacesRequestPDB request, StreamObserver<VisitedPlacesResponsePDB> response) ;

	/**
	 * Mark a city as visited for a user
	 */
	public void visitCity(VisitCityRequestPDB request, StreamObserver<VisitCityResponsePDB> response);

	/**
	 * Get all visited places for a user
	 */
	public void getVisitedCitys(VisitedCitysRequestPDB request, StreamObserver<VisitedCitysResponsePDB> response) ;
	/**
	 * Update user
	 */
	public void updateUserRequest(User request, StreamObserver<CreateUserResponsePDB> response) ;

	/**
	 * Update city
	 */
	public void updateCityRequest(City request, StreamObserver<CityResponsePDB> response) ;

	/**
	 * Update place
	 */
	public void updatePlaceRequest(Place request, StreamObserver<PlaceResponsePDB> response) ;

	/**
	 * Get all places from a city
	 */
	public void getCityPlaces(CityRequestPDB request, StreamObserver<VisitedPlacesResponsePDB> response) ;
}
