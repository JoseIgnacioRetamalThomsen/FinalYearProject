import React, {Component} from 'react';
import {Alert, Image, NativeModules, ScrollView, StyleSheet, View} from 'react-native';
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
        // this.setState({city: this.state.city})
        console.log("in render", this.state.fcity)
        if (this.state.city !== 'undefined') {
            let newCityList = []
            this.state.cities.forEach(element => {
                if (element.name.toLowerCase().includes(this.state.city)) {
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
            <View style={{flex: 1}}>
                <GeoLoc parentCallback={this.callbackFunction}/>
                <CustomHeader style={{flex: 1}} title="Cities" isHome={true} navigation={this.props.navigation}/>
                <View style={{flex: 1}}>
                    <MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()}
                              parentCallback={this.callbackFunction}/>
                </View>
                <ScrollView style={{flex: 1}}>

                    {this.state.cities.map((item, index) => {
                        return (
                            <Card key={this.state.cities.cityId}>
                                <CardItem cardBody>
                                    {/*<Image source={(image)}*/}
                                    {/*       style={{height: 200, width: null, flex: 1}}*/}
                                    {/*/>*/}
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
