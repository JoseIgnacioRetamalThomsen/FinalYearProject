import {combineReducers} from "redux"
import LocationReducer from "./reducer-location"

const allReducers = combineReducers({
      location: LocationReducer,

 })
export default allReducers
