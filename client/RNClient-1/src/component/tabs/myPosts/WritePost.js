import React, {Component} from "react";
import {Text, View} from "react-native";
import CustomHeader from "../../CustomHeader";
import {Root} from "native-base";
import MapInput from "../../MapInput";

export default class WritePost extends Component {
    onClickEvent(){
        this.props.navigation.navigate('Post')
    }
    render() {
        return (
            <Root>
                <View style={{flex: 1}}>
                    <CustomHeader title="Write post" isHome={true} navigation={this.props.navigation}/>
                    <MapInput notifyChange={() => this.onClickEvent()} navigation = {this.props.navigation}/>
                    <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
                    </View>
                </View>
            </Root>
        )
    }
}
