import React,  {Component} from 'react';
import { View, TextInput } from 'react-native';
import {Button, Root, Text} from 'native-base';
import CustomHeader from '../../CustomHeader'
import LoadImage from '../../LoadImage'

import API from "./API";
import PlusButton from "../../PlusButton";
import Post from "../../Post";


class MyPosts extends Component {

    render() {
      return (
        <View style={{ flex: 1 }}>
          <CustomHeader  title="My Posts"  isHome={true} navigation={this.props.navigation}/>

        <View style={{ flex: 1, justifyContent: 'center' }}>
            {/*<Post/>*/}
            <PlusButton/>

        </View>

        </View>
      );
    }
  }

export default MyPosts
