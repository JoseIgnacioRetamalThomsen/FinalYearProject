import React, {Component} from 'react';
import {Image, NativeModules, ScrollView, StyleSheet, View} from 'react-native';
import {Body, CardItem, Icon, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import MapInput from "../../MapInput";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";


class CityDetail extends Component {
    constructor(props) {
        super(props);
        this.state = {
            cityId: 0,
            indexId: 0,
            city: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
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
        const city = this.props.navigation.getParam('city', '')
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
        console.log('componentDidMount' + cityId, city, img)
        this.getCityImages()
    }

    getCityImages() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    console.log("email/token in getCityImages " + email + " " + token)

                    if (token !== null) {

                        NativeModules.PhotosModule.getCityImage(
                            token,
                            email,
                            parseInt(this.state.cityId),
                            (err) => {
                                console.log("error In PhotoModule.getCityImage " + err)
                            },
                            (url) => {
                                console.log("Url is !!!", url)
                                this.setState({url: url})
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title={this.state.city} isHome={false} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <Card>
                        <CardItem>
                            <CardTitle
                                title={this.state.city}
                                subtitle={this.state.country}
                            />
                        </CardItem>

                        <CardItem cardBody>
                            <Image source={this.state.img}
                                   style={{height: 200, width: null, flex: 1}}/>
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
                    </Card>

                </ScrollView>
                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Write a post about this city"
                                       onPress={() => this.props.navigation.navigate('CreateCityPost', {
                                           indexId: this.state.indexId,
                                           city: this.state.city,
                                           country: this.state.country
                                       })}>
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
