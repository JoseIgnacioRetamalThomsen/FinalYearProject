import React, {Component} from 'react';
import {Button, Image, NativeModules, ScrollView, StyleSheet, View} from 'react-native';
import AsyncStorage from "@react-native-community/async-storage";
import {Body, CardItem, Icon, Text} from "native-base";
import {Card, CardAction, CardButton, CardContent, CardImage, CardTitle} from "react-native-material-cards";
import CustomHeader from "../../CustomHeader";
import MapInput from "../../MapInput";
import ActionButton from "react-native-action-button";
import CreateCity from "./CreateCity";
import CityDetail from "./CityDetail";
import GeoLoc from "../../GeoLoc"

export default class DisplayCities extends Component {
    constructor(props) {
        super(props);
        this.state = {
            cities: [
                {
                    indexId: 0,
                    city: 'Dublin',
                    country: 'Ireland',
                    email: '',
                    description: 'CreateCity in the east',
                    img: '../../../img/nuig.jpg',
                    lat: 0,
                    lng: 0,

                },
                {
                    indexId: 1,
                    city: 'Moscow',
                    country: 'Russia',
                    email: '',
                    description: 'CreateCity in the south-west',
                    img: '../../../img/noImage.png',
                    lat: 0,
                    lng: 0,

                },
                {
                    indexId: 2,
                    city: 'Galway',
                    country: 'Ireland',
                    email: '',
                    description: 'CreateCity in the west',
                    img: '../../../img/noImage.png',
                    lat: 0,
                    lng: 0,
                },
            ]
        }
    }

    callbackFunction = (lat, lng, city, country) => {
        this.setState({lat: lat})
        this.setState({lng:lng})
        this.setState({city:city})
        this.setState({country:country})
        console.log(this.state.lat, this.state.lng)
    }

    // componentDidMount() {
    //     AsyncStorage.getAllKeys((err, keys) => {
    //         AsyncStorage.multiGet(keys, (err, stores) => {
    //             stores.map((result, i, store) => {
    //                 let key = store[i][0];
    //                 let value = store[i][1]
    //                 console.log("key/value in displayCity " + key + " " + value)
    //
    //                 if (value !== null) {
    //                     NativeModules.ProfilesModule.getCity(
    //                         value,
    //                         this.state.city,
    //                         this.state.country,
    //                         key,
    //                         this.state.description,
    //                         this.state.lat,
    //                         this.state.lng,
    //                         (err) => {
    //                             console.log("error In ProfilesModule.getCity " + err)
    //                         },
    //                         (valid, city, country, email, description, lat, lng, id) => {
    //                             this.setState({city: city})
    //                             this.setState({country: country})
    //                             this.setState({email: email})
    //                             this.setState({description: description})
    //                             // this.setState({lat: lat})
    //                             // this.setState({lng: lng})
    //                             this.setState({id: id})
    //                             console.log("successful values in getCity!!!" + this.state.id, this.state.city, this.state.country, this.state.description)
    //                         })
    //                 }
    //             })
    //         })
    //     })
    // }
    updateCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in updatecity " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.updateCity(
                            value,
                            this.state.name,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lng,
                            (err) => {
                                console.log("err in updateCity " + err)
                            },
                            (name, country, email, description, lat, lng, id) => {
                                this.setState({name: name})
                                this.setState({country: country})
                                // this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lng: lng})
                                console.log("name, country, email, description, lat, lon id in updateCity is " + name, country, email, description, lat, lng, id)
                                console.log("successfully updated a city!!!")
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View style={{flex: 1}}>
                 <GeoLoc parentCallback={this.callbackFunction} />
                <CustomHeader title="Cities" isHome={true} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <View style={{flex: 1}}>
                        <MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()}
                                  parentCallback={this.callbackFunction}/>
                    </View>
                    { this.state.cities.map((item, index) => {
                        return (
                            <Card key={this.state.cities.indexId}>

                                <CardItem>
                                    <CardTitle
                                        title={item.city}
                                        subtitle={item.country}
                                    />
                                </CardItem>

                                <CardItem cardBody>
                                    <Image source={require('../../../img/noImage.png')}
                                           style={{height: 200, width: null, flex: 1}}/>
                                </CardItem>
                                <CardItem>
                                    <Body>
                                        <Text numberOfLines={1} ellipsizeMode={"tail"}>{item.description} </Text>
                                    </Body>
                                    <CardAction
                                        separator={true}
                                        inColumn={false}>
                                        <CardButton
                                            onPress={() => this.props.navigation.navigate('CityDetail', {indexId:item.indexId, city:item.city, country:item.country, description:item.description, img:item.img})}
                                            title="More"
                                            color="blue"
                                        />
                                    </CardAction>
                                </CardItem>
                            </Card>
                        )
                    })}
                </ScrollView>
                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Add a city"
                                       onPress={() => this.props.navigation.navigate('CreateCity')}>
                        <Icon name="md-create" style={styles.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>
            </View>

        );
    }
}
const styles = StyleSheet.create({
    actionButtonIcon: {
        fontSize: 20,
        height: 22,
        color: 'white',
    },
})
