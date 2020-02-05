package ie.gmit.wcity.profile;

import java.util.ArrayList;

public class City {
	private String name;
	private String country;
	private String creatorEmail;
	private Geolocation geolocation;
	private String description;
	private String picture;
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
	

	public City(String name, String country, String creatorEmail, Geolocation geolocation, String description,
			String picture, ArrayList<Place> places) {
		super();
		this.name = name;
		this.country = country;
		this.creatorEmail = creatorEmail;
		this.geolocation = geolocation;
		this.description = description;
		this.picture = picture;
		this.places = places;
	}

	public String getPicture() {
		return picture;
	}

	public City setPicture(String picture) {
		this.picture = picture;
		return this;
	}

	public String getName() {
		return name;
	}

	public City setName(String name) {
		this.name = name;
		return this;
	}

	public String getCountry() {
		return country;
	}

	public City setCountry(String country) {
		this.country = country;
		return this;
	}

	public String getCreatorEmail() {
		return creatorEmail;
	}

	public City setCreatorEmail(String creatorEmail) {
		this.creatorEmail = creatorEmail;
		return this;
	}

	public Geolocation getGeolocation() {
		return geolocation;
	}

	public City setGeolocation(Geolocation geolocation) {
		this.geolocation = geolocation;
		return this;
	}

	public String getDescription() {
		return description;
	}

	public City setDescription(String description) {
		this.description = description;
		return this;
	}

	public ArrayList<Place> getPlaces() {
		return places;
	}

	public City setPlaces(ArrayList<Place> places) {
		this.places = places;
		return this;
	}

}
