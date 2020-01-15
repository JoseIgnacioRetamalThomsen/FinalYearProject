import React, { Component } from 'react';
import { View, Image, SafeAreaView, ScrollView } from 'react-native';
import { Text, List, ListItem } from 'native-base';
import { IMAGE } from '../constants/Image'

class SideMenu extends Component {
    render() {
        return (
            <SafeAreaView style={{ flex: 1 }}>
                <View style={{ height: 150, alignItems: 'center', justifyContent: 'center' }}>
                    <Image source={IMAGE.ICON_USER_DEFAULT} style={{ height: 120, borderRadius: 60 }} />
                </View>
                <ScrollView>
                    <List>
                        <ListItem onPress={() => this.props.navigation.navigate('Profile')}>
                            <Text>
                                Profile
                            </Text>
                        </ListItem>
                        <ListItem onPress={() => this.props.navigation.navigate('Settings')}>
                            <Text>
                                Settings
                            </Text>
                        </ListItem>
                        <ListItem noBorder>
                            <Text onPress={() => this.props.navigation.navigate('auth')}>
                                Log out
                            </Text>
                        </ListItem>
                    </List>
                </ScrollView>
            </SafeAreaView>
        )

    }
}
export default SideMenu