import {Dimensions, StyleSheet} from 'react-native'
import {constants} from '../constants/Constants'

const {width: width} = Dimensions.get('window')
export default StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#DCDCDC'
    },
    view: {
        flex: 1,
        backgroundColor: '#DCDCDC',
        marginBottom:10
    },
    createContainer: {
        flex: 1,
        padding: 10,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#DCDCDC',
    },
    inputContainer: {
        borderBottomColor: '#F5FCFF',
        backgroundColor: '#FFFFFF',
        borderRadius: 30,
        borderBottomWidth: 1,
        width: 250,
        height: 45,
        marginBottom: 20,
        flexDirection: 'row',
        alignItems: 'center'
    },
    createInputContainer: {
        flex: 1,
        borderBottomColor: '#F5FCFF',
        backgroundColor: '#FFFFFF',
        borderRadius: 30,
        borderBottomWidth: 1,
        width: 350,
        maxHeight: 50,
        marginBottom: 20,
        flexDirection: 'row',
        alignItems: 'center'
    },
    descInputContainer: {
        flex: 1,
        borderBottomColor: '#F5FCFF',
        backgroundColor: '#FFFFFF',
        borderRadius: 30,
        borderBottomWidth: 1,
        width: 350,
        maxHeight: 50,
        marginBottom: 20,
        flexDirection: 'row',
        alignItems: 'center',
        textAlign: 'center'
    },
    inputs: {
        height: 45,
        marginLeft: 16,
        borderBottomColor: '#FFFFFF',
        flex: 1,
    },
    inputIcon: {
        width: 30,
        height: 30,
        marginLeft: 15,
        justifyContent: 'center'
    },
    buttonContainer: {
        height: 45,
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 20,
        width: 250,
        borderRadius: 30,
    },
    loginButton: {
        backgroundColor: "#007AFF",
    },
    loginText: {
        color: 'white',
    },
    content: {
        flex: 1,
        alignItems: 'center',
        marginTop: 50,
        paddingLeft: 30,
        paddingRight: 30,
        marginBottom: 30

    },
    btnPressStyle: {
        backgroundColor: '#007AFF',
        height: 50,
        width: constants.width - 60,
        alignItems: 'center',
        justifyContent: 'center',
        borderRadius: 10
    },
    modalbtn: {
        backgroundColor: '#007AFF',
        height: 50,
        width: constants.width / 1.45,
        alignItems: 'center',
        justifyContent: 'center',
        borderRadius: 10,
        borderColor: 'black'
    },
    txtStyle: {
        color: '#ffffff'
    },
    itemImage: {
        backgroundColor: '#2f455c',
        height: 150,
        width: constants.width - 60,
        borderRadius: 8,
        resizeMode: 'contain'
    },
    touchableButton: {
        position: 'absolute',
        right: 3,
        height: 40,
        width: 35,
        padding: 2
    },
    buttonImage: {
        resizeMode: 'contain',
        height: '100%',
        width: '100%',
    },
    cardContainer: {
        borderColor: '#0080ff',
        backgroundColor: '#FFF',
        borderWidth: 1,
        flex: 1,
        margin: 10,
        padding: 0,
        marginBottom: 25,
    },
    carouselContainer: {
        flexDirection: 'row',
        backgroundColor: '#FFF',
        borderWidth: 1,
        flex: 1,
        margin: 10,
        padding: 0,
        width: width / 1.09,
    },
    noPostsPhoto: {
        height: 350,
        width: width / 1.095,
        flex: 1,

    },
    cardPhoto: {
        height: 200,
        width: width / 1.095,
        flex: 1
    },
    profilePhoto: {
        height: 120,
        width: 120,
        borderRadius: 60,
        margin: 10,
        padding: 10,
        marginTop: 30,
        borderColor: '#0080ff',
        borderWidth: 0.5,
        flex: 0,
        resizeMode: 'cover',//contain?
        justifyContent: 'center',
        alignItems: 'center',
    },
    uploadPhoto: {
        height: 150,
        width: 150,
        margin: 10,
        padding: 10,
        marginTop: 40,
        marginBottom: 50,
        borderColor: '#0080ff',
        borderRadius: 10,
        borderWidth: 0.5,
        flex: 0,
        resizeMode: 'cover',//contain?
        justifyContent: 'center',
        alignItems: 'center',
    },
    actionButtonIcon: {
        fontSize: 20,
        height: 22,
        color: 'white',
    },
    slide: {},
    heading: {
        textAlign: 'left',
        margin: 10,
        fontWeight: 'bold',
        fontFamily: 'sans-serif',
        fontSize: 25
    },
    title: {
        textAlign: 'left',
        margin: 10,
        fontWeight: 'bold',
    },
    text: {
        textAlign: 'left',
        marginLeft: 10,
        marginBottom: 10,
    },
    modal: {},
    modalContent: {
        borderWidth: 1,
        borderColor: '#0080ff',
        backgroundColor: '#FFF',
        width: Dimensions.get('window').width * 0.8, height: Dimensions.get('window').height * 0.6,

    }
})
