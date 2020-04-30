import React, {Component} from 'react'
import {
    View,
    Image,
    Button,
    StyleSheet,
    ImageBackground,
    ScrollView,
    NativeModules, TouchableOpacity, Dimensions, SafeAreaView,
} from 'react-native'
import {Body, CardItem, Text} from 'native-base'
import CustomHeader from '../headers/CustomHeader'
import GeoLoc from "../utils/GeoLoc";
import {IMAGE} from "../../constants/Image";
import PhotoUpload from "react-native-photo-upload"
import {Card} from 'react-native-elements'
import Settings from "./Settings"
import AsyncStorage from "@react-native-community/async-storage"
import Carousel from "react-native-snap-carousel";
import {CardTitle} from "react-native-material-cards";
import Style from '../../styles/Style'

const {width: viewWidth} = Dimensions.get('window')
let email

class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            userId: -999,
            avatar_url: '',
            image: '',
            email: email,
            name: '',
            description: '',
            message: '',
            visitedCitiesPhotoList: [],
            visitedPlacesPhotoList: [],
            visitedCities: [
                {
                    cityId: 0,
                    name: '',
                    postId: 0,
                    country: '',
                    description: '',
                }
            ],
            getVisitedPlaces: [
                {
                    city: '',
                    country: '',
                    description: '',
                    name: '',
                    id: 0,
                }
            ]
        }
    }

    callbackFunction = (lat, lng, city, country) => {
        this.setState({city: city})
        this.setState({country: country})
    }

    componentDidMount() {
        let visitedCities = []
        let keys = [...global.visitedCityMap.keys()]

        keys.forEach(e => {
            if (global.visitedCityMap[e] === true) {
                visitedCities.push(e)
            }
        })

        let visitedPlaces = []
        let pkeys = [...global.visitedPlaceMap.keys()]

        pkeys.forEach(e => {
            if (global.visitedPlaceMap[e] === true) {
                visitedPlaces.push(e)
            }
        })

        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.PhotosModule.getProfilePhoto(
                            email,
                            token,
                            (err) => {
                                console.log(err)
                            },
                            (url) => {
                                this.setState({avatar_url: url})
                            })
                        NativeModules.ProfilesModule.getUser(
                            token,
                            email,
                            (err) => {
                                console.log(err)
                            },
                            (email, name, description, userId) => {
                                this.setState({name: name})
                                this.setState({description: description})
                                this.setState({userId: userId})
                            })
                        NativeModules.ProfilesModule.getVisitedCities(
                            token,
                            email,
                            (err) => {
                                console.log(err)
                            },
                            (json) => {
                                this.setState({visitedCities: JSON.parse(json)})
                            })

                        NativeModules.ProfilesModule.getVisitedPlaces(
                            token,
                            email,
                            (err) => {
                                console.log(err)
                            },
                            (placesJson) => {
                                this.setState({visitedPlaces: JSON.parse(placesJson)})
                            })

                        NativeModules.PhotosModule.getVisitedCitysPhotos(
                            email,
                            token,
                            visitedCities,
                            (err) => {
                                console.log(err)
                            },

                            (visitedCitiesPhotoList) => {
                                this.setState({visitedCitiesPhotoList: visitedCitiesPhotoList})
                            })

                        NativeModules.PhotosModule.getVisitedPlacesPhotos(
                            email,
                            token,
                            visitedPlaces,
                            (err) => {
                                console.log(err)
                            },

                            (visitedPlacesPhotoList) => {
                                this.setState({visitedPlacesPhotoList: visitedPlacesPhotoList})
                            })
                    }
                })
            })
        })
    }

    updatePhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    email = store[i][0];
                    let value = store[i][1]
                    if (value !== null) {
                        NativeModules.PhotosModule.uploadProfilePhoto(
                            email,
                            value,
                            this.state.image,

                            (err) => {
                                console.log(err)
                            },
                            (url) => {
                                this.setState({avatar_url: url})
                            })
                    }
                })
            })
        })
    }

    renderItemVisitedCities = ({item, map}) => {
        return (
            <View>
                <TouchableOpacity onPress={() => this.props.navigation.navigate('CityDetail', {
                    cityId: item.cityId,
                    name: item.name,
                    country: item.country,
                    description: item.description,
                })}>
                    <Image source={{uri: this.state.visitedCitiesPhotoList["" + item.cityId]}}
                           style={Style.cardPhoto}/>

                    <Text style={styles.title}>{item.name}, {item.country}</Text>
                    <Text numberOfLines={1} ellipsizeMode={"tail"} style={styles.title}>{item.description}</Text>

                </TouchableOpacity>
            </View>
        )
    }

    renderItemVisitedPlaces = ({item, map}) => {
        return (
            <View>
                <TouchableOpacity onPress={() => this.props.navigation.navigate('PlaceDetail', {
                    placeId: item.id,
                    name: item.name,
                    city: item.city,
                    country: item.country,
                    description: item.description,
                })}>


                    <Image source={{uri: this.state.visitedPlacesPhotoList["" + item.id]}}
                           style={Style.cardPhoto}/>

                    <Text style={styles.title}>{item.name}</Text>
                    <Text style={styles.title}>{item.city}, {item.country}</Text>
                    <Text numberOfLines={1} ellipsizeMode={"tail"} style={styles.title}>{item.description}</Text>

                </TouchableOpacity>
            </View>
        )
    }

    render() {
        if (this.state.description === 'null') {
            this.setState({description: 'No description provided'})
        }
        if (this.state.name === 'null') {
            this.setState({name: 'No name provided'})
        }
        return (
            <View style={Style.view}>
                <CustomHeader title="Profile" navigation={this.props.navigation}/>
                <ScrollView style={Style.view}>
                    <Card style={Style.cardContainer}>
                        <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
                            <PhotoUpload onPhotoSelect={avatar => {
                                if (avatar) {
                                    this.setState({image: avatar})
                                    this.updatePhoto()
                                }
                            }
                            }>
                                {this.displayPhoto()}
                            </PhotoUpload>
                            <CardItem>
                                <GeoLoc parentCallback={this.callbackFunction}/>
                                <Text style={Style.title}>{this.state.city} {this.state.country} </Text>
                            </CardItem>
                        </View>
                        <View>
                            <CardItem>
                                <Text style={Style.title}>Email: </Text>
                                <Text> {email} </Text>
                            </CardItem>

                            <CardItem>
                                <Text style={Style.title}>Name: </Text>
                                <Text>{this.state.name} </Text>
                            </CardItem>
                            <CardItem>
                                <Text style={Style.title}>Description: </Text>
                                <Text> {this.state.description} </Text>
                            </CardItem>
                        </View>
                        <TouchableOpacity style={Style.btnPressStyle}
                                          onPress={() => this.props.navigation.navigate("Settings", {userId: this.state.userId})}>
                            <Text style={Style.txtStyle}>Edit Profile</Text>
                        </TouchableOpacity>
                    </Card>

                    {/*Visited Cities card*/}
                    <Card style={Style.cardContainer}>
                        <View style={{flex: 1}}>
                            <Text style={Style.title}> Visited Cities</Text>
                        </View>
                        <Carousel
                            ref={(c) => {
                                this._carousel = c;
                            }}
                            data={this.state.visitedCities}
                            map={this.state.visitedCitiesPhotoList}
                            renderItem={this.renderItemVisitedCities}
                            sliderWidth={viewWidth / 1.2}
                            itemWidth={viewWidth}
                        />
                    </Card>

                    {/*Visited Places card*/}

                    <Card style={Style.cardContainer}>
                        <View style={{flex: 1}}>
                            <Text style={Style.title}> Visited Places</Text>
                        </View>
                        <Carousel
                            ref={(c) => {
                                this._carousel = c;
                            }}
                            data={this.state.visitedPlaces}
                            map={this.state.visitedPlacesPhotoList}
                            renderItem={this.renderItemVisitedPlaces}
                            sliderWidth={viewWidth / 1.2}
                            itemWidth={viewWidth}
                        />
                    </Card>
                </ScrollView>
            </View>
        )
    }

    displayPhoto() {
        if (this.state.avatar_url === '') {
            return (<Image source={IMAGE.ICON_DEFAULT_PROFILE} style={Style.profilePhoto}/>)
        } else {
            return (<Image source={{uri: this.state.avatar_url}}
                           style={Style.profilePhoto}/>)
        }
    }
}

