// import React, {Component} from 'react';
// import {StyleSheet, View} from 'react-native';
// import ActionButton from 'react-native-action-button';
// import Icon from 'react-native-vector-icons/Ionicons';
//
// export default class PlusButton extends Component {
//     onClickListener = () => {
//         this.props.navigation.navigate('WritePost');
//     }
//
//     render() {
//
//         return (
//             <View style={{flex: 1, backgroundColor: '#f3f3f3'}}>
//                 {/* Rest of the app comes ABOVE the action button component !*/}
//                 <ActionButton buttonColor='#007AFF'>
//                     <ActionButton.Item buttonColor='#007AFF' title="Add new post"
//                                        onPress={() => this.onClickListener()}>
//                         <Icon name="md-create" style={styles.actionButtonIcon}/>
//                     </ActionButton.Item>
//                 </ActionButton>
//             </View>
//         );
//     }
//
// }
//
// const styles = StyleSheet.create({
//     actionButtonIcon: {
//         fontSize: 20,
//         height: 22,
//         color: 'white',
//
//     },
// });
