import React, {Component} from 'react';
import {View, TextInput} from 'react-native';
import CustomHeader from '../../CustomHeader'
import PlusButton from "../../PlusButton";
import Place from '../../Place'
import DisplayPlace from "../../DisplayPlace";
import {ActionSheet, Root} from "native-base";
class MyPosts extends Component {

    render() {
        const {navigation} = this.props;
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="My Posts" isHome={true} navigation={this.props.navigation}/>

                {/*<View style={{flex: 1, justifyContent: 'center'}}>*/}
                {/*<PlusButton navigation={navigation}/>*/}
                {/*</View>*/}

                    <Place navigation = {this.props.navigation}/>
            </View>
        );
    }
}

export default MyPosts
