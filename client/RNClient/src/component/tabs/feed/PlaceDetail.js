import React, {Component} from 'react';
import {Image, NativeModules, ScrollView, StyleSheet, Text, View} from "react-native";
import CustomHeader from "../../CustomHeader";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import {Body, CardItem, Icon} from "native-base";
import Carousel from "react-native-snap-carousel";
import ActionButton from "react-native-action-button";
import AsyncStorage from "@react-native-community/async-storage";

export default class PlaceDetail extends Component {
    constructor(props) {
        super(props);
        this.state = {
            placeId: 0,
            placeName: '',
            city: '',
            country: '',
            description: '',
            images:[
                {
                    url:'',
                    timestamp: ''
                }
            ]
        }
    }

    componentDidMount() {
        const placeId = this.props.navigation.getParam('placeId', '')
        const placeName = this.props.navigation.getParam('name', '')
        const city = this.props.navigation.getParam('city', '')
        const country = this.props.navigation.getParam('country', '')
        const description = this.props.navigation.getParam('description', '')

        this.setState({
            placeId,
            placeName,
            city,
            country,
            description
        })
        console.log("Id is " , placeId)
        this.getPlaceImages()
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
                                //console.log("image json", images)
                            })
                    }
                })
            })
        })
    }
    _renderItem = ({item, index}) => {
        console.log(item, index);
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
                <CustomHeader title={this.state.city} isHome={false} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                    <Card>
                        <CardItem>
                            <CardTitle
                                title={this.state.placeName}
                                subtitle={this.state.city}
                            />
                        </CardItem>

                        <CardItem cardBody>
                            <Image source={this.state.img}
                                   style={{height: 200, width: null, flex: 1}}/>
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
                    <ActionButton.Item buttonColor='#007AFF' title="Write a post about this place"
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
const styles = StyleSheet.create({
    actionButtonIcon: {
        fontSize: 20,
        height: 22,
        color: 'white',

    },
})
