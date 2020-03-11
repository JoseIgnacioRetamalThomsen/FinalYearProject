// import React, {Component} from 'react';
// import {Button, Image, NativeModules, ScrollView, StyleSheet, View} from 'react-native';
// import AsyncStorage from "@react-native-community/async-storage";
// import {Body, CardItem, Icon, Text} from "native-base";
// import {Card, CardAction, CardButton, CardContent, CardImage, CardTitle} from "react-native-material-cards";
// import CustomHeader from "../../CustomHeader";
// import MapInput from "../../MapInput";
// import ActionButton from "react-native-action-button";
// import CreateCity from "./CreateCity";
// import CityDetail from "./CityDetail";
// import GeoLoc from "../../GeoLoc"
//
// export default class DisplayPlaces extends Component {
//     constructor(props) {
//         super(props);
//         this.state = {}
//     }
//
//     componentDidMount() {
//         AsyncStorage.getAllKeys((err, keys) => {
//             AsyncStorage.multiGet(keys, (err, stores) => {
//                 stores.map((result, i, store) => {
//                     let email = store[i][0];
//                     let token = store[i][1]
//
//                     if (value != null) {
//                         NativeModules.ProfilesModule.getAllPlaces(
//                             this.state.max,
//                             (err) => {
//                                 console.log(err)
//                             },
//                             (jsonPlaceList) => {
//                                 this.setState({cities: JSON.parse(jsonPlaceList)})
//                             })
//                     }
//                 })
//             })
//         })
//     }
//
//
//     render() {
//         return (
//             <View style={{flex: 1}}>
//                 {/*<GeoLoc parentCallback={this.callbackFunction}/>*/}
//                 <CustomHeader title="Posts" isHome={true} navigation={this.props.navigation}/>
//                 <ScrollView style={{flex: 1}}>
//                     <View style={{flex: 1}}>
//                         {/*<MapInput navigation={this.props.navigation} notifyChange={() => this.onClickEvent()}*/}
//                         {/*          parentCallback={this.callbackFunction}/>*/}
//                     </View>
//                     {this.state.cities.map((item, index) => {
//                         return (
//                             <Card key={this.state.cities.cityId}>
//
//                                 <CardItem>
//                                     <CardTitle
//                                         title={item.name}
//                                         subtitle={item.country}
//                                     />
//                                 </CardItem>
//
//                                 <CardItem>
//                                     <Body>
//                                         <Text numberOfLines={1} ellipsizeMode={"tail"}>{item.description} </Text>
//                                     </Body>
//                                     <CardAction
//                                         separator={true}
//                                         inColumn={false}>
//                                         <CardButton
//                                             onPress={() => this.props.navigation.navigate('CityDetail', {
//                                                 cityId: item.cityId,
//                                                 indexId: item.indexId,
//                                                 name: item.name,
//                                                 country: item.country,
//                                                 description: item.description,
//                                                 //img: item.img
//                                             })}
//                                             title="More"
//                                             color="blue"
//                                         />
//                                     </CardAction>
//                                 </CardItem>
//                             </Card>
//                         )
//                     })}
//                 </ScrollView>
//                 <ActionButton buttonColor='#007AFF'>
//                     <ActionButton.Item buttonColor='#007AFF' title="Add a city"
//                                        onPress={() => this.props.navigation.navigate('CreateCity')}>
//                         <Icon name="md-create" style={styles.actionButtonIcon}/>
//                     </ActionButton.Item>
//                 </ActionButton>
//             </View>
//
//         );
//     }
// }
// // const styles = StyleSheet.create({
// //     actionButtonIcon: {
// //         fontSize: 20,
// //         height: 22,
// //         color: 'white',
// //     },
// // })
