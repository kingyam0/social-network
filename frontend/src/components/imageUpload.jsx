//handling image upload for posts
const ImageUpload = (event) => {
  const file = event.target.files[0];
  setPostPhotoUp(file);

  // Read the image file as a data URL
  const reader = new FileReader();
  reader.onloadend = () => {
    setImageDataUrl(reader.result);
  };

  reader.readAsDataURL(file);

return (
  <>
    <label id="PostPhotoUp" htmlFor="PostPhotoUp">
      {" "}
      Image:{" "}
    </label>
    <input
      id="PhotoUp"
      type="file"
      accept="image/*"
      onChange={ImageUpload}
    />
    {PostPhotoUp && <img id="photoShow" src={ImageDataUrl} alt="uploaded" />}
  </>
);

};

export default ImageUpload