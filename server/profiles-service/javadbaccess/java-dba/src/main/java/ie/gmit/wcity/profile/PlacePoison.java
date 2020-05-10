package ie.gmit.wcity.profile;

import java.util.List;
import java.util.Map;

import com.google.protobuf.ByteString;
import com.google.protobuf.Descriptors.Descriptor;
import com.google.protobuf.Descriptors.FieldDescriptor;
import com.google.protobuf.Descriptors.OneofDescriptor;
import com.google.protobuf.Message;
import com.google.protobuf.UnknownFieldSet;

import io.grpc.wcity.profiles.Geolocation;
import io.grpc.wcity.profiles.GeolocationOrBuilder;
import io.grpc.wcity.profiles.PlaceOrBuilder;

public class PlacePoison implements PlaceOrBuilder{

	@Override
	public Message getDefaultInstanceForType() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public List<String> findInitializationErrors() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public String getInitializationErrorString() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public Descriptor getDescriptorForType() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public Map<FieldDescriptor, Object> getAllFields() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public boolean hasOneof(OneofDescriptor oneof) {
		// TODO Auto-generated method stub
		return false;
	}

	@Override
	public FieldDescriptor getOneofFieldDescriptor(OneofDescriptor oneof) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public boolean hasField(FieldDescriptor field) {
		// TODO Auto-generated method stub
		return false;
	}

	@Override
	public Object getField(FieldDescriptor field) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public int getRepeatedFieldCount(FieldDescriptor field) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public Object getRepeatedField(FieldDescriptor field, int index) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public UnknownFieldSet getUnknownFields() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public boolean isInitialized() {
		// TODO Auto-generated method stub
		return false;
	}

	@Override
	public String getName() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public ByteString getNameBytes() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public String getCity() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public ByteString getCityBytes() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public String getCountry() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public ByteString getCountryBytes() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public String getCreatorEmail() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public ByteString getCreatorEmailBytes() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public boolean hasLocation() {
		// TODO Auto-generated method stub
		return false;
	}

	@Override
	public Geolocation getLocation() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public GeolocationOrBuilder getLocationOrBuilder() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public String getDescription() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public ByteString getDescriptionBytes() {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public int getPlaceId() {
		// TODO Auto-generated method stub
		return 0;
	}

}
