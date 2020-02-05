package ie.gmit.wcity.profile;

import static org.neo4j.driver.Values.parameters;

import org.neo4j.driver.AuthTokens;
import org.neo4j.driver.Driver;
import org.neo4j.driver.GraphDatabase;
import org.neo4j.driver.Record;
import org.neo4j.driver.Result;
import org.neo4j.driver.Session;
import org.neo4j.driver.Transaction;
import org.neo4j.driver.TransactionWork;

import profiles.Test;

public class DAO implements AutoCloseable {

	private final Driver driver;

	public DAO(String uri, String user, String password) {
		driver = GraphDatabase.driver(uri, AuthTokens.basic(user, password));
	}

	public String createUser(final String email, final String name, final String description) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {
				@Override
				public String execute(Transaction tx) {
					Result result = tx.run(
							"Create (a:User) " + "SET a.name = $name " + "SET a.email = $email "
									+ "SET a.description = $description " + "RETURN a.name + ', from node ' + id(a)",
							parameters("name", name, "email", email, "description", description));
					return result.single().get(0).asString();
				}
			});
			return result;

		}
	}

	public User getUser(final String email) {
		User user = null;

		try (Session session = driver.session()) {
			user = session.writeTransaction(new TransactionWork<User>() {
				@Override
				public User execute(Transaction tx) {
					Result result = tx.run("MATCH (a:User) " + "where a.email = $email " + "RETURN a",
							parameters("email", email));
					User u = new User();
					u.setEmail(email);
				
					Record r = result.next();
					u.setName(r.get(0).get("name").asString());
					//u.setDescription(result.single().get(0).get("name").asString());
					
					u.setDescription(r.get(0).get("description").asString());
					return u;
				}
			});
		}
		return user;
	}
	
	public String createCity(City city) {
		try (Session session = driver.session()) {
			String result = session.writeTransaction(new TransactionWork<String>() {
				@Override
				public String execute(Transaction tx) {
					Result result = tx.run(
							"Create (a:City) " + "SET a.name = $name " +
					"SET a.creatorEmail = $creatorEmail "
									+ "SET a.description = $description " + 
									"SET a.lat = $lat " +
									"SET a.lon = $lon " + 
									"SET a.country = $country " + 
									"SET a.picture = $picture " + 
									"RETURN a.name + ', from node ' + id(a)",
							parameters("name", city.getName(), 
									"creatorEmail", city.getCreatorEmail(),
									"description", city.getDescription(),
									"lat",city.getGeolocation().getLat(),
									"lon",city.getGeolocation().getLon(),
									"country",city.getCountry(),
									"picture",city.getPicture()
									));
					return result.single().get(0).asString();
				}
			});
			return result;

		}
	}

	public static void main(String... args) throws Exception {
		try (DAO dao = new DAO("bolt://192.168.43.58:7687", "neo4j", "test")) {
			//dao.AddUser("email1", "name1", "description1");
			//User u = dao.getUser("one");
			//System.out.println(u);
			City t = new City();
			t.setCreatorEmail("creator");
			t.setName("name");
			t.setDescription("description");
			t.setCountry("country");
			t.setPicture("picture");
			t.setGeolocation(new Geolocation(3,3));
			dao.createCity(t);
		}
	}

	@Override
	public void close() throws Exception {
		driver.close();

	}
}
//CREATE (jaedcom:User {email:'j@e.com',name: 'John', description : ' max 160 characters', picture: 'the picture address'})