import React, {Component} from "react";
import {default as Alert, Text, View} from "react-native";

export default class NotFoundCity extends Component {
    constructor(props) {
        super(props);
        this.state = {}
    }
    showAlert1() {
        Alert.alert(
            'Alert Title',
            'My Alert Msg',
            [
                {
                    text: 'Cancel',
                    onPress: () => console.log('Cancel Pressed'),
                    style: 'cancel',
                },
                {text: 'OK', onPress: () => console.log('OK Pressed')},
            ]
        );
    }
    render()
{
    return (
        <View>
            <Text>
                City is not created yet! Would you like to create a city?
            </Text>
        </View>
    )
}
}
