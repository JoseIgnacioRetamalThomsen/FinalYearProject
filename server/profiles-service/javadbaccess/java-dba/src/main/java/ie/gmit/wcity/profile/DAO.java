/**
 * Jose I. Retamal
 * GMIT - 2020
 */

package ie.gmit.wcity.profile;

import static org.neo4j.driver.Values.parameters;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

import org.neo4j.driver.AuthTokens;
import org.neo4j.driver.Driver;
import org.neo4j.driver.GraphDatabase;
import org.neo4j.driver.Record;
import org.neo4j.driver.Result;
import org.neo4j.driver.Session;
import org.neo4j.driver.Transaction;
import org.neo4j.driver.TransactionWork;

import io.grpc.wcity.profiles.City;
import io.grpc.wcity.profiles.CityOrBuilder;
import io.grpc.wcity.profiles.CityRequestPDB;
import io.grpc.wcity.profiles.Geolocation;
import io.grpc.wcity.profiles.GetAllRequest;
import io.grpc.wcity.profiles.Place;
import io.grpc.wcity.profiles.PlaceOrBuilder;
import io.grpc.wcity.profiles.User;
import io.grpc.wcity.profiles.User.Builder;
import io.grpc.wcity.profiles.VisitCityRequestPDB;
import io.grpc.wcity.profiles.VisitPlaceRequestPDB;
import profiles.Test;

/**
 * Provide access to neo4j database
 * 
 * @author Jose I. Retamal
 *
 */
public class DAO implements AutoCloseable {

	private final Driver driver;

	public DAO(String uri, String user, String password) {
		driver = GraphDatabase.driver(uri, AuthTokens.basic(user, password));
	}

	/**
	 * Returns all city's
	 * @param request city's request
	 * @param queue blocking queue with all city's
	 */
	public void getAllCitys(GetAllRequest request, BlockingQueue<CityOrBuilder> queue) {
		try (Session session = driver.session()) {
			int result = session.writeTransaction(new TransactionWork<Integer>() {
				@Override
				public Integer execute(Transaction tx) {
					Result result = tx.run("MATCH (a:City)" + "RETURN a,ID(a)");

					for (Record r : result.list()) {
						queue.offer(
								City.newBuilder().setName(r.get(0).get("name").asString())
										.setCountry(r.get(0).get("country").asString())
										.setCreatorEmail(r.get(0).get("creatorEmail").asString())
										.setDescription(r.get(0).get("description").asString())
										.setLocation(Geolocation.newBuilder().setLat(r.get(0).get("lat").asFloat())
												.setLon(r.get(0).get("lon").asFloat()))
										.setCityId(r.get(1).asInt()).build());
					}

					queue.offer(new CityPoison());
					return 1;
				}
			});
		}

	}

	/**
	 * Return all places
	 * 
	 * @param request the request
	 * @param queue   all places collection
	 */
	public void getAllPlaces(GetAllRequest request, BlockingQueue<PlaceOrBuilder> queue) {
		try (Session session = driver.session()) {
			int result = session.writeTransaction(new TransactionWork<Integer>() {
				@Override
				public Integer execute(Transaction tx) {
					Result result = tx.run("MATCH (a:Place)" + "RETURN a,ID(a)");
					
					for (Record r : result.list()) {
						queue.offer(
								Place.newBuilder().setName(r.get(0).get("name").asString())
										.setCity(r.get(0).get("city").asString())
										.setCountry(r.get(0).get("country").asString())
										.setCreatorEmail(r.get(0).get("creatorEmail").asString())
										.setDescription(r.get(0).get("description").asString())
										.setLocation(Geolocation.newBuilder().setLat(r.get(0).get("lat").asFloat())
												.setLon(r.get(0).get("lon").asFloat()))
										.setPlaceId(r.get(1).asInt()).build());
					}
					queue.offer(new PlacePoison());
					return 1;
				}
			});
		}

	}

