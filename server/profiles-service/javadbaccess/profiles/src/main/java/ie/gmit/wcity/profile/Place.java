package ie.gmit.wcity.profile;

public class Place {
	private String name;
	private String cityName;
	private String cityCountry;
	private String creatorEmail;
	private String Description;
	private Geolocation geolocation;

	public Place() {
	}

	public Place(String name, String cityName, String cityCountry, String creatorEmail, String description,
			Geolocation geolocation) {
		super();
		this.name = name;
		this.cityName = cityName;
		this.cityCountry = cityCountry;
		this.creatorEmail = creatorEmail;
		Description = description;
		this.geolocation = geolocation;
	}

	public String getDescription() {
		return Description;
	}

	public void setDescription(String description) {
		Description = description;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getCityName() {
		return cityName;
	}

	public void setCityName(String cityName) {
		this.cityName = cityName;
	}

	public String getCityCountry() {
		return cityCountry;
	}

	public void setCityCountry(String cityCountry) {
		this.cityCountry = cityCountry;
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

}
