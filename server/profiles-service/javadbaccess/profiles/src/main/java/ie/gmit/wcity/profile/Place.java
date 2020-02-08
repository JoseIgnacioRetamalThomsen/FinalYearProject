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

	public Place setDescription(String description) {
		Description = description;
		return this;
	}

	public String getName() {
		return name;
	}

	public Place setName(String name) {
		this.name = name;
		return this;
	}

	public String getCityName() {
		return cityName;
	}

	public Place setCityName(String cityName) {
		this.cityName = cityName;
		return this;
	}

	public String getCityCountry() {
		return cityCountry;
	}

	public Place setCityCountry(String cityCountry) {
		this.cityCountry = cityCountry;
		return this;
	}

	public String getCreatorEmail() {
		return creatorEmail;
	}

	public Place setCreatorEmail(String creatorEmail) {
		this.creatorEmail = creatorEmail;
		return this;
	}

	public Geolocation getGeolocation() {
		return geolocation;
	}

	public Place setGeolocation(Geolocation geolocation) {
		this.geolocation = geolocation;
		return this;
	}

	@Override
	public String toString() {
		return "Place [name=" + name + ", cityName=" + cityName + ", cityCountry=" + cityCountry + ", creatorEmail="
				+ creatorEmail + ", Description=" + Description + ", geolocation=" + geolocation + "]";
	}

	
}