	/**
	 * Create a new user
	 * 
	 * @param email       user email
	 * @param name        user name
	 * @param description user description
	 * @return true if is created
	 */
	public int createUser(User user) {
		try (Session session = driver.session()) {
			Integer result = session.writeTransaction(new TransactionWork<Integer>() {
				@Override
				public Integer execute(Transaction tx) {
					Result result = tx.run(
							"Create (a:User) " + "SET a.name = $name " + "SET a.email = $email "
									+ "SET a.description = $description " + "RETURN   id(a)",
							parameters("name", user.getName().toLowerCase(), "email", user.getEmail().toLowerCase(), "description",
									user.getDescripiton()));
					return result.single().get(0).asInt();
				}
			});
			return result;

		}
	}

	/**
	 * Get user data
	 * 
	 * @param email user email
	 * @return user data
	 */
	public User getUser(final String email) {
		User user = null;

		try (Session session = driver.session()) {
			user = session.writeTransaction(new TransactionWork<User>() {
				@Override
				public User execute(Transaction tx) {
					Result result = tx.run("MATCH (a:User) " + "where a.email = $email " + "RETURN a",
							parameters("email", email.toLowerCase()));
					Record r = result.next();
					User u = User.newBuilder().setEmail(email).setName(r.get(0).get("name").asString())
							.setDescripiton(r.get(0).get("description").asString()).build();
					return u;
				}
			});
		}
		return user;
	}

	/**
	 * Create a new city
	 * 
	 * @param city the new city
	 * @return true if created
	 */
	public int createCity(City city) {
		try (Session session = driver.session()) {
			int result = session.writeTransaction(new TransactionWork<Integer>() {

				@Override
				public Integer execute(Transaction tx) {
					Result result = tx.run(
							"Create (a:City) " + "SET a.name = $name " + "SET a.creatorEmail = $creatorEmail "
									+ "SET a.description = $description " + "SET a.lat = $lat " + "SET a.lon = $lon "
									+ "SET a.country = $country " + "RETURN  id(a)",
							parameters("name", city.getName().toLowerCase(), "creatorEmail", city.getCreatorEmail().toLowerCase(), "description",
									city.getDescription(), "lat", city.getLocation().getLat(), "lon",
									city.getLocation().getLon(), "country", city.getCountry().toLowerCase()));

					return result.single().get(0).asInt();
				}
			});
			return result;

		}
	}

	/**
	 * Get city data
	 * 
	 * @param name    city name
	 * @param country city country
	 * @return city data
	 */
	public City getCity(String name, String country) {
		City city = null;
		try (Session session = driver.session()) {
			city = session.writeTransaction(new TransactionWork<City>() {
				@Override
				public City execute(Transaction tx) {
					Result result = tx.run("MATCH (a:City) " + "where a.name = $name " + " AND "
							+ " a.country = $country " + "RETURN a,ID(a)",
							parameters("name", name.toLowerCase(), "country", country.toLowerCase()));
					Record r = result.next();

					return City.newBuilder().setName(r.get(0).get("name").asString())
							.setCountry(r.get(0).get("country").asString())
							.setCreatorEmail(r.get(0).get("creatorEmail").asString())
							.setDescription(r.get(0).get("description").asString()).setLocation(Geolocation.newBuilder()
									.setLat(r.get(0).get("lat").asFloat()).setLon(r.get(0).get("lon").asFloat()))
							.setCityId(r.get(1).asInt()).build();

				}
			});
		}
		return city;
	}

