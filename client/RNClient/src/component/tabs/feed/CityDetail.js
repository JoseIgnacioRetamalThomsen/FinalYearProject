import React, {Component} from 'react';
import {Image, NativeModules, ScrollView, StyleSheet, View} from 'react-native';
import {Body, CardItem, Icon, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import MapInput from "../../MapInput";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";
import Carousel from 'react-native-snap-carousel';
import CreatePlace from './CreatePlace'

class CityDetail extends Component {
    constructor(props) {
        super(props);
        this.state = {
            cityId: 0,
            indexId: 0,
            city: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
            images: [
                {
                    url: "",
                    timestamp: ""
                },
                {
                    url: "",
                    timestamp: ""

                }

            ],
            places: [
                {
                    id: 0,
                    name: '',
                    city: '',
                    country: '',
                    description: '',
                    url:''
                }
            ],
            posts: [
                {
                    indexId: '',
                    creatorEmail: '',
                    cityName: '',
                    cityCountry: '',
                    title: '',
                    body: '',
                    timeStamp: '',
                    likes: [],
                    mongoId: '',
                }
            ]
        }
    }

    componentDidMount() {
        const cityId = this.props.navigation.getParam('cityId', '')
        const indexId = this.props.navigation.getParam('indexId', '')
        const city = this.props.navigation.getParam('name', '')
        const country = this.props.navigation.getParam('country', '')
        const description = this.props.navigation.getParam('description', '')
        const img = this.props.navigation.getParam('img', '')

        this.setState({
            cityId,
            city,
            country,
            description,
            img
        })
        console.log('cityDetail' + cityId, city, country)
        this.getCityImages()
        this.getCityPlaces()
    }

    getCityImages() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    console.log("email/token in getCityImages " + email + " " + token)

                    if (token != null) {
                        NativeModules.PhotosModule.getCityImage(
                            token,
                            email,
                            parseInt(this.state.cityId),
                            (err) => {
                                console.log("error In PhotoModule.getCityImage " + err)
                            },
                            (images) => {
                                this.setState({images: JSON.parse(images)})
                                console.log("image json", this.state.images.url)
                            })
                    }
                })
            })
        })
    }

    getCityPlaces() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    this.setState({email: email})
                    let token = store[i][1]
                    if (token != null) {
                        NativeModules.ProfilesModule.getCityPlaces(
                            token,
                            email,
                            this.state.city,
                            this.state.country,
                            (err) => {
                                console.log(err)
                            },
                            (placesList) => {
                                console.log(" placesList is !!!", placesList, this.state.city, this.state.country)
                                this.setState({places: JSON.parse(placesList)})
                            })
                    }
                })
            })
        })
    }

    _renderItem = ({item, index}) => {
        console.log(item, index);
        return (
            <View style={styles.slide}>
                <Text style={styles.title}>{item.timestamp}</Text>
                <Image source={{uri: item.url}}
                       style={{height: 200, width: null, flex: 1}}/>
            </View>
        )
    }

    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title={this.state.city} isHome={false} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <Card>
                        <CardItem>
                            <CardTitle
                                title={this.state.city}
                                subtitle={this.state.country}
                            />
                        </CardItem>

                        <CardItem cardBody>
                            <Image source={this.state.img}
                                   style={{height: 200, width: null, flex: 1}}/>
                        </CardItem>
                        <Carousel
                            ref={(c) => {
                                this._carousel = c;
                            }}
                            data={this.state.images}
                            renderItem={this._renderItem}
                            sliderWidth={500}
                            itemWidth={500}
                        />
                        <CardItem>
                            <Body>
                                <Text>{this.state.description} </Text>
                            </Body>
                            <CardAction
                                separator={true}
                                inColumn={false}>
                                <CardButton
                                    onPress={() => this.props.navigation.navigate('CreateCity')}
                                    title="Edit"
                                    color="blue"
                                />
                            </CardAction>
                        </CardItem>
                    </Card>

                    <View style={{flex: 1}}>

                    </View>
                    {this.state.places.map((item, index) => {
                        return (
                            <Card key={this.state.places.placeId}>

                                <CardItem>
                                    <CardTitle
                                        title={item.name}
                                    />
                                </CardItem>
                                {/*<CardItem cardBody>*/}
                                {/*    <Image source={{url: this.state.url}}*/}
                                {/*           style={{height: 200, width: null, flex: 1}}/>*/}
                                {/*</CardItem>*/}
                                <CardItem>
                                    <Body>
                                        <Text numberOfLines={1} ellipsizeMode={"tail"}>{item.description} </Text>
                                    </Body>
                                    <CardAction
                                        separator={true}
                                        inColumn={false}>
                                        <CardButton
                                            onPress={() => this.props.navigation.navigate('PlaceDetail', {
                                                placeId: item.id,
                                                name: item.name,
                                                city: this.state.city,
                                                cityId: this.state.cityId,
                                                country: this.state.country,
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
                    <ActionButton.Item buttonColor='#007AFF' title="Write a post about this city"
                                       onPress={() => this.props.navigation.navigate('CreateCityPost', {
                                           indexId: this.state.indexId,
                                           city: this.state.city,
                                           country: this.state.country
                                       })}>
                        <Icon name="md-create" style={styles.actionButtonIcon}/>
                    </ActionButton.Item>
                    <ActionButton.Item buttonColor='#007AFF' title="Add a place"
                                       onPress={() => this.props.navigation.navigate('CreatePlace', {
                                           indexId: this.state.indexId,
                                           name: this.state.city,
                                           country: this.state.country,
                                           email: this.state.email,
                                           image: this.state.image
                                       })}>
                        <Icon name="md-create" style={styles.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>
            </View>

        )
    }
}

export default CityDetail
const styles = StyleSheet.create({
    actionButtonIcon: {
        fontSize: 20,
        height: 22,
        color: 'white',

    },
})
