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

	public String AddUser(final String email, final String name, final String description) {
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

	public User GetUser(final String email) {
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

	public static void main(String... args) throws Exception {
		try (DAO dao = new DAO("bolt://192.168.43.58:7687", "neo4j", "test")) {
			//dao.AddUser("email1", "name1", "description1");
			User u = dao.GetUser("one");
			System.out.println(u);
		}
	}

	@Override
	public void close() throws Exception {
		driver.close();

	}
}
//CREATE (jaedcom:User {email:'j@e.com',name: 'John', description : ' max 160 characters', picture: 'the picture address'})