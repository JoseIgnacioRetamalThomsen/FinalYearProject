import React, {Component} from 'react';
import {Alert, StyleSheet, Text, View, TouchableOpacity} from 'react-native';
import Geolocation from '@react-native-community/geolocation';
import Geocoder from 'react-native-geocoding';

Geocoder.init("AIzaSyAB7wbhzqnnlXl3k1ugJgQ_tjS3Ks6Jycc", {language: "en"});

export default class GeoLoc extends Component {
    constructor(props) {
        super(props);
        this.state = {
            latitude: null,
            longitude: null,
            error: null,
            city: '',
            country: '',
        }
    }
    sendData = () => {
        this.props.parentCallback(this.state.city, this.state.country, this.state.latitude, this.state.longitude)
    }

    componentDidMount() {
        Geolocation.getCurrentPosition(
            (position) => {
                this.setState({
                    latitude: position.coords.latitude,
                    longitude: position.coords.longitude,
                });
                Geocoder.from(position.coords.latitude, position.coords.longitude)
                    .then(json => {
                       // console.log(json);
                        // json.results[0].address_components[1].long_name;//address_components[0] GMIT
                        //var user_city = results[0].address_components.filter(ac=>~ac.types.indexOf('locality'))[0].long_name
                        let city = json.results[0].address_components.filter(ac => ~ac.types.indexOf('locality'))[0].long_name
                        let country = json.results[0].address_components.filter(ac => ~ac.types.indexOf('country'))[0].long_name
                        this.setState({
                            city: city
                        })
                        this.setState({
                            country: country
                        })
                        this.sendData()
                       //console.log("!!!!", city, country)
                    })
                    .catch(error => console.warn(error));
            },
            (error) => {
                // See error code charts below.
                this.setState({
                    error: error.message
                }),
                    console.log(error.code, error.message);
            },

            {
                enableHighAccuracy: false,
                timeout: 10000,
                maximumAge: 100000
            }
        );
    }

    render() {
        let text = '';
        if (this.state.error) {
            text = this.state.error;
        } else if (this.state.city) {
            text = this.state.city
        }
        return (
            <View>
            </View>
        )
    }
}
