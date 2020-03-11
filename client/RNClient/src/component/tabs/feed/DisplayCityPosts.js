import React, {Component} from 'react';
import {View, Image, ScrollView, FlatList, StyleSheet, NativeModules} from 'react-native';
import CustomHeader from '../../CustomHeader'
import MapInput from "../../MapInput";
import {GooglePlacesAutocomplete} from "react-native-google-places-autocomplete";
import {
    Card,
    CardItem,
    Text,
    Icon,
    Body,
} from 'native-base';
import {CardAction, CardButton, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";
import {CardList} from 'react-native-card-list';

class DisplayCityPosts extends Component {
    constructor(props) {
        super(props)
        this.state = {
            posts: [
                // {
                //     valid: true,
                //     indexId: 0,
                //     index: 0,
                //     creatorEmail: '',
                //     cityName: 'Galway',
                //     cityCountry: 'Ireland',
                //     title: '',
                //     body: 'CreateCity in the west',
                //     img: '../../../img/gmit.jpg',
                //     timeStamp: '',
                //     likes: [],
                //     mongoId: 0,
                //     lat: 0,
                //     lon: 0,
                //     id: 0,
                // },
                {
                    indexId: 0,
                    city: 'GMIT',
                    country: 'Ireland',
                    email: '',
                    description: 'CreateCity in the east',
                    img: '../../../img/nuig.jpg',
                    lat: 0,
                    lon: 0,
                    id: 0,
                },
                {
                    indexId: 1,
                    city: 'NUIG',
                    country: 'Ireland',
                    email: '',
                    description: 'CreateCity in the south-west CreateCity in the south-west CreateCity in the south-west CreateCity in the south-westCity in the south-west CreateCity in the south-west' +
                        'CreateCity in the south-westCity in the south-westCity in the south-westCity in the south-westCity in the south-westCity in the south-westCity in the south-west' +
                        'CreateCity in the south-westCity in the south-westCity in the south-westCity in the south-westCity in the south-westCity in the south-west' +
                        'CreateCity in the south-westCity in the south-westCity in the south-westv',
                    img: '../../../img/noImage.png',
                    lat: 0,
                    lon: 0,
                    id: 0,
                },
                {
                    indexId: 1,
                    city: 'GTI',
                    country: 'Ireland',
                    email: '',
                    description: 'CreateCity in the west',
                    img: '../../../img/noImage.png',
                    lat: 0,
                    lon: 0,
                    id: 0,
                },
            ]

        }
    }

    // componentDidMount() {
    //     AsyncStorage.getAllKeys((err, keys) => {
    //         AsyncStorage.multiGet(keys, (err, stores) => {
    //             stores.map((result, i, store) => {
    //                 let key = store[i][0];
    //                 let value = store[i][1]
    //                 console.log("key/value in displayCity " + key + " " + value)
    //
    //                 if (value !== null) {
    //                     NativeModules.PostModule.getCityPosts(
    //                         this.state.indexId,
    //                         (err) => {
    //                             console.log("error In PostsModule.getCityPosts " + err)
    //                         },
    //                         (valid, indexId, index, creatorEmail, cityName, cityCountry, title, body, timeStamp, likes, mongoId) => {
    //                             //this.setState({valid: valid})
    //                            // this.setState({indexId: indexId})
    //                            // this.setState({index: index})
    //                            // this.setState({creatorEmail: creatorEmail})
    //                            //  this.setState({cityName: cityName})
    //                            //  this.setState({cityCountry: cityCountry})
    //                            //  this.setState({title: title})
    //                            //  this.setState({body: body})
    //                            // this.setState({timeStamp: timeStamp})
    //                            // this.setState({likes: likes})
    //                            // this.setState({mongoId: mongoId})
    //                             //console.log("successful values in getCityPosts!!!" + this.state.indexId, this.state.cityName, this.state.cityCountry)
    //                         })
    //                 }
    //             })
    //         })
    //     })
    // }
    getCityPosts() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getCity() " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getCity(
                            value,
                            this.state.city,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,

                            (err) => {
                                console.log("err in getCity " + err)
                            },
                            (valid, name, country, email, description, lat, lon, id) => {
                                this.setState({name: name})
                                this.setState({country: country})
                                // this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                console.log("name id  is " + name, id)
                                console.log("successfully got a city!!!")
                                //this.props.navigation.navigate('DisplayCities')
                            })
                    }
                })
            })
        })
    }
    updateCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in updatecity " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.updateCity(
                            value,
                            this.state.name,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,
                            (err) => {
                                console.log("err in updateCity " + err)
                            },
                            (name, country, email, description, lat, lon, id) => {
                                this.setState({name: name})
                                this.setState({country: country})
                                // this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                console.log("name, country, email, description, lat, lon id in updateCity is " + name, country, email, description, lat, lon, id)
                                console.log("successfully updated a city!!!")
                            })
                    }
                })
            })
        })
    }

    callbackFunction = (lat, lng, city, country) => {
        this.setState({lat: lat})
        this.setState({lng: lng})
        this.setState({city: city})
        this.setState({country: country})
    }

    render() {
        return (
            <View style={{flex: 1}}>

                <CustomHeader title={this.state.city +" Posts"} isHome={false} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <View style={{flex: 1}}>
                        <MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()}
                                  parentCallback={this.callbackFunction}/>
                    </View>
                    { this.state.posts.map((item, index) => {
                        return (
                            <Card key={this.state.posts.city}>
                                <CardItem>
                                    <CardTitle
                                        title={item.city}
                                        subtitle={item.country}
                                    />
                                </CardItem>

                                <CardItem cardBody>
                                    <Image source={require('../../../img/noImage.png')}
                                           style={{height: 200, width: null, flex: 1}}/>
                                </CardItem>
                                <CardItem>
                                    <Body>
                                        <Text numberOfLines={1} ellipsizeMode={"tail"}>{item.description} </Text>
                                    </Body>
                                    <CardAction
                                        separator={true}
                                        inColumn={false}>
                                        <CardButton
                                            onPress={() => this.props.navigation.navigate('FeedDetail')}
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
                    <ActionButton.Item buttonColor='#007AFF' title="Add new post"
                                       onPress={() => this.props.navigation.navigate('WriteCityPost')}>
                        <Icon name="md-create" style={styles.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>
            </View>

        )
    }
}

export default DisplayCityPosts
const styles = StyleSheet.create({
    actionButtonIcon: {
        fontSize: 20,
        height: 22,
        color: 'white',
    },
});
