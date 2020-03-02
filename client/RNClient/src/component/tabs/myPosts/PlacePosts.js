import React, {Component} from 'react';
import {View, Image, ScrollView, StyleSheet, NativeModules} from 'react-native';
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

class PlacePosts extends Component {
    constructor(props) {
        super(props)
        this.state = {
            city: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
            id: 0,
        }
    }
    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in displayCity " + key + " " + value)

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
                                console.log("error In ProfilesModule.getCity " + err)
                            },
                            (city, country, email, description, lat, lon, id) => {
                                this.setState({city: city})
                                this.setState({country: country})
                                this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                this.setState({id: id})
                                console.log("successful values in getCity!!!" + this.state.city, this.state.country, this.state.description)
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
    goToWritePost = () => {
        this.props.navigation.navigate('WriteCityPost');
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
                <CustomHeader title="DisplayCityPosts" isHome={true} navigation={this.props.navigation}/>
                <ScrollView style={{flex:1}}>
                    <View style={{flex: 1}}>
                        <MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()}
                                  parentCallback={this.callbackFunction}/>
                    </View>
                    {/*<Text> {this.state.lat.toFixed(2)} {this.state.lng.toFixed(2)} {this.state.city} {this.state.country}</Text>*/}
                    {/*<View style={{flex: 1}}>*/}

                    <Card>
                        <CardItem>
                            <CardTitle
                                title="GMIT"
                                subtitle="Institute"
                            />
                        </CardItem>

                        <CardItem cardBody>
                            <Image source={require('../../../img/gmit.jpg')}
                                   style={{height: 200, width: null, flex: 1}}/>
                        </CardItem>
                        <CardItem>
                            <Body>
                                <Text>Galway-Mayo Institute of Technology (Irish: Institúid Teicneolaíochta na
                                    Gaillimhe-Maigh Eo)
                                    is a third level institute of education and is based at five locations in the west
                                    of
                                    Ireland. </Text>
                            </Body>
                            <CardAction
                                separator={true}
                                inColumn={false}>
                                <CardButton
                                    onPress={() => {
                                    }}
                                    title="More"
                                    color="blue"
                                />
                            </CardAction>
                        </CardItem>
                    </Card>
                    <Card>
                        <CardItem>
                            <CardTitle
                                title="NUIG"
                                subtitle="University"
                            />
                        </CardItem>

                        <CardItem cardBody>
                            <Image source={require('../../../img/nuig.jpg')}
                                   style={{height: 200, width: null, flex: 1}}/>
                        </CardItem>
                        <CardItem>
                            <Body>
                                <Text>The National University of Ireland Galway (NUI Galway; Irish: OÉ Gaillimh) is located in the city of Galway in Ireland.
                                    A third-level teaching and research institution, the University has been awarded the full five QS stars for excellence,
                                    and is ranked among the top 1 percent of universities according to the 2018 QS World University Rankings.[3] </Text>
                            </Body>
                            <CardAction
                                separator={true}
                                inColumn={false}>
                                <CardButton
                                    onPress={() => {
                                    }}
                                    title="More"
                                    color="blue"
                                />
                            </CardAction>
                        </CardItem>
                    </Card>

                </ScrollView>

                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Add new post"
                                       onPress={() => this.goToWritePost()}>
                        <Icon name="md-create" style={styles.actionButtonIcon}/>
                    </ActionButton.Item>
                </ActionButton>


            </View>

        )
    }
}

export default PlacePosts
const styles = StyleSheet.create({
    actionButtonIcon: {
        fontSize: 20,
        height: 22,
        color: 'white',

    },
});
