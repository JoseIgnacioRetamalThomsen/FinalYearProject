import React, {Component} from 'react';
import {View, TextArea} from 'react-native';
import {Button, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import MapInput from "../../MapInput";

class Feed extends Component {
    onClickEvent(){
        // this.props.navigation.navigate('Post')
        alert("called")
    }
    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="Feed" isHome={true} navigation={this.props.navigation}/>
                <MapInput notifyChange={() => this.onClickEvent()} />
                <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
                </View>
            </View>
        );
    }
}

export default Feed