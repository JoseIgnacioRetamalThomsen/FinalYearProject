import React from "react";
import {Button, View} from "react-native";

export default class Test extends React.Component {
    print(){
        console.log("Clicked");
    }
    render() {
        return (
            <View>
                <Button onPress={this.print()}>Hello </Button>
            </View>
        )
    }
}
