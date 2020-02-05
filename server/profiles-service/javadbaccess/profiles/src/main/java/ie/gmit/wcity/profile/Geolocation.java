package ie.gmit.wcity.profile;

public class Geolocation {
	private float lon;
	private float lat;

	public Geolocation(float lon, float lat) {
		super();
		this.lon = lon;
		this.lat = lat;
	}

	public float getLon() {
		return lon;
	}

	public void setLon(float lon) {
		this.lon = lon;
	}

	public float getLat() {
		return lat;
	}

	public void setLat(float lat) {
		this.lat = lat;
	}

}
