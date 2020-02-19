import React, {Component} from 'react';
import {View, TextArea} from 'react-native';
import {Button, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import City from '../../City'
import DisplayCity from '../../DisplayCity'
import MapInput from "../../MapInput";
class Feed extends Component {
    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="Feed" isHome={true} navigation={this.props.navigation}/>
                <MapInput notifyChange={() => this.onClickEvent()} />
                {/*<View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>*/}
                    <City/>
                {/*</View>*/}
            </View>

        );
    }
}

export default Feed
