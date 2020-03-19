import React, {Component} from 'react';
import {Image, NativeModules, ScrollView, StyleSheet, View} from 'react-native';
import AsyncStorage from "@react-native-community/async-storage";
import {Body, CardItem, Icon, Text} from "native-base";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
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
            max: 999999,
            cities: [

                {
                    cityId: 0,
                    name: '',
                    country: '',
                    creatorEmail: '',
                    description: '',
                    lat: 0.0,
                    lng: 0.0,
                },
            ],
            photoMap:[

            ],
            isUpdated: false
        }
    }

    // callbackFunction = (lat, lng, city, country) => {
    //     this.setState({lat: lat})
    //     this.setState({lng: lng})
    //     this.setState({city: city})
    //     this.setState({country: country})
    //     console.log(this.state.lat, this.state.lng)
    // }
    // shouldComponentUpdate(prevProps, prevState){
    //     console.log("in  shouldComponentUpdate1")
    //     //this.displayCities()
    //
    //     if (prevState.cities !== this.state.cities) {
    //         console.log("in  shouldComponentUpdate")
    //         this.setState({isUpdated: false})
    //
    //     }
    // }
    getCitiesPhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.PhotosModule.getCitysPhoto(
                            token,
                            email,
                            (err) => {
                                console.log(err)
                            },

                            (jsonCityPhotoList) => {
                                this.setState({photoMap:jsonCityPhotoList})

                            })
                    }
                })
            })
        })
    }

    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]

                    if (value != null) {
                        NativeModules.ProfilesModule.getAllCities(
                            this.state.max,
                            (err) => {
                                console.log(err)
                            },

                            (jsonCityList) => {
                                this.setState({cities: JSON.parse(jsonCityList)})
                                // this.setState({name: JSON.stringify(jsonCityList.name)})

                            })
                    }
                })
            })
        })
        this.getCitiesPhoto()
    }

    updateCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]

                    if (value !== null) {
                        // NativeModules.ProfilesModule.updateCity(
                        //     value,
                        //     this.state.name,
                        //     this.state.country,
                        //     key,
                        //     this.state.description,
                        //     this.state.lat,
                        //     this.state.lng,
                        //     (err) => {
                        //         console.log("err in updateCity " + err)
                        //     },
                        //     (name, country, email, description, lat, lng, id) => {
                        //         this.setState({name: name})
                        //         this.setState({country: country})
                        //         // this.setState({email: email})
                        //         this.setState({description: description})
                        //         this.setState({lat: lat})
                        //         this.setState({lng: lng})
                        //         console.log("name, country, email, description, lat, lon id in updateCity is " + name, country, email, description, lat, lng, id)
                        //     })

                    }
                })
            })
        })
    }

    render() {
       // let image = this.state.showDefault ? require('../../../img/noImage.png') : {uri: ImageUrl};source={image ? { uri: image } : defaultImg}

        return (
            <View style={{flex: 1}}>
                {/*<GeoLoc parentCallback={this.callbackFunction}/>*/}
                <CustomHeader title="Cities" isHome={true} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <View style={{flex: 1}}>
                        {/*<MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()}*/}
                        {/*          parentCallback={this.callbackFunction}/>*/}
                    </View>
                    {this.state.cities.map((item, index) => {
                        return (
                            <Card key={this.state.cities.cityId}>
                                <CardItem cardBody>

                                    <Image source={{uri: this.state.photoMap[item.cityId]}}
                                           style={{height: 200, width: null, flex: 1}}
                                    />
                                </CardItem>
                                <CardItem>
                                    <CardTitle
                                        title={item.name}
                                        subtitle={item.country}
                                    />
                                </CardItem>
                                <CardItem>
                                    <Body>
                                        <Text numberOfLines={1} ellipsizeMode={"tail"}>{item.description} </Text>
                                    </Body>
                                    <CardAction
                                        separator={true}
                                        inColumn={false}>
                                        <CardButton
                                            onPress={() => this.props.navigation.navigate('CityDetail', {
                                                cityId: item.cityId,
                                                name: item.name,
                                                indexId: item.indexId,
                                                country: item.country,
                                                description: item.description,
                                            })}
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
