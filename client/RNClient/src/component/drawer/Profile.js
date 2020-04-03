import React, {Component} from 'react'
import {
    View,
    Image,
    Button,
    StyleSheet,
    ImageBackground,
    ScrollView,
    NativeModules, TouchableOpacity, Dimensions,
} from 'react-native'
import {Body, CardItem, Text} from 'native-base'
import CustomHeader from '../headers/CustomHeader'
import GeoLoc from "../GeoLoc";
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
                                // console.log("placesJson", placesJson)
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
                                console.log("visitedPhotoCityList", visitedCitiesPhotoList)
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
                                console.log("visitedPhotoPlaceList", visitedPlacesPhotoList)
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

                <Text style={styles.title}>{item.name}</Text>
                <Text style={styles.title}>{item.country}</Text>
                <Text style={styles.title}>{item.description}</Text>

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
                <Text style={styles.title}>{item.city}</Text>
                <Text style={styles.title}>{item.country}</Text>
                <Text style={styles.title}>{item.description}</Text>

                </TouchableOpacity>
            </View>
        )
    }

    render() {
        if (this.state.avatar_url === '') {
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
                                    <Image source={IMAGE.ICON_DEFAULT_PROFILE} style={Style.profilePhoto}/>
                                </PhotoUpload>
                                <GeoLoc></GeoLoc>
                            </View>
                            <View>
                                <CardItem>
                                    <Text>Email: {email} </Text>
                                </CardItem>
                                <CardItem>
                                    <Text>Name: {this.state.name} </Text>
                                </CardItem>
                                <CardItem>
                                    <Text>Description: {this.state.description} </Text>
                                </CardItem>
                            </View>
                            <Button title="Edit Profile" style={[Style.buttonContainer, Style.loginButton] }
                                    onPress={() => this.props.navigation.navigate("Settings")}/>
                        </Card>

                        {/*Visited Cities card*/}

                            <Card style={Style.cardContainer}>
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
                                <CardItem>
                                    <CardTitle
                                        title={this.state.city}
                                        subtitle={this.state.country}
                                    />
                                </CardItem>
                            </Card>


                        {/*Visited Places card*/}
                            <Card style={Style.cardContainer} >
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
                                <CardItem>
                                    <CardTitle
                                        title={this.state.city}
                                        subtitle={this.state.country}
                                    />
                                </CardItem>
                            </Card>
                    </ScrollView>
                </View>
            )
        } else {
            return (
                <View style={Style.view}>
                    <CustomHeader title="Profile" navigation={this.props.navigation}/>
                    <ScrollView style={Style.view}>
                        <Card style={Style.cardContainer}>
                            <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
                                <PhotoUpload onPhotoSelect={avatar => {
                                    // console.log("Avatar!!   ", avatar)
                                    if (avatar) {
                                        this.setState({image: avatar})
                                        this.updatePhoto()
                                    }
                                }
                                }>
                                    <Image source={{uri: this.state.avatar_url}}
                                           style={Style.profilePhoto}/>
                                </PhotoUpload>
                                <GeoLoc></GeoLoc>
                            </View>
                            <View>
                                <CardItem>
                                    <Text>Email: {email} </Text>
                                </CardItem>
                                <CardItem>
                                    <Text>Name: {this.state.name} </Text>
                                </CardItem>
                                <CardItem>
                                    <Text>Description: {this.state.description} </Text>
                                </CardItem>
                            </View>
                            <Button style={styles.buttonContainer} title="Edit Profile"
                                    onPress={() => this.props.navigation.navigate("Settings")}/>
                        </Card>

                        {/*Visited Cities card*/}
                            <Card style={Style.cardContainer}>
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
                                <CardItem>
                                    <CardTitle
                                        title={this.state.city}
                                        subtitle={this.state.country}
                                    />
                                </CardItem>

                                <CardItem>
                                    <Body>

                                    </Body>
                                </CardItem>
                            </Card>

                        {/*Visited Places card*/}

                            <Card style={Style.cardContainer} >
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
                                <CardItem>
                                    <CardTitle
                                        title={this.state.city}
                                        subtitle={this.state.country}
                                    />
                                </CardItem>
                            </Card>
                    </ScrollView>
                </View>
            )
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
