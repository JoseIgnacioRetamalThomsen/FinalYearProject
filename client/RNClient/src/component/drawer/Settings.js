import React, {Component} from 'react';
import {Button, NativeModules, TextInput, View} from 'react-native';
import CustomHeader from '../CustomHeader'
import styles from "../../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";

class Settings extends Component {
    constructor(props) {
        super(props);
        this.state = {
            userId: 0,
            avatar_url: '',
            name: '',
            description: '',
        }
    }

    onClickListener() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in settings " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.updateUser(
                            value,
                            key,
                            key,
                            this.state.name,
                            this.state.description,
                            this.state.userId,
                            (err) => {
                                console.log("error In settings " + err)
                            },
                            (email, name, description, userId) => {
                                this.setState({name: name})
                                this.setState({description: description})
                                this.userId({description: userId})
                                console.log("successful!!!" + this.state.name, this.state.description)
                            })
                    }
                })
            })
        })
    }

    render() {

        return (

            <View style={{flex: 1}}>
                <CustomHeader title="Settings" navigation={this.props.navigation}/>

                <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>

                    <TextInput
                        style={styles.inputs}
                        placeholder="Name"
                        onChangeText={(name) => this.setState({name})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="Description"
                        onChangeText={(description) => this.setState({description})}/>
                    <Button title="Save changes"
                            onPress={() => this.onClickListener()}/>
                </View>
            </View>
        )
    }
}

export default Settings
