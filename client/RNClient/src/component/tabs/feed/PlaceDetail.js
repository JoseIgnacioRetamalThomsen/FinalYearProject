import React, {Component} from 'react';
import {Dimensions, Image, NativeModules, ScrollView, StyleSheet, Text, TextInput, View} from "react-native";
import Style from '../../../styles/Style'
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import {Body, CardItem, Icon} from "native-base";
import Carousel from "react-native-snap-carousel";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";
import Modal, {ModalContent} from 'react-native-modals';
import {Button, Divider} from "react-native-elements";
import SlideAnimation from "react-native-modals/dist/animations/SlideAnimation";
import PhotoUpload from "react-native-photo-upload";
import PlaceHeader from "../../headers/PlaceHeader"

const {width: viewportWidth} = Dimensions.get('window')

export default class PlaceDetail extends Component {

    constructor(props) {
        super(props);
        this.state = {
            //place
            isVisible: false,
            placeId: 0,
            placeName: '',
            city: '',
            country: '',
            description: '',
            title: '',
            images: [
                {
                    url: '',
                    timestamp: ''
                }
            ],
            photoMap: [],
            //post
            postTitle: '',
            postDescription: '',
            postMongoId: '',
            postImage: '',
            postUrl: '',
            posts: [
                {
                    body: '',
                    placePostId: '',
                    creatorEmail: '',
                    timeStamp: '',
                    title: '',
                    // likes:[]
                    mongoId: '',
                }
            ]
        }
    }

    componentDidMount() {
        const placeId = this.props.navigation.getParam('placeId', '')
        const placeName = this.props.navigation.getParam('name', '')
        const city = this.props.navigation.getParam('city', '')
        const cityId = this.props.navigation.getParam('cityId', '')
        const country = this.props.navigation.getParam('country', '')
        const description = this.props.navigation.getParam('description', '')

        this.setState({
            placeId,
            placeName,
            city,
            cityId,
            country,
            description
        })
        this.getPlaceImages()
        this.getPlacePosts()
        this.getPostsPhotosIdP()
    }

    getPlaceImages() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token != null) {
                        NativeModules.PhotosModule.getPlacePhoto(
                            token,
                            email,
                            parseInt(this.state.placeId),
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

    _renderItem = ({item, index}) => {
        return (
            <View style={Style.slide}>
                <Text style={Style.title}>{item.timestamp}</Text>
                <Image source={{uri: item.url}}
                       style={{height: 200, width: null, flex: 1}}/>
            </View>
        )
    }

    createPlacePost() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let value = store[i][1]

                    if (value !== null) {
                        NativeModules.PostModule.createPlacePost(
                            this.state.placeId,
                            email,
                            this.state.city,
                            this.state.country,
                            this.state.placeName,
                            this.state.postTitle,
                            this.state.postDescription,
                            (err) => {
                                console.log(err)
                            },
                            (mongoId) => {
                                this.setState({postMongoId: mongoId})
                                this.uploadPostImage();
                                this.setState({isVisible: false})
                            })
                    }
                })
            })
        })
        this.getPlacePosts()
        this.getPlaceImages()
    }

    getPlacePosts() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let value = store[i][1]

                    if (value !== null) {
                        NativeModules.PostModule.getPlacePosts(
                            this.state.placeId,
                            (err) => {
                                console.log(err)
                            },
                            (postsList) => {
                                this.setState({posts: JSON.parse(postsList)})
                            }
                        )
                    }
                })
            })
        })
    }

    uploadPostImage() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.PhotosModule.uploadPostImage(
                            token,
                            email,
                            this.state.postMongoId,
                            this.state.postImage,
                            1,
                            this.state.placeId,
                            (err) => {
                                console.log(err)
                            },
                            (postUrl) => {
                                this.setState({postUrl: postUrl})
                            }
                        )
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
                            1,
                            this.state.placeId,
                            (err) => {
                                console.log(err)
                            },
                            (photoList) => {
                                this.setState({photoMap: photoList})
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View style={Style.view}>

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
                            onChangeText={(postTitle) => this.setState({postTitle})}/>

                        <TextInput
                            placeholder="Description"
                            underlineColorAndroid='transparent'
                            onChangeText={(postDescription) => this.setState({postDescription})}/>

                        <PhotoUpload onPhotoSelect={postImage => {
                            if (postImage) {
                                this.setState({postImage: postImage})
                            }
                        }
                        }>
                            <Image source={{image: this.state.postImage}}
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

                        <Button title="Add place post" onPress={() => this.createPlacePost()}>

                        </Button>
                    </ModalContent>
                </Modal>

                <PlaceHeader title={this.state.placeName} placeIdFromParent={this.state.placeId} isHome={false}
                             navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <Card style={Style.cardContainer}>
                        <Carousel
                            ref={(c) => {
                                this._carousel = c;
                            }}
                            data={this.state.images}
                            renderItem={this._renderItem}
                            sliderWidth={viewportWidth / 1.055}
                            itemWidth={viewportWidth}

                        />
                        <CardItem>
                            <CardTitle
                                title={this.state.placeName}
                                subtitle={this.state.city}
                            />
                        </CardItem>

                        <CardItem cardBody>
                            <Image source={this.state.img}
                                   style={{height: 20, width: null, flex: 1}}/>
                        </CardItem>

                        <CardItem>
                            <Body>
                                <Text>{this.state.description} </Text>
                            </Body>
                            <CardAction
                                separator={true}
                                inColumn={false}>
                            </CardAction>
                        </CardItem>
                    </Card>

                    {/*Posts about the place*/}
                    <Divider style={{backgroundColor: 'black'}}/>
                    <Text style={Style.heading}> Posts about {this.state.placeName}</Text>
                    <Divider style={{backgroundColor: 'black'}}/>

                    {this.state.posts.map((item, index) => {
                        return (
                            <Card style = {Style.cardContainer} key={this.state.posts.placePostId}>
                                <CardItem cardBody style={Style.cardContainer}>
                                    <Image source={{uri: this.state.photoMap[item.mongoId]}}
                                           style={{height: 200, width: null, flex: 1}}/>
                                </CardItem>
                                <CardItem>
                                    <CardTitle
                                        title={item.title}
                                        subtitle={item.creatorEmail}
                                    />
                                </CardItem>

                                <CardItem>
                                    <Body>
                                        <Text numberOfLines={1} ellipsizeMode={"tail"}>{item.body} </Text>
                                        <Text>{item.timeStamp} </Text>
                                    </Body>
                                </CardItem>
                            </Card>
                        )
                    })}
                </ScrollView>


                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Write a post about this place"
                                       onPress={() => this.setState({isVisible: true})}>
                        <Icon name="md-create" style={Style.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>

            </View>

        )
    }
}
