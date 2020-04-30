import React, {Component} from 'react';
import {Alert, Image, NativeModules, ScrollView, StyleSheet, TouchableOpacity, View} from 'react-native';
import AsyncStorage from "@react-native-community/async-storage";
import {Body, CardItem, Icon, Text} from "native-base";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import CustomHeader from "../../headers/CustomHeader";
import MapInput from "../../utils/MapInput";
import CreateCity from "./CreateCity";
import CityDetail from "./CityDetail";
import Style from "../../../styles/Style";
export default class SearchCity extends Component {
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
        console.log("in callbackFunction", this.state.fcity)
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
        console.log("city", city)
        console.log("in alert", this.state.fcity)
        Alert.alert(
            'City is not created yet.',
            ' Would you like to create a city? ',
            [
                {
                    text: 'Yes', onPress: () => this.props.navigation.navigate('CreateCity',
                        { // country: item.country,
                            city: this.state.fcity,
                            country: this.state.country,
                        })
                },
                {
                    text: 'No',
                    onPress: () => this.props.navigation.navigate('DisplayCities'),
                    style: 'cancel',

                }
            ]
        );
    }

    getCity(){
        console.log("in render", this.state.city, this.state.country)
        if (this.state.city !== 'undefined') {
            let newCityList = []
            this.state.cities.forEach(element => {
                if (element.name.includes(this.state.city) && element.country.includes(this.state.country)) {
                    newCityList.push(element)
                    this.setState({cities: newCityList})
                }
            })
            this.state.city = 'undefined'
            if (newCityList.length === 0) {
                this.showAlert(this.state.city)
            }
        } else {
            //console.log("is undefined")
        }
    }

    render() {
        if (this.state.city === 'undefined') {
            return (
            <View style={Style.view}>
                <CustomHeader style={{flex: 1}} title="Search City" isHome={false} navigation={this.props.navigation}/>
                <View style={{flex: 1}}>
                    <MapInput navigation={this.props.navigation} notifyChange={() => this.getCity()}
                              parentCallback={this.callbackFunction}/>
                </View>
            </View>
        )
        }
        else{


        return (
            <View style={Style.view}>
                <CustomHeader style={{flex: 0.4}} title="Found City" isHome={false} navigation={this.props.navigation}/>
                {/*<View style={{flex: 0.4}}>*/}
                {/*    <MapInput navigation={this.props.navigation} notifyChange={() => this.getCity()}*/}
                {/*              parentCallback={this.callbackFunction}/>*/}
                {/*</View>*/}
                <ScrollView style={{flex: 1}}>

                    {this.state.cities.map((item, index) => {
                        return (
                            <TouchableOpacity onPress={() => this.props.navigation.navigate('CityDetail', {
                                cityId: item.cityId,
                                name: item.name,
                                indexId: item.indexId,
                                country: item.country,
                                description: item.description,
                            })}>
                                <Card style={Style.cardContainer} key={this.state.cities.cityId} pointerEvents="none">
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
                                    </CardItem>
                                </Card>
                            </TouchableOpacity>
                        )

                    })}
                </ScrollView>
            </View>

        )
        }
    }
}
