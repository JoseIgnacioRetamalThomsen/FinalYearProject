import React, { Component } from 'react'
import {
    Keyboard,
    StyleSheet,
    Text,
    View,
    TouchableHighlight,
    TextInput
} from "react-native";
import MapView  from 'react-native-maps'
import _ from 'lodash'
import Geolocation from '@react-native-community/geolocation';
import{constants} from '../../constants/Constants'

export default class API extends Component {

    constructor(props) {
        super(props);
        this.state = {
            error:"",
            latitude: 0,
            longitude: 0,
            destination:"",
            predictions: []
        }
        this.onChangeDestinationDebounced = _.debounce(this.onChangeDestination)
    }
    componentDidMount() {
        Geolocation.getCurrentPosition(
            position => {
                this.setState({
                    latitude: position.coords.latitude,
                    longitude: position.coords.longitude
                })
            },
            error => this.setState({
                error: error.message
            }),
            {enableHighAccuracy: true, maximumAge: 2000, timeout: 2000}
        )
    }
    async onChangeDestination(destination){
        this.setState({destination})
        const apiUrl = constants.apiUri +
            '&input=${destination}&location=${this.state.latitude}, ${this.state.longitude}' +
            '&radius=2000';
        try{
            const result = await fetch(apiUrl);
            const json = await result.json();
            console.log(json)
            this.setState({})
            predictions: json.predictions
        }catch(err){
            console.error(err)
        }
    }
    pressedPrediction(prediction){
        console.log(prediction)
        Keyboard.dismiss();
        this.setState({
            predictions: [],
            destination: prediction.description
        })
        Keyboard
    }
    render() {
        const predictions = this.state.predictions.map(
            prediction =>(
                <TouchableHighlight
                key={prediction.id}
                onPress={() => this.pressedPrediction(prediction)}
                >
                    <Text style={styles.suggestions} key = {prediction.id}>
                       {prediction.description}
                    </Text>)
                </TouchableHighlight>
                )
        )
    return(
        <View style={styles.container}>
            <MapView
            region={{
                latitude: this.state.latitude,
                longitude: this.state.longitude,
                latitudeDelta: 0.015,
                longitudeDelta: 0.0121
            }}
            showsUserLocation = {true}
            />
            <TextInput
                placeholder="Enter destination"
                onChangeText={destination => this.onChangeDestinationDebounced(destination)}
                value={this.state.destination}
            />
            {predictions}
        </View>
    )
}
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        marginTop: 50,
        paddingLeft: 30,
        paddingRight: 30,
        marginBottom: 30
    },
    textInput:{
        height: 40,
        width:300,
        borderWidth: 0.5,
        marginTop: 50,
        marginLeft: 5,
        marginRight: 5,
        padding: 30,
        backgroundColor: 'white'
    },
    destinationInput:{
        height: 40,
        borderWidth: 0.5,
        marginTop: 50,
        marginLeft: 5,
        marginRight: 5,
        padding: 30,
        backgroundColor: 'white'
    },

    suggestions:{
        backgroundColor: 'black',
        padding: 5,
        fontSize: 18,
        borderWidth: 0.5,
        marginLeft: 5,
        marginRight: 5
    },
    map:{
        padding: 5,
        borderBottomWidth: 1,
        borderBottomColor:"#eee"
    }
})
