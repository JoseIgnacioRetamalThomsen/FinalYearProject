import React, {Component} from 'react';
import {View, TextInput} from 'react-native';
import CustomHeader from '../../headers/CustomHeader'
import PlusButton from "../../PlusButton";
import Place from '../../Place'
import DisplayPlace from "../../DisplayPlace";
import {ActionSheet, Root} from "native-base";
import LoadImage from "../../LoadImage";
class MyPosts extends Component {

    render() {
        const {navigation} = this.props;
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="My Posts" isHome={true} navigation={this.props.navigation}/>

                <View style={{flex: 1,  alignItems: 'center', justifyContent: 'center',  bottom:0,
                    }}>
                {/*<PlusButton navigation={navigation}/>*/}
                    <LoadImage></LoadImage>
                </View>

                    {/*<Place navigation = {this.props.navigation}/>*/}

            </View>
        );
    }
}

export default MyPosts
