package ie.gmit.wcity.profile;

import java.util.ArrayList;

public class City {
	private String name;
	private String country;
	private String creatorEmail;
	private Geolocation geolocation;
	private String description;
	private ArrayList<Place> places;

	public City() {
	}

	public City(String name, String country, String creatorEmail, Geolocation geolocation, String description,
			ArrayList<Place> places) {
		super();
		this.name = name;
		this.country = country;
		this.creatorEmail = creatorEmail;
		this.geolocation = geolocation;
		this.description = description;
		this.places = places;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getCountry() {
		return country;
	}

	public void setCountry(String country) {
		this.country = country;
	}

	public String getCreatorEmail() {
		return creatorEmail;
	}

	public void setCreatorEmail(String creatorEmail) {
		this.creatorEmail = creatorEmail;
	}

	public Geolocation getGeolocation() {
		return geolocation;
	}

	public void setGeolocation(Geolocation geolocation) {
		this.geolocation = geolocation;
	}

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
	}

	public ArrayList<Place> getPlaces() {
		return places;
	}

	public void setPlaces(ArrayList<Place> places) {
		this.places = places;
	}

}
