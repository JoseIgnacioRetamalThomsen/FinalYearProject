import React, {Component} from 'react';
import {Header, Left, Body, Right, Button, Icon, Title} from 'native-base';
import {NativeModules, Text, View} from "react-native";
import AsyncStorage from "@react-native-community/async-storage";

class SpecialHeader extends Component {
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
    render() {//global.visitedCityMap[this.props.cityIdFromParent]
        let {title, isHome} = this.props;
        console.log("this.state.cityId", global.visitedCityMap[this.props.cityIdFromParent])
        if (global.visitedCityMap[this.props.cityIdFromParent] === true)

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
                        <Button hasText transparent onPress={() => this.visitCity()}>
                            <Text> Visited </Text>
                        </Button>
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

                            (isValid) => {
                                global.visitedCityMap[this.props.cityIdFromParent] = true
                                this.setState({pressed:true})
                            })
                    }
                })
            })
        })
    }

}

export default SpecialHeader
