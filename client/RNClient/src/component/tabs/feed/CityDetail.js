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

const {width: viewWidth} = Dimensions.get('window')

class CityDetail extends Component {
    constructor(props) {
        global.isVisitedPlace = true
        super(props);
        this.state = {
            clicked: false,
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

    componentDidMount() {
        const cityId = this.props.navigation.getParam('cityId', '')
        const indexId = this.props.navigation.getParam('indexId', '')
        const city = this.props.navigation.getParam('name', '')
        const country = this.props.navigation.getParam('country', '')
        const description = this.props.navigation.getParam('description', '')
        const img = this.props.navigation.getParam('img', '')

        this.setState({
            cityId,
            indexId,
            city,
            country,
            description,
            img
        })
        console.log("this.state.city.Id!!", cityId)
        this.setState({cityId: cityId})
        this.getCityImages()
        this.getCityPlaces()
        this.getPlacesPerCityPhoto()
        this.getCityPosts()
        this.getPostsPhotosIdP()
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
        console.log("this.state.city,", this.state.city)
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
                                console.log("placesList", placesList)
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
                                this.props.navigation.navigate('DisplayCities')
                            })
                    }
                })
            })
        })
    }

    createPlace() {
        this.setState({clicked: true})
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
                            this.props.navigation.navigate('DisplayCities')
                            console.log("placeId", placeId)
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
        console.log("this.state.cityId, in citydet", this.state.cityId, this.state.city)
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
                                console.log("ddd", jsonCityPhotoList)
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
        if (item.url !== '') {
            return (
                <View style={Style.slide}>
                    <Text style={Style.title}>{item.timestamp}</Text>
                    <Image source={{uri: item.url}}
                           style={{height: 200, width: null, flex: 1}}/>
                </View>
            )
        } else {
            return (
                <View style={Style.slide}>
                    <Text style={Style.title}>{item.timestamp}</Text>
                    <Image source={IMAGE.NO_IMG}
                           style={{height: 200, width: null, flex: 1}}/>
                </View>
            )
        }
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

    onClickPlace() {
        if (this.state.placeName === '' || this.state.placeDescription === '' || this.state.placeImage === '') {

            alert("Please upload image and provide title & description")
        } else {
            this.createPlace()
        }
    }

    onClickPost() {
        if (this.state.cityPostTitle === '' || this.state.cityPostBody === '' || this.state.cityPostImage === '') {
            alert("Please upload image and provide title & description")
        } else {
            this.createCityPost()
        }
    }

    render() {
        return (
            <View style={Style.view}>
                {/*create city Post*/}
                <Modal style={Style.modal}
                       visible={this.state.isVisible}
                       modalAnimation={new SlideAnimation({
                           slideFrom: 'bottom',
                       })}
                       onTouchOutside={() => {
                           this.setState({isVisible: false});
                       }}
                >
                    <ModalContent style={Style.modalContent}>
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
                            {this.displayCityPostPhoto()}
                        </PhotoUpload>

                        <TouchableOpacity style={Style.modalbtn}
                                          onPress={() => this.onClickPost()}>
                            <Text style={Style.txtStyle}>Add city post</Text>
                        </TouchableOpacity>
                    </ModalContent>

                </Modal>

                {/*create place*/}
                <Modal style={Style.modal}
                       visible={this.state.isVisible2}
                       modalAnimation={new SlideAnimation({
                           slideFrom: 'bottom',
                       })}
                       onTouchOutside={() => {
                           this.setState({isVisible2: false});
                       }}
                >
                    <ModalContent style={Style.modalContent}>
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
                            {this.displayPlacePostPhoto()}
                        </PhotoUpload>

                        <TouchableOpacity style={Style.modalbtn}
                                          onPress={() => this.onClickPlace()}>
                            <Text style={Style.txtStyle}>Create Place</Text>
                        </TouchableOpacity>
                    </ModalContent>
                </Modal>

                <SpecialHeader title={this.state.city} cityIdFromParent={this.state.cityId} isHome={false}
                               navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}} keyboardShouldPersistTaps='handled'>
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
                        if (this.state.posts[index].mongoId === '') {
                            return (
                                <Card style={Style.cardContainer} key={this.state.posts.postId}>
                                    <CardItem cardBody>
                                        <Image source={IMAGE.NO_POSTS}
                                               style={Style.noPostsPhoto}/>
                                    </CardItem>
                                </Card>
                            )
                        } else {
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

    displayCityPostPhoto() {
        if (this.state.cityPostImage !== '') {
            return <Image style={Style.uploadPhoto} source={{cityPostImage: this.state.cityPostImage}}/>
        } else return <Image source={IMAGE.UPLOAD_IMG}/>
    }

    displayPlacePostPhoto() {
        if (this.state.placeImage !== '') {
            return <Image style={Style.uploadPhoto} source={{placeImage: this.state.placeImage}}/>
        } else return <Image source={IMAGE.UPLOAD_IMG}/>
    }

    displayPlaces() {
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

export default CityDetail
