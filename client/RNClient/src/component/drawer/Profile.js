import React, {Component} from 'react'
import {
    View,
    Image,
    Button,
    StyleSheet,
    ImageBackground,
    ScrollView,
    NativeModules,
} from 'react-native'
import {Text} from 'native-base'
import CustomHeader from '../CustomHeader'
import GeoLoc from "../GeoLoc";
import {IMAGE} from "../../constants/Image";
import PhotoUpload from "react-native-photo-upload"
import {Card} from 'react-native-elements'
import Settings from "./Settings"
import AsyncStorage from "@react-native-community/async-storage"

let email

class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            avatar_url: null,
            image: '../../img/profile.png',
            email: email,
            name: '',
            description: '',
            message: ''
        }
    }

    componentDidMount() {
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
                                if (url == null) {
                                    this.setState({avatar_url: 'https://storage.googleapis.com/wcity-images-1/profile-1/1896665_1780468.jpg'})
                                } else {
                                    this.setState({avatar_url: url})
                                }

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
                                console.log("json", json)
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

    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="Profile" navigation={this.props.navigation}/>
                <ScrollView style={styles.scroll}>
                    <ImageBackground
                        style={styles.headerBackgroundImage}
                        blurRadius={10}
                        source={{
                            //uri: avatarBackground,
                        }}
                    />
                    <View style={styles.headerContainer}>
                        <Card containerStyle={styles.cardContainer}>
                            <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
                                <PhotoUpload onPhotoSelect={avatar => {
                                    if (avatar) {
                                        this.setState({image: avatar})
                                        this.updatePhoto()
                                    }
                                }
                                }>
                                    <Image source={{uri: this.state.avatar_url}}
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
                                {/*<GeoLoc></GeoLoc>*/}
                            </View>
                        </Card>
                        <View>
                            <View>
                                <Text>Email {email} </Text>
                            </View>
                            <View>
                                <Text>Name {this.state.name} </Text>
                            </View>
                            <View>
                                <Text>Description {this.state.description} </Text>
                            </View>
                        </View>
                        <Button style={styles.buttonContainer} title="Edit Profile"
                                onPress={() => this.props.navigation.navigate("Settings")}/>
                    </View>
                </ScrollView>
            </View>
        )
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
