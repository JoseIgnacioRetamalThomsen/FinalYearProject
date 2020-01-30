package ie.gmit.wcity.profile;

import static org.neo4j.driver.Values.parameters;

import org.neo4j.driver.AuthTokens;
import org.neo4j.driver.Driver;
import org.neo4j.driver.GraphDatabase;
import org.neo4j.driver.Result;
import org.neo4j.driver.Session;
import org.neo4j.driver.Transaction;
import org.neo4j.driver.TransactionWork;

import profiles.Test;

public class DAO  implements AutoCloseable {

	
	private final Driver driver;
	
	public DAO(String uri, String user, String password) {
		driver = GraphDatabase.driver(uri, AuthTokens.basic(user, password));
	}
	
	public void AddUser(final String email, final String name, final String description) {
		try (Session session = driver.session()) {
			//String greeting =
					session.writeTransaction(new TransactionWork<String>() {
				@Override
				public String execute(Transaction tx) {
					Result result = tx.run("CREATE (a:User) " + "SET a.name = $name "  
							+ "SET a.email = $email " + "SET a.description = $description " 
							+ "RETURN a.name + ', from node ' + id(a)", parameters("name", name,"email",email,"description",description));
					return result.single().get(0).asString();
				}
			});
			
		}
	}

	public static void main(String... args) throws Exception {
	try (DAO dao = new DAO("bolt://0.0.0.0:7687", "neo4j", "test")) {
		dao.AddUser("email","name","description");
	}
}

	@Override
	public void close() throws Exception {
		driver.close();
		
	}	
}
//CREATE (jaedcom:User {email:'j@e.com',name: 'John', description : ' max 160 characters', picture: 'the picture address'})