	/**
	 * Create a new place
	 * 
	 * @param place the new place
	 * @return true if created
	 */
	public int createPlace(Place place) {
		try (Session session = driver.session()) {
			int result = session.writeTransaction(new TransactionWork<Integer>() {

				@Override
				public Integer execute(Transaction tx) {
					Result result = tx.run(
							"Create (a:Place) " + "SET a.name = $name " + "SET a.creatorEmail = $creatorEmail "
									+ "SET a.description = $description " + "SET a.lat = $lat " + "SET a.lon = $lon "
									+ "SET a.country = $country " + "SET a.city = $city " + "RETURN id(a)",
							parameters("name", place.getName().toLowerCase(), "creatorEmail", place.getCreatorEmail().toLowerCase(), "description",
									place.getDescription(), "lat", place.getLocation().getLat(), "lon",
									place.getLocation().getLon(), "country", place.getCountry().toLowerCase(), "city",
									place.getCity().toLowerCase()));

					// add relation
					tx.run("MATCH (a:City),(b:Place) " + "WHERE a.name = $cityName AND  a.country = $country "
							+ "AND b.name = $placeNane AND b.country =$country AND b.city = $cityName "
							+ " CREATE (a)-[r:ISIN]->(b)" + "RETURN r",
							parameters("cityName", place.getCity().toLowerCase(), "country", place.getCountry().toLowerCase(), "placeNane",
									place.getName().toLowerCase(), "city", place.getCity().toLowerCase()));
					return result.single().get(0).asInt();
				}
			});
			return result;

		} catch (Exception e) {
			e.printStackTrace();
			return -1;
		}
	}

	/**
	 * Get a place
	 * 
	 * @param name    the place name
	 * @param city    the place city
	 * @param country the place country
	 * @return the place data
	 */

	public Place getPlace(String name, String city, String country) {
		Place place = null;
		try (Session session = driver.session()) {
			place = session.writeTransaction(new TransactionWork<Place>() {

				@Override
				public Place execute(Transaction tx) {
					Result result = tx.run(
							"MATCH (a:Place) " + "where a.name = $name " + " AND " + " a.country = $country " + " AND"
									+ " a.city = $city " + "RETURN a,ID(a)",
							parameters("name", name.toLowerCase(), "country", country.toLowerCase(), "city", city.toLowerCase()));
					Record r = result.next();

					return Place.newBuilder().setName(r.get(0).get("name").asString())
							.setCity(r.get(0).get("city").asString()).setCountry(r.get(0).get("country").asString())
							.setCreatorEmail(r.get(0).get("creatorEmail").asString())
							.setDescription(r.get(0).get("description").asString()).setLocation(Geolocation.newBuilder()
									.setLat(r.get(0).get("lat").asFloat()).setLon(r.get(0).get("lon").asFloat()))
							.setPlaceId(r.get(1).asInt()).build();
				}
			});
		} catch (Exception e) {
			// not found
			return place;
		}
		return place;
	}

	/**
	 * Mark a place as visited
	 * 
	 * @param userEmail    the user email
	 * @param placeName    the place name
	 * @param placeCity    the place city
	 * @param placeCountry the place country
	 * @return true if sucess
	 */

