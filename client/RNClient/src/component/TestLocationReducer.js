// import {bindActionCreators} from "redux";
// import {connect} from "react-redux";
// import React from "react"
// import {Text, View, FlatList} from "react-native";
// import {selectLocation} from '../actions/index'
//
// class TestLocationReducer extends React.Component {
//
//     createListItems() {
//         return this.props.location.map((location) => {
//             return (
//                 <Text
//                     key={location.id}
//                     onClick={() => this.props.selectLocation(location)}>
//                     {location.lat} {location.lng}
//                 </Text>
//             )
//         })
//     }
//
//     render() {
//         return (
//             <View>
//                {this.createListItems()}
//             </View>
//         )
//     }
// }
//
// function mapStateToProps(state) {
//     return {
//         location: state.location
//     }
//
// }
//
// function mapDispatchToProps(dispatch) {//dispatch - call a func
//     return bindActionCreators({selectLocation: selectLocation}, dispatch)
// }
// export default connect(mapStateToProps)(TestLocationReducer)
