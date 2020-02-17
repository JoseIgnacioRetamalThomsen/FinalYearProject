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

let key;

class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            avatar_url: null,
            image: null,
            email: key,
            name: 'Default Name',
            description: 'Default Description',
        }
    }

    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in profile " + key + " " + value)

                    if (value !== null) {
                        NativeModules.PhotosModule.getProfilePhoto(
                            key,
                            value,
                            (err) => {
                                console.log("In profile photo " + err)
                            },
                            (url) => {
                                this.setState({avatar_url: url})
                                console.log("successful photo getsss() " +this.state.avatar_url)
                            })
                        NativeModules.ProfilesModule.getUser(
                            value,
                            key,
                            (err) => {
                                console.log("In profile!!! " + err)
                            },
                            (name, description) => {
                                this.setState({name: name})
                                this.setState({description: description})
                                console.log("successful getUser")
                            })

                    }
                })
            })
        })
    }

    async updatePhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in profile " + key + " " + value)
                    console.log("this.state.image " + this.state.image)
                    if (value !== null) {
                        NativeModules.PhotosModule.uploadProfilePhoto(
                            key,
                            value,
                            this.state.image,

                            (err) => {
                                console.log("In uploadPhoto " + err)
                            },
                            (url) => {
                                this.setState({avatar_url: url})
                                console.log("avatar_url  is " + this.state.avatar_url)
                                console.log("successful upload!!!")
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
                                        console.log('Image base64 string: ', avatar)
                                    //(image)=>this.updatePhoto()

                                    // image => {
                                    //     if (image) {
                                    //         this.updatePhoto();
                                    //         console.log('Image base64 string: ', image)
                                    //     }
                                     }}
                                }>
                                    <Image source={{uri: this.state.avatar_url}}
                                           /*{this.state.avatar_url ? {uri: this.state.avatar_url } : IMAGE.ICON_DEFAULT_PROFILE}*/
                                           style={{
                                               height: 120,
                                               width: 120,
                                               borderRadius: 600,
                                               borderColor: 'black',
                                               borderWidth: 5,
                                               flex: 0,
                                               resizeMode: 'contain'
                                           }}/>
                                </PhotoUpload>
                                {/*<GeoLoc></GeoLoc>*/}
                            </View>
                        </Card>
                        <View>
                            <View>
                                <Text>Email {key} </Text>
                            </View>
                            <View>
                                <Text>Name {this.state.name} </Text>
                            </View>
                            <View>
                                <Text>Description {this.state.description} </Text>
                            </View>
                        </View>
                        <Button style={styles.buttonContainer} title="Edit Profile"
                                onPress={() => this.props.navigation.navigate("Settings")}></Button>
                    </View>

                </ScrollView>
            </View>
        );
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