	public String visitPlace(String userEmail, String placeName, String placeCity, String placeCountry) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {

				@Override
				public String execute(Transaction tx) {
					Result result = tx.run("Match (a:User) " + "WHERE a.email = $email " + "MATCH (b:Place) "
							+ "WHERE b.name = $name " + "AND " + "b.city =$city " + " AND b.country = $country "
							+ "CREATE (a)-[r:VISIT{date:datetime()}]->(b)" + "RETURN r.date  AS currentDateTime",
							parameters("email", userEmail.toLowerCase(), "name", placeName.toLowerCase(), "city", placeCity.toLowerCase(), "country",
									placeCountry));

					return result.single().get(0).toString();
				}
			});
			return result;
		} catch (

		Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * Mark a place as visited
	 * 
	 * @param placeId
	 * @return the timestamp
	 */
	public String visitPlace(VisitPlaceRequestPDB request) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {

				@Override
				public String execute(Transaction tx) {
					Result result = tx.run(
							"Match (a:User) " + "WHERE a.email = $email " + "MATCH (b:Place) " + "WHERE id(b) = $id "
									+ "CREATE (a)-[r:VISIT{date:datetime()}]->(b)"
									+ "RETURN r.date  AS currentDateTime",
							parameters("email", request.getUserEmail().toLowerCase(), "id", request.getPlaceId()));

					return result.single().get(0).toString();
				}
			});
			return result;
		} catch (

		Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * Return a list of all visited places
	 * 
	 * @param userEmail the user email
	 * @return list with all visited places
	 */

	public List<Place> getVisitedPlaces(String userEmail) {
		List<Place> places = new ArrayList<>();
		try (Session session = driver.session()) {
			boolean result = session.writeTransaction(new TransactionWork<Boolean>() {

				@Override
				public Boolean execute(Transaction tx) {
					Result result = tx
							.run("Match (a:User) " + " -[r:VISIT]-> " + " (b:Place) " + "WHERE a.email = $email " +

									"RETURN b,id(b)", parameters("email", userEmail.toLowerCase())); // return
					// result.single().get(0).asString();
					for (Record r : result.list()) {
						places.add(Place.newBuilder().setName(r.get(0).get("name").asString())
								.setCity(r.get(0).get("city").asString()).setCountry(r.get(0).get("country").asString())
								.setCreatorEmail(r.get(0).get("creatorEmail").asString())
								.setDescription(r.get(0).get("description").asString())
								.setLocation(Geolocation.newBuilder().setLat(r.get(0).get("lat").asFloat())
										.setLon(r.get(0).get("lon").asFloat()).build())
								.setPlaceId(r.get(1).asInt()).build());
					}
					return true;
				}
			});
			return places;
		} catch (Exception e) {
			e.printStackTrace();
			return places;
		}
	}

	/**
	 * MArk a city as visited
	 * 
	 * @param request the user request
	 * @return list with all visited city's
	 */

	public boolean visitCity(VisitCityRequestPDB request) {
		try (Session session = driver.session()) {
			List<Record> result = session.writeTransaction(new TransactionWork<List<Record>>() {

				@Override
				public List<Record> execute(Transaction tx) {
					Result result = tx.run(
							"Match (a:User) " + "WHERE a.email = $email " + "MATCH (b:City) " + "WHERE id(b) = $id "
									+ "CREATE (a)-[r:VISIT{date:date()}]->(b)" + "RETURN r.date",
							parameters("email", request.getUserEmail().toLowerCase(), "id", request.getCityId()));

					return result.list();
				}
			});
			return true;
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
	}

	/**
	 * Returns all visited city;s for a user
	 * 
	 * @param userEmail the user email
	 * @return list with all visited citys
	 */

	public List<City> getVisitedCitys(String userEmail) {
		List<City> citys = new ArrayList<>();
		try (Session session = driver.session()) {
			boolean result = session.writeTransaction(new TransactionWork<Boolean>() {
				@Override
				public Boolean execute(Transaction tx) {
					Result result = tx.run("Match (a:User) " + " -[r:VISIT]-> " + " (b:City) "
							+ "WHERE a.email = $email " + "RETURN b,id(b)", parameters("email", userEmail.toLowerCase())); 
			
					for (Record r : result.list()) {
						citys.add(City.newBuilder().setName(r.get(0).get("name").asString())
								.setCreatorEmail(r.get(0).get("creatorEmail").asString())
								.setCountry(r.get(0).get("country").asString())
								.setDescription(r.get(0).get("description").asString())
								.setLocation(Geolocation.newBuilder().setLat(r.get(0).get("lat").asFloat())
										.setLon(r.get(0).get("lon").asFloat()).build())
								.setCityId(r.get(1).asInt()).build());
					}
					return true;
				}
			});
			return citys;

		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * Update user
	 * 
	 * @param request user data
	 * @return true if success
	 */

	public boolean updateUser(User request) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {
				@Override
				public String execute(Transaction tx) {
					Result result = tx.run(
							"Match (a:User) " + "WHERE a.email = $email " + " SET a.name = $name "
									+ " SET a.description = $description " + "RETURN a",
							parameters("name", request.getName().toLowerCase(), "email", request.getEmail().toLowerCase(), "description",
									request.getDescripiton()));
					return result.toString();
				}
			});
			return true;
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
	}

	/**
	 * Update a city
	 * 
	 * @param request city data
	 * @return true if success
	 */

	public boolean updateCity(City request) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {

				@Override
				public String execute(Transaction tx) {
					Result result = tx.run(
							"MATCH (a:City) " + "WHERE a.name =  $name " + " AND a.country = $country "
									+ "AND   a.creatorEmail = $creatorEmail "
									+ "SET  a.description = $description, a.lat = $lat , a.lon = $lon "

									+ "RETURN a.name + ', from node ' + id(a)",
							parameters("name", request.getName().toLowerCase(), "creatorEmail", request.getCreatorEmail().toLowerCase(),
									"description", request.getDescription(), "lat", request.getLocation().getLat(),
									"lon", request.getLocation().getLon(), "country", request.getCountry()));
					return result.single().get(0).asString();
				}
			});
			return true;
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
	}

	/**
	 * Update a place
	 * 
	 * @param request place data
	 * @return true if sucess
	 */

	public boolean updatePlace(Place request) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {

				@Override
				public String execute(Transaction tx) {
					Result result = tx.run(
							"MATCH (a:Place) " + "WHERE a.name =  $name " + " AND a.country = $country "
									+ "AND   a.city = $city "
									+ "SET  a.description = $description, a.lat = $lat , a.lon = $lon "
									+ "RETURN a.name + ', from node ' + id(a)",
							parameters("name", request.getName().toLowerCase(), "city", request.getCity().toLowerCase(), "description",
									request.getDescription(), "lat", request.getLocation().getLat(), "lon",
									request.getLocation().getLon(), "country", request.getCountry()));
					return result.single().get(0).asString();
				}
			});
			return true;
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
	}

	/**
	 * Returns all places from a city
	 * 
	 * @param request city request
	 * @return list with places of the city
	 */
	public List<Place> getPlacesInCity(CityRequestPDB request) {
		List<Place> places = new ArrayList<>();
		try (Session session = driver.session()) {
			boolean result = session.writeTransaction(new TransactionWork<Boolean>() {

				@Override
				public Boolean execute(Transaction tx) {
					Result result = tx.run(
							"Match (a:City) " + " -[r:ISIN]-> " + " (b:Place) "
									+ "WHERE a.name= $name AND a.country = $country " + "RETURN b,id(b)",
							parameters("name", request.getName().toLowerCase(), "country", request.getCountry().toLowerCase())); //
				

					for (Record r : result.list()) {
						places.add(Place.newBuilder().setName(r.get(0).get("name").asString())
								.setCity(r.get(0).get("city").asString()).setCity(r.get(0).get("country").asString())
								.setCreatorEmail(r.get(0).get("creatorEmail").asString())
								.setDescription(r.get(0).get("description").asString())
								.setLocation(Geolocation.newBuilder().setLat(r.get(0).get("lat").asFloat())
										.setLon(r.get(0).get("lon").asFloat()).build())
								.setPlaceId(r.get(1).asInt()).build());
					}
					return true;
				}
			});
			return places;

		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	public static void main(String... args) throws Exception {
		try (DAO dao = new DAO("bolt://172.17.0.1:7687", "neo4j", "test")) {

			User user = User.newBuilder().setName("TestUser4").setEmail("testUser4@email.com").setDescripiton("new  description").build();
			dao.updateUser(user);
			
			//City temp = dao.getCity("san Pedro", "Chile");
			//System.out.println(temp);
//			 City city =
//			 City.newBuilder().setName("Dublin").setCountry("Ireland").setDescription("A city in Ireland.").setCreatorEmail("testUser4@email.com").build();
//			 dao.createCity(city);

//			 User user =
//			 User.newBuilder().setName("TestUser4").setEmail("testUser4@email.com").setDescripiton("Test user description").build();
//			 dao.createUser(user);

			// System.out.println(dao.visitCity(VisitCityRequestPDB.newBuilder().setCityId(5).setUserEmail("testUser@email.com").build()));

			// System.out.println(dao.getVisitedCitys("email"));

			/*
			 * BlockingQueue<CityOrBuilder> queue = new
			 * ArrayBlockingQueue<CityOrBuilder>(1000); dao.getAllCitys(null,queue);
			 * while(true) { CityOrBuilder temp = queue.take(); if(temp instanceof
			 * CityPoison) break;
			 * 
			 * System.out.println((City)temp); }
			 */
			
//			 dao.createPlace(Place.newBuilder().setCity("Dublin").setCountry("Ireland")
//					 .setName("Congress").setDescription("A Place in Santiago").setCreatorEmail("testuser2@email.com").build());
//			
			
			//dao.visitPlace(VisitPlaceRequestPDB.newBuilder().setPlaceId(16).setUserEmail("testuser4@email.com").build());

			// System.out.println(dao.getVisitedPlaces("email"));

			// System.out.print(dao.createUser(User.newBuilder().setName("name1").setDescripiton("descript1")
			// .setEmail("m1m").build()));
			// User u = dao.getUser("mm");
			// System.out.println(u);
			/*
			 * City t = new City(); t.setCreatorEmail("creator"); t.setName("name");
			 * t.setDescription("description"); t.setCountry("country");
			 * t.setPicture("picture"); t.setGeolocation(new Geolocation(3,3));
			 * dao.createCity(t);
			 */
			// System.out.println(dao.getCity("galway", "ireland"));

			// System.out.println(
			// dao.createPlace(Place.newBuilder().setName("place2").setCity("galway2").setCountry(
			// "ireland").setCreatorEmail("user1@email.com") .setDescription("description")
			// .setLocation(Geolocation.newBuilder().setLat(1).setLon(1).build()).build()));

			// System.out.println(dao.getPlace("place2", "galway2", "ireland"));

			// String l = dao.visitPlace("user1@email.com", "place2", "galway2", "ireland");
			// CityResponsePDBvisiSystem.out.println(l);
			// for(Record r : l) { System.out.println(r.get(0)); }

			// dao.getVisitedPlaces("user1@email.com");
			// dao.visitCity(VisitCityRequestPDB.newBuilder().setEmail("user1@email.com")
			// .setCityName("galway")
			// .setCityCountry("ireland").build());
			// dao.getVisitedPlaces("user1@email.com");
			// System.out.println(dao.updateUser(CreateUserRequestPDB.newBuilder().setEmail("user1@email.com")
			// .setName("nnnn").setDescription("nnn").build()));
			// dao.updateCity(CityPDB.newBuilder().setName("galway").setCountry("ireland")
			// .setCreatorEmail("user1@email.com").setDescription("xxxx")
			// .setLocation(GeolocationPDB.newBuilder().setLat(7).setLon(7)).build());
			// dao.updatePlace(PlacePDB.newBuilder().setName("gmit")
			// .setCity("galway")
			// .setCountry("ireland")
			// .setDescription("new Description xxx")
			// .setLocation(GeolocationPDB.newBuilder().setLat(7).setLon(5).build()).build());
			/*
			 * List<PlacePDB> res = dao.getPlacesInCity(CityRequestPDB.newBuilder()
			 * .setName("galway").setCountry("ireland").build()); for ( PlacePDB p : res ) {
			 * System.out.println(p.getName()); }
			 */
		}
	}

	@Override
	public void close() throws Exception {
		driver.close();

	}
}
