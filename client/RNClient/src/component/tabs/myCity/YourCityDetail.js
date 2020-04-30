import React, {Component} from 'react';
import {
    Dimensions,
    Image,
    NativeModules,
    ScrollView,
    TextInput,
    TouchableOpacity,
    View
} from 'react-native';
import {Body, CardItem, Icon, Text} from 'native-base';
import {Card, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";
import Carousel from 'react-native-snap-carousel';
import Style from '../../../styles/Style'
import Modal, {ModalContent} from 'react-native-modals';
import SlideAnimation from "react-native-modals/dist/animations/SlideAnimation";
import PhotoUpload from "react-native-photo-upload";
import {Button, Divider} from "react-native-elements";
import SpecialHeader from "../../headers/SpecialHeader";
import {IMAGE} from "../../../constants/Image";
import GeoLoc from "../../utils/GeoLoc";

const {width: viewWidth} = Dimensions.get('window')

class YourCityDetail extends Component {
    constructor(props) {
        global.isVisitedPlace = true
        super(props);
        this.state = {
            placesPhotoMap: [],
            postsPhotoMap: [],
            //city
            isVisible: false,
            isVisible2: false,
            cityId: -99,
            indexId: 0,
            city: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
            cityImage: '',
            cityUrl: '',
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
            //cityPost
            cityPostId: '',
            cityPostTitle: '',
            cityPostBody: '',
            cityPostImage: '',
            cityPostUrl: '',

            //place
            placeId: 0,
            placeImage: '',
            placeUrl: '',
            placeName: '',
            placeDescription: '',

            places: [
                {
                    id: 0,
                    name: '',
                    city: '',
                    country: '',
                    description: '',
                    url: ''
                }
            ],
            posts: [
                {
                    body: '',
                    cityCountry: '',
                    cityName: '',
                    creatorEmail: '',
                    indexId: '',
                    likes: [],
                    mongoId: '',
                    timeStamp: '',
                    title: '',
                }
            ]
        }
    }

    callbackFunction = (lat, lng, city, country) => {
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
                                console.log("gggg", err)
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
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.ProfilesModule.getCity(
                            email,
                            token,
                            this.state.city,
                            this.state.country,
                            (err) => {
                                console.log("err", err)
                            },
                            (city, country, email, lat, lng, description, id) => {
                                this.setState({city: city})
                                this.setState({country: country})
                                this.setState({email: email})
                                this.setState({lat: lat})
                                this.setState({lng: lng})
                                this.setState({description: description})
                                console.log("successful in getCity values " + city, country, email, description, id)

                            })
                    }
                })
            })
        })
        this.getCitiesPhoto()
    }

    getCityPosts() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let token = store[i][1]
                    if (token !== null) {
                        NativeModules.PostModule.getCityPosts(
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },
                            (cityPosts) => {
                                this.setState({posts: JSON.parse(cityPosts)})
                            })

                    }

                })
            })
        })
    }

    getCityImages() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token != null) {
                        NativeModules.PhotosModule.getCityImage(
                            token,
                            email,
                            parseInt(this.state.cityId),
                            (err) => {
                                console.log(err)
                            },
                            (images) => {
                                this.setState({images: JSON.parse(images)})
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
                                this.setState({places: JSON.parse(placesList)})
                            })
                    }
                })
            })
        })
    }

    createCityPost() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let value = store[i][1]

                    if (value !== null) {
                        NativeModules.PostModule.createCityPost(
                            this.state.cityId,
                            email,
                            this.state.city,
                            this.state.country,
                            this.state.cityPostTitle,
                            this.state.cityPostBody,
                            (err) => {
                                console.log(err)
                            },
                            (cityPostId) => {
                                this.setState({cityPostId: cityPostId})
                                this.uploadPostImage();
                                this.setState({isVisible: false})

                            })
                    }
                })
            })
        })
    }

    createPlace() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    NativeModules.ProfilesModule.createPlace(
                        token,
                        email,
                        this.state.placeName,
                        this.state.city,
                        this.state.country,
                        89,
                        90,
                        //parseFloat(this.state.lon),
                        this.state.placeDescription,
                        (err) => {
                            console.log(err)
                        },
                        (placeId) => {
                            this.setState({placeId: placeId})
                            this.uploadPlacePhoto();
                            this.setState({isVisible2: false})
                            this.getCityPlaces()
                            this.getPlacesPerCityPhoto()
                            //this.props.navigation.navigate('DisplayPlaces')
                        })
                })
            })
        })
    }

    uploadPlacePhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    if (token != null) {
                        NativeModules.PhotosModule.uploadPlacePhoto(
                            token,
                            email,
                            parseInt(this.state.placeId),
                            this.state.placeImage,
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },
                            (placeUrl) => {
                                this.setState({placeUrl: placeUrl})
                            })
                    }
                })
            })
        })
    }

    uploadCityPhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    if (token !== null) {
                        NativeModules.PhotosModule.uploadCityPhoto(
                            token,
                            email,
                            parseInt(this.state.cityId),
                            this.state.cityImage,
                            (err) => {
                                console.log(err)
                            },
                            (cityUrl) => {
                                this.setState({cityUrl: cityUrl})
                            })
                    }
                })
            })
        })
    }

    uploadPostImage() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0]
                    let token = store[i][1]
                    if (token !== null) {
                        NativeModules.PhotosModule.uploadPostImage(
                            token,
                            email,
                            this.state.cityPostId,
                            this.state.cityPostImage,
                            0,
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },
                            (cityUrl) => {
                                this.setState({cityUrl: cityUrl})
                            })
                    }
                })
            })
        })
    }

    getPlacesPerCityPhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.PhotosModule.getPlacesPerCityPhoto(
                            token,
                            email,
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },

                            (jsonCityPhotoList) => {
                                this.setState({placesPhotoMap: jsonCityPhotoList})
                            })
                    }
                })
            })
        })
    }

    getPostsPhotosIdP() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.PhotosModule.getPostsPhotosIdP(
                            token,
                            email,
                            0,
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },

                            (photoList) => {
                                this.setState({postsPhotoMap: photoList})

                            })
                    }
                })
            })
        })
    }

    _renderItem = ({item, index}) => {
        return (
            <View style={Style.slide}>
                <Text style={Style.title}>{item.timestamp}</Text>
                <Image source={{uri: item.url}}
                       style={{height: 200, width: null, flex: 1}}/>
            </View>
        )
    }

    renderItemPlaces = ({item, index}) => {
        return (
            <View style={Style.carouselContainer}>
                <TouchableOpacity onPress={() => this.props.navigation.navigate('PlaceDetail', {
                    placeId: item.id,
                    name: item.name,
                    city: this.state.city,
                    cityId: this.state.cityId,
                    country: this.state.country,
                    description: item.description,
                })}>

                    <View>
                        <View>
                            <Image source={{uri: this.state.placesPhotoMap["" + item.id]}}
                                   style={Style.cardPhoto}/>
                        </View>

                        <View>
                            <Text style={Style.title}>{item.name}</Text>
                            <Text style={Style.text}>{item.description}</Text>
                        </View>
                    </View>
                </TouchableOpacity>
            </View>
        )
    }

    render() {
        return (
            <View style={Style.view}>
                {/*create city Post*/}
                <Modal
                    visible={this.state.isVisible}
                    modalAnimation={new SlideAnimation({
                        slideFrom: 'bottom',
                    })}
                    onTouchOutside={() => {
                        this.setState({isVisible: false});
                    }}
                >
                    <ModalContent style={{
                        width: Dimensions.get('window').width * 0.8, height: Dimensions.get('window').height * 0.6
                    }}>
                        <TextInput
                            placeholder="Title"
                            underlineColorAndroid='transparent'
                            onChangeText={(cityPostTitle) => this.setState({cityPostTitle})}/>

                        <TextInput
                            placeholder="Description"
                            underlineColorAndroid='transparent'
                            onChangeText={(cityPostBody) => this.setState({cityPostBody})}/>

                        <PhotoUpload onPhotoSelect={cityPostImage => {
                            if (cityPostImage) {
                                this.setState({cityPostImage: cityPostImage})
                            }
                        }
                        }>
                            <Image source={IMAGE.UPLOAD_IMG}/>
                        </PhotoUpload>

                        <Button title="Add city post" style={{height: 200, width: 200, borderColor: 'black'}}
                                onPress={() => this.createCityPost()}>

                        </Button>
                    </ModalContent>

                </Modal>

                {/*create place*/}
                <Modal
                    visible={this.state.isVisible2}
                    modalAnimation={new SlideAnimation({
                        slideFrom: 'bottom',
                    })}
                    onTouchOutside={() => {
                        this.setState({isVisible2: false});
                    }}
                >
                    <ModalContent style={{
                        width: Dimensions.get('window').width * 0.8, height: Dimensions.get('window').height * 0.6
                    }}>
                        <TextInput
                            placeholder="Title"
                            underlineColorAndroid='transparent'
                            onChangeText={(placeName) => this.setState({placeName})}/>

                        <TextInput
                            placeholder="Description"
                            underlineColorAndroid='transparent'
                            onChangeText={(placeDescription) => this.setState({placeDescription})}/>

                        <PhotoUpload onPhotoSelect={placeImage => {
                            if (placeImage) {
                                this.setState({placeImage: placeImage})
                            }
                        }
                        }>
                            <Image source={IMAGE.UPLOAD_IMG}/>
                        </PhotoUpload>

                        <Button title="Add a place" style={Style.btnPressStyle} onPress={() => this.createPlace()}>

                        </Button>
                    </ModalContent>
                </Modal>

                <SpecialHeader title={this.state.city} isHome={true}
                               navigation={this.props.navigation}/>
                <GeoLoc parentCallback={this.callbackFunction}/>

                <ScrollView style={{flex: 1}}>
                    <Card style={Style.cardContainer}>
                        <Carousel
                            ref={(c) => {
                                this._carousel = c;
                            }}
                            data={this.state.images}
                            renderItem={this._renderItem}
                            sliderWidth={viewWidth / 1.055}
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
                                <Text>{this.state.description} </Text>
                            </Body>

                        </CardItem>
                    </Card>

                    <Divider style={{backgroundColor: 'black', margin: 10}}/>

                    {/*Places*/}

                    <Text style={Style.heading}> Places in {this.state.city} </Text>
                    <Divider style={{backgroundColor: 'black'}}/>

                    {this.displayPlaces()}

                    <Carousel
                        ref={(c) => {
                            this._carousel = c;
                        }}
                        data={this.state.places}
                        index={this.state.placesPhotoMap}
                        renderItem={this.renderItemPlaces}
                        sliderWidth={viewWidth / 1.055}
                        itemWidth={viewWidth}
                    />


                    {/*Posts*/}
                    <Divider style={{backgroundColor: 'black'}}/>
                    <Text style={Style.heading}> Posts about {this.state.city}</Text>
                    <Divider style={{backgroundColor: 'black'}}/>

                    {this.state.posts.map((e, index) => {
                        if (this.state.posts.postId === undefined){
                            return (
                                <Card style={Style.cardContainer} key={this.state.posts.postId}>
                                    <CardItem cardBody>
                                        <Image source={IMAGE.NO_POSTS}
                                               style={Style.noPostsPhoto}/>
                                    </CardItem>
                                </Card>
                            )
                        }else {
                            return (
                                <Card style={Style.cardContainer} key={this.state.posts.postId}>
                                    <CardItem cardBody>
                                        <Image source={{uri: this.state.postsPhotoMap[e.mongoId]}}
                                               style={Style.cardPhoto}/>
                                    </CardItem>

                                    <CardItem>
                                        <Text style={Style.title}>{e.title} </Text>
                                    </CardItem>

                                    <CardItem>
                                        <Body>
                                            <Text style={Style.text} numberOfLines={1}
                                                  ellipsizeMode={"tail"}>{e.body} </Text>
                                        </Body>
                                    </CardItem>
                                </Card>
                            )
                        }
                    })}

                </ScrollView>
                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Write a post about this city"
                                       onPress={() => this.setState({isVisible: true})}>
                        <Icon name="md-create" style={Style.actionButtonIcon}/>
                    </ActionButton.Item>
                    <ActionButton.Item buttonColor='#007AFF' title="Add a place"
                                       onPress={() => this.setState({isVisible2: true})}>
                        <Icon name="md-create" style={Style.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>
            </View>
        )
    }
    displayPlaces() {
        console.log("this.state.places.length", this.state.places.length)
        if (this.state.places.length === 0) {
            return (
                <Card style={Style.cardContainer} key={this.state.posts.postId}>
                    <CardItem cardBody>
                        <Image source={IMAGE.NO_PLACES}
                               style={Style.noPostsPhoto}/>
                    </CardItem>
                </Card>
            )
        }
    }
}

export default YourCityDetail
