import React, {Component} from 'react';
import {Dimensions, Image, NativeModules, ScrollView, StyleSheet, TextInput, View} from 'react-native';
import {Body, CardItem, Icon, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import MapInput from "../../MapInput";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";
import Carousel from 'react-native-snap-carousel';
import CreatePlace from './CreatePlace'
import Modal, {ModalContent} from 'react-native-modals';
import SlideAnimation from "react-native-modals/dist/animations/SlideAnimation";
import PhotoUpload from "react-native-photo-upload";
import {Button} from "react-native-elements";

class CityDetail extends Component {
    constructor(props) {
        super(props);
        this.state = {
            //city
            isVisible: false,
            isVisible2: false,
            cityId: 0,
            indexId: 0,
            city: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
            cityImage:'',
            cityUrl:'',
            photoMap:[

            ],
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
            cityPostId:'',
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
        this.getCityImages()
        this.getCityPlaces()
        this.getPlacesPerCityPhoto()
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

                    if (value != null) {
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
                                // this.props.navigation.navigate('DisplayCityPosts')
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

    uploadPostImage(){
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
                                this.setState({photoMap:jsonCityPhotoList})
                            })
                    }
                })
            })
        })
    }

    _renderItem = ({item, index}) => {
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
                            <Image source={{cityPostImage: this.state.cityPostImage}}
                                   style={{
                                       height: 120,
                                       width: 120,
                                       borderRadius: 60,
                                       borderColor: 'black',
                                       borderWidth: 5,
                                       flex: 0,
                                       resizeMode: 'cover'
                                   }}/>
                        </PhotoUpload>

                        <Button title = "Add city post" style = {{ height:200, width: 200,   borderColor: 'black'}} onPress={() => this.createCityPost()} >

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
                            <Image source={{placeImage: this.state.placeImage}}
                                   style={{
                                       height: 120,
                                       width: 120,
                                       borderRadius: 60,
                                       borderColor: 'black',
                                       borderWidth: 5,
                                       flex: 0,
                                       resizeMode: 'cover'
                                   }}/>
                        </PhotoUpload>

                        <Button  title = "Add place post" onPress={() => this.createPlace()}>

                        </Button>
                    </ModalContent>
                </Modal>


                <CustomHeader title={this.state.city} isHome={false} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <Card>
                        <CardItem cardBody>
                            <Image source={this.state.img}
                                   style={{height: 200, width: null, flex: 1}}/>
                        </CardItem>
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
                        <Carousel
                            ref={(c) => {
                                this._carousel = c;
                            }}
                            data={this.state.images}
                            renderItem={this._renderItem}
                            sliderWidth={500}
                            itemWidth={500}
                        />
                    </Card>

                    {/*<View style={{flex: 1}}>*/}

                    {/*</View>*/}
                    {this.state.places.map((e, index) => {
                        return (
                            <Card key={this.state.places.placeId}>


                                <CardItem cardBody>
                                    <Image source={{uri: this.state.photoMap[e.id]}}
                                           style={{height: 200, width: null, flex: 1}}/>
                                </CardItem>

                                <CardItem>
                                    <CardTitle
                                        title={e.name}
                                    />
                                </CardItem>

                                <CardItem>
                                    <Body>
                                        <Text numberOfLines={1} ellipsizeMode={"tail"}>{e.description} </Text>
                                    </Body>
                                    <CardAction
                                        separator={true}
                                        inColumn={false}>
                                        <CardButton
                                            onPress={() => this.props.navigation.navigate('PlaceDetail', {
                                                placeId: e.id,
                                                name: e.name,
                                                city: this.state.city,
                                                cityId: this.state.cityId,
                                                country: this.state.country,
                                                description: e.description,
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

                                       onPress={() => this.setState({isVisible: true})}>
                        <Icon name="md-create" style={styles.actionButtonIcon}/>
                    </ActionButton.Item>
                    <ActionButton.Item buttonColor='#007AFF' title="Add a place"

                                       onPress={() => this.setState({isVisible2: true})}>
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
