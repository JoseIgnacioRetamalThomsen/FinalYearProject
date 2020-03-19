import React, {Component} from 'react';
import {Header, Left, Body, Right, Button, Icon, Title} from 'native-base';
import {NativeModules, Text, View} from "react-native";
import AsyncStorage from "@react-native-community/async-storage";

class CustomHeader extends Component {
    constructor(props) {
        super(props);
        this.state = {
            cityId: -999,
            pressed: false
        }
        this.setState({
            cityId: this.props.cityIdFromParent
        })

    }

    componentDidMount() {
        this.setState({pressed: false})
    }

    returnData(pressed) {
        this.setState({pressed: pressed});
    }
    render() {
        let {title, isHome} = this.props;

        if (this.state.pressed === false)
            return (
                <Header style={{backgroundColor: '#007AFF'}}>
                    <Left>
                        {
                            isHome ?
                                <Button transparent onPress={() => this.props.navigation.openDrawer()}>
                                    <Icon name='menu'/>
                                </Button> :
                                <Button transparent onPress={() =>this.props.navigation.navigate('DisplayCities',
                                    {returnData: this.returnData.bind(this)})}>
                                    <Icon name='arrow-back'/>
                                </Button>
                        }
                    </Left>
                    <Body>
                        <Title>{title}</Title>
                    </Body>
                    <Right>
                        <Button hasText transparent onPress={() => this.visitCity()}>
                            <Text> Visited </Text>
                        </Button>
                    </Right>
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
                                <Button transparent onPress={() =>this.props.navigation.navigate('DisplayCities',
                                    {returnData: this.returnData.bind(this)})}>
                                    <Icon name='arrow-back'/>
                                </Button>
                        }
                    </Left>
                    <Body>
                        <Title>{title}</Title>
                    </Body>
                    <Right>
                    </Right>
                </Header>
            )
    }

    visitCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.ProfilesModule.visitCity(
                            token,
                            email,
                            parseFloat(this.props.cityIdFromParent),
                            (err) => {
                                console.log(err)
                            },

                            (timestamp) => {
                                this.setState({pressed: true})
                                console.log("this.state.pressed should be true", this.state.pressed)
                            })
                    }
                })
            })
        })
    }

}

export default CustomHeader
