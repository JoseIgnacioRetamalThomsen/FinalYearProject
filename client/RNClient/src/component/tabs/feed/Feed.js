import React, {Component} from 'react';
import {View, TextArea} from 'react-native';
import {Button, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import City from '../../City'
import DisplayCity from '../../DisplayCity'
import MapInput from "../../MapInput";
import {GooglePlacesAutocomplete} from "react-native-google-places-autocomplete";
class Feed extends Component {
    constructor(props) {
        super(props)
        this.state = {
            lat: 0,
            lng: 0,
            city: '',
            country: '',
        }
    }
    callbackFunction = (lat, lng, city, country) => {
        this.setState({lat: lat})
        this.setState({lng:lng})
        this.setState({city:city})
        this.setState({country:country})
    }
    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="Feed" isHome={true} navigation={this.props.navigation}/>
                <MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()} parentCallback = {this.callbackFunction} />

                <Text> {this.state.lat.toFixed(2)} {this.state.lng.toFixed(2)} {this.state.city} {this.state.country}</Text>
            </View>

        )
    }
}

export default Feed
