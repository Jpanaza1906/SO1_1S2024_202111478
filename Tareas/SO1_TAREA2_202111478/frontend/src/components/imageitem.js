import React from "react";

const ImageItem = ({image}) => {

    return (
        <div className="image-item">
            <div className="image-container">
                <img src={image.imgb64} alt="Captured" />
            </div>
        </div>
    )
}

export default ImageItem;