import React, { Component } from 'react'
import ImagePicker from 'react-native-image-crop-picker'
import {Text, View, FlatList, Image, StyleSheet, Dimensions, TouchableOpacity} from 'react-native';
// import styles from '../styles/Style'
import { ActionSheet, Root } from "native-base";
// import {STYLES} from "../constants/Styles";

const width = Dimensions.get('window').width

export default class LoadImage extends Component {
    constructor(props) {
        super(props);
        this.state= {
            fileList:[]
        }
    }
    onSelectedImage = (image) => {
        let newDataImg = this.state.fileList;
        const source = {uri: image.path};
        let item={
            id: Date.now(),
            url: source,
            content: image.data
        }
        newDataImg.push(item);
        this.setState({fileList: newDataImg})
    }
    takePhoto = () => {
        ImagePicker.openCamera({
            compressImageMaxWidth: 500,
            compressImageMaxHeight: 500,
            compressImageQuality: 0.7,
            includeBase64: true,
            cropping: false,
        }).then(image => {
            console.log('takePhoto ', image);
            this.onSelectedImage(image)

        });
    }
    selectPhotoFromGallery = () => {
        ImagePicker.openPicker({
            compressImageMaxWidth: 500,
            compressImageMaxHeight: 500,
            compressImageQuality: 0.7,
            includeBase64: true,
            cropping: false
        }).then(image => {
            console.log('selectPhotoFromGallery', image);
            this.onSelectedImage(image)
            console.log(image);
        });
    }

    onClickAddImage = () => {
        console.log("here")
        const BUTTONS = ['Take Photo', 'Choose Photo from Library', 'Cancel'];
        ActionSheet.show(
            {options: BUTTONS,
                cancelButtonIndex: 2,
                title:'Select a Photo'},
            buttonIndex => {
            switch (buttonIndex) {
                case 0:
                    this.takePhoto();
                    break;
                case 1:
                    this.selectPhotoFromGallery();
                    break;
                default:
                    break;
            }
            }
        )
    }

    renderItem = ({item, index}) => {
return(
    <View style={styles.itemViewImage}>
        <Image source={item.url} style={styles.itemImage}/>
    </View>

)
};
    render() {
        let{content, btnPressStyle, txtStyle} = styles;
        let{fileList} = this.state;
        return (
            <Root>
                <View style = {content}>
                    <FlatList
                        data={fileList}
                        renderItem={this.renderItem}
                        keyExtractor={(item, index) => index.toString()}
                        extraData={this.state}
                    />
                    <TouchableOpacity onPress={this.onClickAddImage} style={btnPressStyle}>
                        <Text style={txtStyle}> Press Add Image</Text>
                    </TouchableOpacity>
                </View>
            </Root>
        )
    }
}

const styles = StyleSheet.create({
        content:{
            flex: 1,
            alignItems:'center',
            marginTop:50,
            paddingLeft:30,
            paddingRight:30,
            marginBottom:30

        },
    btnPressStyle: {
        backgroundColor: '#007AFF',
        height: 50,
        width: width - 60,
        alignItems: 'center',
        justifyContent: 'center'
    },
    txtStyle: {
        color: '#ffffff'
    },
    itemImage: {
        backgroundColor: '#2f455c',
        height: 150,
        width:width - 60,
        borderRadius: 8,
        resizeMode: 'contain'
    },
    itemViewImage:{
        alignItems: 'center',
        borderRadius: 8,
        marginTop: 10
    }
}
)