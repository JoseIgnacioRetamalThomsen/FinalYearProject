/**
 * Jose I. Retamal
 * GMIT - 2020
 */

package ie.gmit.wcity.profile;

import static org.neo4j.driver.Values.parameters;

import java.util.ArrayList;
import java.util.List;

import org.neo4j.driver.AuthTokens;
import org.neo4j.driver.Driver;
import org.neo4j.driver.GraphDatabase;
import org.neo4j.driver.Record;
import org.neo4j.driver.Result;
import org.neo4j.driver.Session;
import org.neo4j.driver.Transaction;
import org.neo4j.driver.TransactionWork;

import io.grpc.wcity.profilesDB.CityPDB;
import io.grpc.wcity.profilesDB.CityRequestPDB;
import io.grpc.wcity.profilesDB.CreateUserRequestPDB;
import io.grpc.wcity.profilesDB.GeolocationPDB;
import io.grpc.wcity.profilesDB.PlacePDB;
import io.grpc.wcity.profilesDB.VisitCityRequestPDB;
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
   * Create a new user
   * 
   * @param email       user email
   * @param name        user name
   * @param description user description
   * @return true if is created
   */
  public String createUser(final String email, final String name,
      final String description) {
    try (Session session = driver.session()) {
      String result = session.writeTransaction(new TransactionWork<String>() {
        @Override
        public String execute(Transaction tx) {
          Result result = tx.run("Create (a:User) " + "SET a.name = $name "
              + "SET a.email = $email " + "SET a.description = $description "
              + "RETURN a.name + ', from node ' + id(a)",
              parameters("name", name, "email", email, "description",
                  description));
          return result.single().get(0).asString();
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
          Result result = tx.run(
              "MATCH (a:User) " + "where a.email = $email " + "RETURN a",
              parameters("email", email));
          User u = new User();
          u.setEmail(email);

          Record r = result.next();
          u.setName(r.get(0).get("name").asString());
          u.setDescription(r.get(0).get("description").asString());
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
  public String createCity(City city) {
    try (Session session = driver.session()) {
      String result = session.writeTransaction(new TransactionWork<String>() {
        @Override
        public String execute(Transaction tx) {
          Result result = tx.run(
              "Create (a:City) " + "SET a.name = $name "
                  + "SET a.creatorEmail = $creatorEmail "
                  + "SET a.description = $description " + "SET a.lat = $lat "
                  + "SET a.lon = $lon " + "SET a.country = $country "
                  + "SET a.picture = $picture "
                  + "RETURN a.name + ', from node ' + id(a)",
              parameters("name", city.getName(), "creatorEmail",
                  city.getCreatorEmail(), "description", city.getDescription(),
                  "lat", city.getGeolocation().getLat(), "lon",
                  city.getGeolocation().getLon(), "country", city.getCountry(),
                  "picture", city.getPicture()));

          return result.single().get(0).asString();
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
          Result result = tx.run(
              "MATCH (a:City) " + "where a.name = $name " + " AND "
                  + " a.country = $country " + "RETURN a",
              parameters("name", name, "country", country));
          Record r = result.next();
          City c = new City().setName(r.get(0).get("name").asString())
              .setCountry(r.get(0).get("country").asString())
              .setCreatorEmail(r.get(0).get("creatorEmail").asString())
              .setDescription(r.get(0).get("description").asString())
              .setGeolocation(new Geolocation(r.get(0).get("lat").asFloat(),
                  r.get(0).get("lon").asFloat()));

          return c;
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
  public String createPlace(Place place) {
    try (Session session = driver.session()) {
      String result = session.writeTransaction(new TransactionWork<String>() {
        @Override
        public String execute(Transaction tx) {
          Result result = tx.run(
              "Create (a:Place) " + "SET a.name = $name "
                  + "SET a.creatorEmail = $creatorEmail "
                  + "SET a.description = $description " + "SET a.lat = $lat "
                  + "SET a.lon = $lon " + "SET a.country = $country "
                  + "SET a.city = $city "
                  + "RETURN a.name + ', from node ' + id(a)",
              parameters("name", place.getName(), "creatorEmail",
                  place.getCreatorEmail(), "description",
                  place.getDescription(), "lat",
                  place.getGeolocation().getLat(), "lon",
                  place.getGeolocation().getLon(), "country",
                  place.getCityCountry(), "city", place.getCityName()));
          tx.run("MATCH (a:City),(b:Place) "
              + "WHERE a.name = $cityName AND  a.country = $country "
              + "AND b.name = $placeNane AND b.country =$country AND b.city = $cityName "
              + " CREATE (a)-[r:ISIN]->(b)" + "RETURN r",
              parameters("cityName", place.getCityName(), "country",
                  place.getCityCountry(), "placeNane", place.getName(), "city",
                  place.getCityName()));

          return result.single().get(0).asString();
        }
      });
      return result;

    }catch(Exception e) {
      e.printStackTrace();
      return null;
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
              "MATCH (a:Place) " + "where a.name = $name " + " AND "
                  + " a.country = $country " + " AND" + " a.city = $city "
                  + "RETURN a",
              parameters("name", name, "country", country, "city", city));
          Record r = result.next();
          Place p = new Place().setName(r.get(0).get("name").asString())
              .setCityCountry(r.get(0).get("country").asString())
              .setCityName(r.get(0).get("city").asString())
              .setCreatorEmail(r.get(0).get("creatorEmail").asString())
              .setDescription(r.get(0).get("description").asString())
              .setGeolocation(new Geolocation(r.get(0).get("lat").asFloat(),
                  r.get(0).get("lon").asFloat()));
          return p;
        }
      });
    } catch (Exception e) {
      e.printStackTrace();
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
  public List<Record> visitPlace(String userEmail, String placeName,
      String placeCity, String placeCountry) {
    try (Session session = driver.session()) {
      List<Record> result = session
          .writeTransaction(new TransactionWork<List<Record>>() {
            @Override
            public List<Record> execute(Transaction tx) {
              Result result = tx.run(
                  "Match (a:User) " + "WHERE a.email = $email "
                      + "MATCH (b:Place) " + "WHERE b.name = $name " + "AND "
                      + "b.city =$city " + " AND b.country = $country "
                      + "CREATE (a)-[r:VISIT{date:date()}]->(b)"
                      + "RETURN r.date",
                  parameters("email", userEmail, "name", placeName, "city",
                      placeCity, "country", placeCountry));
              // return result.single().get(0).asString();
              return result.list();
            }
          });
      return result;
    } catch (Exception e) {
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
  public List<PlacePDB> getVisitedPlaces(String userEmail) {
    List<PlacePDB> places = new ArrayList<>();
    try (Session session = driver.session()) {
      boolean result = session.writeTransaction(new TransactionWork<Boolean>() {
        @Override
        public Boolean execute(Transaction tx) {
          Result result = tx.run("Match (a:User) " + " -[r:VISIT]-> "
              + " (b:Place) " + "WHERE a.email = $email " +

              "RETURN b", parameters("email", userEmail));
          // return result.single().get(0).asString();
          for (
            Record r : result.list()
          ) {
            places.add(PlacePDB.newBuilder()
                .setName(r.get(0).get("name").asString())
                .setCity(r.get(0).get("city").asString())
                .setCity(r.get(0).get("country").asString())
                .setCreatorEmail(r.get(0).get("creatorEmail").asString())
                .setDescription(r.get(0).get("description").asString())
                .setLocation(GeolocationPDB.newBuilder()
                    .setLat(r.get(0).get("lat").asFloat())
                    .setLon(r.get(0).get("lon").asFloat()).build())
                .build());
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
      List<Record> result = session
          .writeTransaction(new TransactionWork<List<Record>>() {
            @Override
            public List<Record> execute(Transaction tx) {
              Result result = tx.run("Match (a:User) "
                  + "WHERE a.email = $email " + "MATCH (b:City) "
                  + "WHERE b.name = $name " + "AND " + "b.country = $country "
                  + "CREATE (a)-[r:VISIT{date:date()}]->(b)" + "RETURN r.date",
                  parameters("email", request.getEmail(), "name",
                      request.getCityName(), "country",
                      request.getCityCountry()));

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
  public List<CityPDB> getVisitedCitys(String userEmail) {
    List<CityPDB> citys = new ArrayList<>();
    try (Session session = driver.session()) {
      boolean result = session.writeTransaction(new TransactionWork<Boolean>() {
        @Override
        public Boolean execute(Transaction tx) {
          Result result = tx.run("Match (a:User) " + " -[r:VISIT]-> "
              + " (b:City) " + "WHERE a.email = $email " +

              "RETURN b", parameters("email", userEmail));
          // return result.single().get(0).asString();
          for (
            Record r : result.list()
          ) {
            citys.add(CityPDB.newBuilder()
                .setName(r.get(0).get("name").asString())
                .setCreatorEmail(r.get(0).get("creatorEmail").asString())
                .setDescription(r.get(0).get("description").asString())
                .setLocation(GeolocationPDB.newBuilder()
                    .setLat(r.get(0).get("lat").asFloat())
                    .setLon(r.get(0).get("lon").asFloat()).build())
                .build());

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
  public boolean updateUser(CreateUserRequestPDB request) {
    try (Session session = driver.session()) {
      String result = session.writeTransaction(new TransactionWork<String>() {
        @Override
        public String execute(Transaction tx) {
          Result result = tx.run(
              "Match (a:User) " + "WHERE a.email = $email "
                  + " SET a.name = $name "
                  + " SET a.description = $description " + "RETURN a",
              parameters("name", request.getName(), "email", request.getEmail(),
                  "description", request.getDescription()));
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
  public boolean updateCity(CityPDB request) {
    try (Session session = driver.session()) {
      String result = session.writeTransaction(new TransactionWork<String>() {
        @Override
        public String execute(Transaction tx) {
          Result result = tx.run("MATCH (a:City) " + "WHERE a.name =  $name "
              + " AND a.country = $country "
              + "AND   a.creatorEmail = $creatorEmail "
              + "SET  a.description = $description, a.lat = $lat , a.lon = $lon "

              + "RETURN a.name + ', from node ' + id(a)",
              parameters("name", request.getName(), "creatorEmail",
                  request.getCreatorEmail(), "description",
                  request.getDescription(), "lat",
                  request.getLocation().getLat(), "lon",
                  request.getLocation().getLon(), "country",
                  request.getCountry()));
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
  public boolean updatePlace(PlacePDB request) {
    try (Session session = driver.session()) {
      String result = session.writeTransaction(new TransactionWork<String>() {
        @Override
        public String execute(Transaction tx) {
          Result result = tx.run("MATCH (a:Place) " + "WHERE a.name =  $name "
              + " AND a.country = $country " + "AND   a.city = $city "
              + "SET  a.description = $description, a.lat = $lat , a.lon = $lon "
              + "RETURN a.name + ', from node ' + id(a)",
              parameters("name", request.getName(), "city", request.getCity(),
                  "description", request.getDescription(), "lat",
                  request.getLocation().getLat(), "lon",
                  request.getLocation().getLon(), "country",
                  request.getCountry()));
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
  public List<PlacePDB> getPlacesInCity(CityRequestPDB request) {
    List<PlacePDB> places = new ArrayList<>();
    try (Session session = driver.session()) {
      boolean result = session.writeTransaction(new TransactionWork<Boolean>() {
        @Override
        public Boolean execute(Transaction tx) {
          Result result = tx.run(
              "Match (a:City) " + " -[r:ISIN]-> " + " (b:Place) "
                  + "WHERE a.name= $name AND a.country = $country "
                  + "RETURN b",
              parameters("name", request.getName(), "country",
                  request.getCountry()));
          // return result.single().get(0).asString();
          for (
            Record r : result.list()
          ) {
            places.add(PlacePDB.newBuilder()
                .setName(r.get(0).get("name").asString())
                .setCity(r.get(0).get("city").asString())
                .setCity(r.get(0).get("country").asString())
                .setCreatorEmail(r.get(0).get("creatorEmail").asString())
                .setDescription(r.get(0).get("description").asString())
                .setLocation(GeolocationPDB.newBuilder()
                    .setLat(r.get(0).get("lat").asFloat())
                    .setLon(r.get(0).get("lon").asFloat()).build())
                .build());
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
    try (DAO dao = new DAO("bolt://192.168.43.58:7687", "neo4j", "test")) {
      // dao.AddUser("email1", "name1", "description1");
      // User u = dao.getUser("one");
      // System.out.println(u);
      /*
       * City t = new City(); t.setCreatorEmail("creator"); t.setName("name");
       * t.setDescription("description"); t.setCountry("country");
       * t.setPicture("picture"); t.setGeolocation(new Geolocation(3,3));
       * dao.createCity(t);
       */
      // System.out.println(dao.getCity("galway", "ireland"));

      /*
       * Place p = new
       * Place().setName("place1").setCityName("galway").setCityCountry(
       * "ireland").setCreatorEmail("user1@email.com")
       * .setDescription("description") .setGeolocation(new Geolocation(4,5));
       * dao.createPlace(p);
       */

      // System.out.println(dao.getPlace("gmit","galway", "ireland"));
      /*
       * List<Record> l = dao.visitPlace("user1@email.com", "gmit", "galway",
       * "ireland"); for(Record r : l) { System.out.println(r.get(0)); }
       */
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

      List<PlacePDB> res = dao.getPlacesInCity(CityRequestPDB.newBuilder()
          .setName("galway").setCountry("ireland").build());
      for (
        PlacePDB p : res
      ) {
        System.out.println(p.getName());
      }
    }
  }

  @Override
  public void close() throws Exception {
    driver.close();

  }
}
