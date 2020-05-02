import React, {Component} from 'react';
import {Alert, Image, NativeModules, ScrollView, StyleSheet, TouchableOpacity, View} from 'react-native';
import AsyncStorage from "@react-native-community/async-storage";
import {Body, CardItem, Icon, Text} from "native-base";
import {Card, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import CreateCity from "./CreateCity";
import CityDetail from "./CityDetail";
import HomeHeader from "../../headers/HomeHeader";
import Style from "../../../styles/Style";
import {IMAGE} from "../../../constants/Image";

export default class DisplayCities extends Component {
    constructor(props) {
        super(props);
        this.state = {

            isVisible: false,
            max: 999999,
            img: '../../../img/noImage.png',
            city: 'undefined',
            fcity: '',
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
            photoMap: [],
            isUpdated: false
        }
    }

    callbackFunction = (lat, lng, city, country) => {
        this.setState({lat: lat})
        this.setState({lng: lng})
        this.setState({fcity: city})
        this.setState({city: city})

        this.setState({country: country})
    }

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
                                this.setState({photoMap: require('../../../img/noImage.png')})
                                console.log(err)
                            },

                            (jsonCityPhotoList) => {
                                this.setState({photoMap: jsonCityPhotoList})

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
                            })
                    }
                })
            })
        })

        this.getCitiesPhoto()
    }

    showAlert(city) {
        Alert.alert(
            'City is not created yet.',
            ' Would you like to create a city? ',
            [
                {
                    text: 'No',
                    onPress: () => this.props.navigation.navigate('DisplayCities'),
                    style: 'cancel',
                },
                {
                    text: 'Yes', onPress: () => this.props.navigation.navigate('CreateCity',
                        { // country: item.country,
                            city: this.state.fcity,
                            country: this.state.country,
                        })
                }
            ]
        );
    }

    render() {
        if (this.state.city !== 'undefined') {
            let newCityList = []
            this.state.cities.forEach(element => {
                if (element.name.includes(this.state.city)) {
                    newCityList.push(element)
                    this.setState({cities: newCityList})
                }
            })
            this.state.city = 'undefined'
            if (newCityList.length === 0) {
                this.showAlert(this.state.fcity)
            }
        } else {
            //console.log("is undefined")
        }

        return (
            <View style={Style.view}>
                <HomeHeader style={{flex: 1}} title="Cities" isHome={true} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>

                    {this.state.cities.map((item, index) => {
                        if (this.state.photoMap[item.cityId] !== 'undefined\
                        ') {
                            return (
                                <TouchableOpacity onPress={() => this.props.navigation.navigate('CityDetail', {
                                    cityId: item.cityId,
                                    name: item.name,
                                    indexId: item.indexId,
                                    country: item.country,
                                    description: item.description,
                                })}>
                                    <Card style={Style.cardContainer} key={this.state.cities.cityId}
                                          pointerEvents="none">
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
                                                <Text numberOfLines={1}
                                                      ellipsizeMode={"tail"}>{item.description} </Text>
                                            </Body>
                                        </CardItem>
                                    </Card>
                                </TouchableOpacity>
                            )
                        } else {
                            return (
                                <TouchableOpacity onPress={() => this.props.navigation.navigate('CityDetail', {
                                    cityId: item.cityId,
                                    name: item.name,
                                    indexId: item.indexId,
                                    country: item.country,
                                    description: item.description,
                                })}>
                                    <Card style={Style.cardContainer} key={this.state.cities.cityId}
                                          pointerEvents="none">
                                        <CardItem cardBody>
                                            <Image source={IMAGE.NO_IMG}
                                                   style={{height: 200, width: null, flex: 1}}/>
                                        </CardItem>
                                        <CardItem>
                                            <CardTitle
                                                title={item.name}
                                                subtitle={item.country}
                                            />
                                        </CardItem>
                                        <CardItem>
                                            <Body>
                                                <Text numberOfLines={1}
                                                      ellipsizeMode={"tail"}>{item.description} </Text>
                                            </Body>
                                        </CardItem>
                                    </Card>
                                </TouchableOpacity>
                            )
                        }
                    })}
                </ScrollView>
                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Add a city"
                                       onPress={() => this.props.navigation.navigate('SearchCity')}>
                        <Icon name="md-create" style={Style.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>
            </View>
        )
    }
}
