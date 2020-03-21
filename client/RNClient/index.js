/**
 * @format
 */
import 'react-native-gesture-handler'
import {AppRegistry} from 'react-native';
import App from './App';
import {name as appName} from './app.json';
import './src/component/global'

AppRegistry.registerComponent(appName, () => App);
//https://stackoverflow.com/a/52314596
// import {AppRegistry} from 'react-native';
// import App from './App';
// import React from 'react';
// import {name as appName} from './app.json';
// import {Provider} from 'react-redux';
// import {applyMiddleware, createStore} from "redux";
// import allReducers from './src/reducers'
// import {createLogger} from 'redux-logger'
//
// const logger = createLogger()
// const initialState = {};
// const store = createStore(
//     allReducers, initialState, applyMiddleware(logger)
// )
//
// const Root = () => (
//     <Provider store={store}>
//         <App />
//     </Provider>
// )
//
// AppRegistry.registerComponent(appName, () => Root);
