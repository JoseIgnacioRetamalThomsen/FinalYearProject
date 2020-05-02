// import ImagePicker from "react-native-image-crop-picker";
// import {ActionSheet} from "native-base";
//
// function onSelectedImage (image) {
//     let newDataImg = this.state.fileList;
//     const source = {uri: image.path};
//     let item = {
//         id: Date.now(),
//         url: source,
//         content: image.data
//     }
//     newDataImg.push(item);
//     this.setState({fileList: newDataImg})
//     this.setState({ defaultImage: false })
// }
//
// function takePhoto() {
//     ImagePicker.openCamera({
//         compressImageMaxWidth: 500,
//         compressImageMaxHeight: 500,
//         compressImageQuality: 0.7,
//         includeBase64: true,
//         cropping: false,
//     }).then(image => {
//         console.log('takePhoto ', image);
//         onSelectedImage(image)
//         this.setState({ defaultImage: false })
//     });
// }
// function selectPhotoFromGallery() {
//     ImagePicker.openPicker({
//         compressImageMaxWidth: 500,
//         compressImageMaxHeight: 500,
//         compressImageQuality: 0.7,
//         includeBase64: true,
//         cropping: false
//     }).then(image => {
//         console.log('selectPhotoFromGallery', image);
//         onSelectedImage(image)
//         this.setState({ defaultImage: false })
//     });
// }
//
// export function onClickAddImage () {
//     const BUTTONS = ['Take Photo', 'Choose Photo from Library', 'Cancel'];
//     ActionSheet.show(
//         {
//             options: BUTTONS,
//             cancelButtonIndex: 2,
//             title: 'Select a Photo'
//         },
//         buttonIndex => {
//             switch (buttonIndex) {
//                 case 0:
//                     takePhoto();
//                     break;
//                 case 1:
//                     selectPhotoFromGallery();
//                     break;
//                 default:
//                     break;
//             }
//         }
//     )
// }
//
