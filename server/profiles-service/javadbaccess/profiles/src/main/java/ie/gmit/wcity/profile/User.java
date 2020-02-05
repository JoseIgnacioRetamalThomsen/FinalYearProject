package ie.gmit.wcity.profile;

import java.util.ArrayList;

public class User {

	private String email;
	private String name;
	private String description;
	
	private ArrayList<City> citys;
	private ArrayList<Place> places;
	
	public User() {}

	public User(String email, String name, String description, ArrayList<City> citys, ArrayList<Place> places) {
		super();
		this.email = email;
		this.name = name;
		this.description = description;
		this.citys = citys;
		this.places = places;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
	}

	public ArrayList<City> getCitys() {
		return citys;
	}

	public void setCitys(ArrayList<City> citys) {
		this.citys = citys;
	}

	public ArrayList<Place> getPlaces() {
		return places;
	}

	public void setPlaces(ArrayList<Place> places) {
		this.places = places;
	}
	
}