export default Profile
const styles = StyleSheet.create({
    cardContainer: {
        backgroundColor: '#FFF',
        borderWidth: 0,
        flex: 1,
        margin: 0,
        padding: 0,
    },
    container: {
        flex: 1,
    },
    emailContainer: {
        backgroundColor: '#FFF',
        flex: 1,
        paddingTop: 30,
    },
    headerBackgroundImage: {
        paddingBottom: 20,
        paddingTop: 35,
    },
    headerContainer: {},
    headerColumn: {
        backgroundColor: 'transparent',
        ...Platform.select({
            ios: {
                alignItems: 'center',
                elevation: 1,
                marginTop: -1,
            },
            android: {
                alignItems: 'center',
            },
        }),
    },
    placeIcon: {
        color: 'white',
        fontSize: 26,
    },
    scroll: {
        backgroundColor: '#FFF',
    },
    userImage: {
        // borderColor: mainColor,
        borderRadius: 85,
        borderWidth: 3,
        height: 170,
        marginBottom: 15,
        width: 170,
    },
    name: {
        fontSize: 22,
        color: "#FFFFFF",
        fontWeight: '600',
    },
    buttonContainer: {
        marginTop: 10,
        height: 45,
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 20,
        width: 250,
        borderRadius: 30,
        backgroundColor: "#007AFF",
    },
})
