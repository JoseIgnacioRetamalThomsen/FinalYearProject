import React, {Component} from 'react';
import {Header, Left, Body, Right, Button, Icon, Title} from 'native-base';
import {NativeModules, Text, View} from "react-native";
import AsyncStorage from "@react-native-community/async-storage";

class PlaceHeader extends Component {
    constructor(props) {
        super(props);
        this.state = {
            placeId: -999,
            pressed: false
        }
        this.setState({
            placeId: this.props.placeIdFromParent
        })

    }

    componentDidMount() {
        this.setState({pressed: false})
    }


    render() {
        let {title, isHome} = this.props;

        if (global.visitedPlaceMap[this.props.placeIdFromParent] === true)
            return (
                <Header style={{backgroundColor: '#007AFF'}}>
                    <Left>
                        {
                            isHome ?
                                <Button transparent onPress={() => this.props.navigation.openDrawer()}>
                                    <Icon name='menu'/>
                                </Button> :
                                <Button transparent onPress={() =>this.props.navigation.navigate('DisplayCities')}>
                                    <Icon name='arrow-back'/>
                                </Button>
                        }
                    </Left>
                    <Body>
                        <Title>{title}</Title>
                    </Body>

                </Header>
            )
        else
            return (
                <Header style={{backgroundColor: '#007AFF'}}>
                    <Left>
                        {
                            isHome ?
                                <Button transparent onPress={() => this.props.navigation.openDrawer()}>
                                    <Icon name='menu'/>
                                </Button> :
                                <Button transparent onPress={() =>this.props.navigation.navigate('CityDetail')}>
                                    <Icon name='arrow-back'/>
                                </Button>
                        }
                    </Left>
                    <Body>
                        <Title>{title}</Title>
                    </Body>
                    <Right>
                        <Button hasText transparent onPress={() => this.visitPlace()}>
                            <Text> Visited this Place </Text>
                        </Button>
                    </Right>
                </Header>
            )
    }

    visitPlace() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.ProfilesModule.visitPlace(
                            token,
                            email,
                            parseFloat(this.props.placeIdFromParent),
                            (err) => {
                                console.log(err)
                            },

                            (isValid) => {
                                global.visitedPlaceMap[this.props.placeIdFromParent] = true
                                this.setState({pressed:true})
                            })
                    }
                })
            })
        })
    }

}

export default PlaceHeader
