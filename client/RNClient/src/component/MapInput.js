import React from 'react';
import {GooglePlacesAutocomplete} from 'react-native-google-places-autocomplete';

export default class MapInput extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            lat: 0,
            lng: 0,
            city: '',
            country: '',
            flag: 'false',
        }
    }
    sendData = () => {
        this.props.parentCallback(this.state.lat, this.state.lng, this.state.city, this.state.country)
    }
    render() {
        return (
            <GooglePlacesAutocomplete
                placeholder='Search'
                minLength={2} // minimum length of text to search
                autoFocus={false}
                fetchDetails={true}
                returnKeyType={'search'} // Can be left out for default return key https://facebook.github.io/react-native/docs/textinput.html#returnkeytype
                keyboardAppearance={'light'} // Can be left out for default keyboardAppearance https://facebook.github.io/react-native/docs/textinput.html#keyboardappearance
                listViewDisplayed='auto'    // true/false/undefined
                fetchDetails={true}
                renderDescription={row => row.description} // custom description render
                onPress={(data, details = null) => { // 'details' is provided when fetchDetails = true
                    //console.log(data, details);
                    this.setState({
                        lat: details.geometry.location.lat,
                        lng: details.geometry.location.lng,
                        city: details.address_components.filter(ac => ~ac.types.indexOf('locality'))[0].long_name,
                        country: details.address_components.filter(ac => ~ac.types.indexOf('country'))[0].long_name,
                        flag: true,
                    })

                    console.log("lat&lng " + this.state.lat + " " + this.state.lng, this.state.city, this.state.country)
                    this.sendData()

                    // this.props.notifyChange(details.geometry.location)
                }}
                query={{
                    // available options: https://developers.google.com/places/web-service/autocomplete
                    key: 'AIzaSyAB7wbhzqnnlXl3k1ugJgQ_tjS3Ks6Jycc',
                    language: 'en', // language of the results
                    types: '(cities)' // default: 'geocode'
                }}
                nearbyPlacesAPI='GooglePlacesSearch' // Which API to use: GoogleReverseGeocoding or GooglePlacesSearch
                debounce={200} // debounce the requests in ms. Set to 0 to remove debounce. By default 0ms.
            >

              </GooglePlacesAutocomplete>


        )
    }
}

