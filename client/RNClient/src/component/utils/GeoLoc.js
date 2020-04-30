import React, {Component} from 'react';
import {Alert, StyleSheet, Text, View, TouchableOpacity} from 'react-native';
import Geolocation from '@react-native-community/geolocation';
import Geocoder from 'react-native-geocoding';
import {constants} from '../../constants/Constants'

Geocoder.init(constants.apiKey, {language: "en"});

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
        this.props.parentCallback(this.state.latitude, this.state.longitude, this.state.city, this.state.country)
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
        if (this.state.error) {
            this.state.city = this.state.error;
        }
        return (
            <View>
            </View>
        )
    }
}
