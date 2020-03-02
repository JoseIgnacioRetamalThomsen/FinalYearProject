import React, { Component } from 'react';
import {Image, ScrollView, StyleSheet, View} from 'react-native';
import {Body, CardItem, Icon, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import MapInput from "../../MapInput";
import {Card, CardAction, CardButton, CardTitle} from "react-native-material-cards";
import ActionButton from "react-native-action-button";
import  WriteCityPost from './WriteCityPost'

class CityDetail extends Component {
    constructor(props) {
        super(props);
        this.state = {
            indexId: 0,
            city: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
        }
    }
    componentDidMount() {
        const indexId = this.props.navigation.getParam('indexId', 0)
        const city = this.props.navigation.getParam('city', '')
        const country = this.props.navigation.getParam('country', '')
        const description = this.props.navigation.getParam('description', '')
        const img = this.props.navigation.getParam('img', '')

        this.setState({
            indexId,
            city,
            country,
            description,
            img
        })
        console.log('componentDidMount'+indexId, city, description)
    }

    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title={this.state.city} isHome={false} navigation={this.props.navigation}/>
                <ScrollView style={{flex: 1}}>
                            <Card >
                                <CardItem>
                                    <CardTitle
                                        title={this.state.city}
                                        subtitle={this.state.country}
                                    />
                                </CardItem>

                                <CardItem cardBody>
                                    <Image source={require('../../../img/noImage.png')}
                                           style={{height: 200, width: null, flex: 1}}/>
                                </CardItem>
                                <CardItem>
                                    <Body>
                                        <Text >{this.state.description} </Text>
                                    </Body>
                                    <CardAction
                                        separator={true}
                                        inColumn={false}>
                                        <CardButton
                                            onPress={() => this.props.navigation.navigate('CityDetail')}
                                            title="Edit"
                                            color="blue"
                                        />
                                    </CardAction>
                                </CardItem>
                            </Card>


                </ScrollView>
                <ActionButton buttonColor='#007AFF'>
                    <ActionButton.Item buttonColor='#007AFF' title="Write a post about this city"
                                       onPress={() => this.props.navigation.navigate('WriteCityPost')}>
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
