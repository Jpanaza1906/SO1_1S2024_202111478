import React, { useRef, useState } from 'react';
import './camera.css';
import Webcam from 'react-webcam';


const CameraComponent = () => {
    const webcamRef = useRef(null);
    const [capturedImage, setCapturedImage] = useState(null);
    const [showCamera, setShowCamera] = useState(true);
  
    const captureImage = () => {
      const imageSrc = webcamRef.current.getScreenshot();
      setCapturedImage(imageSrc);
      setShowCamera(false);
    };
  
    const retakePhoto = () => {
      setCapturedImage(null);
      setShowCamera(true);
    };

    return (
        <div>
            {showCamera && (
                <div
                    style={{
                        position: 'absolute',
                        top: '50%',
                        left: '50%',
                        transform: 'translate(-50%, -50%)',
                    }}
                >
                    <Webcam
                        audio={false}
                        ref={webcamRef}
                        screenshotFormat="image/jpeg"
                        style={{
                            width: '100%',
                            height: 'auto',
                            position: 'relative',
                        }}
                    />
                </div>
            )}

            <div
                style={{
                    position: 'absolute',
                    textAlign: 'center',
                    top: '80%',
                    left: '32%',
                }}
            >
                <button class="button" onClick={captureImage}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24" height="24" fill="none" class="svg-icon"><g stroke-width="2" stroke-linecap="round" stroke="#fff" fill-rule="evenodd" clip-rule="evenodd"><path d="m4 9c0-1.10457.89543-2 2-2h2l.44721-.89443c.33879-.67757 1.03131-1.10557 1.78889-1.10557h3.5278c.7576 0 1.4501.428 1.7889 1.10557l.4472.89443h2c1.1046 0 2 .89543 2 2v8c0 1.1046-.8954 2-2 2h-12c-1.10457 0-2-.8954-2-2z"></path><path d="m15 13c0 1.6569-1.3431 3-3 3s-3-1.3431-3-3 1.3431-3 3-3 3 1.3431 3 3z"></path></g></svg>
                    <span class="lable">Take a Photo</span>
                </button>
            </div>

            <div
                style={{
                    position: 'absolute',
                    textAlign: 'center',
                    top: '80%',
                    left: '42%',
                }}
            >
                <button class="button" onClick={retakePhoto}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24" height="24" fill="none" class="svg-icon"><g stroke-width="2" stroke-linecap="round" stroke="#fff" fill-rule="evenodd" clip-rule="evenodd"><path d="m4 9c0-1.10457.89543-2 2-2h2l.44721-.89443c.33879-.67757 1.03131-1.10557 1.78889-1.10557h3.5278c.7576 0 1.4501.428 1.7889 1.10557l.4472.89443h2c1.1046 0 2 .89543 2 2v8c0 1.1046-.8954 2-2 2h-12c-1.10457 0-2-.8954-2-2z"></path><path d="m15 13c0 1.6569-1.3431 3-3 3s-3-1.3431-3-3 1.3431-3 3-3 3 1.3431 3 3z"></path></g></svg>
                    <span class="lable">Retake</span>
                </button>
            </div>

            <div
                style={{
                    position: 'absolute',
                    textAlign: 'center',
                    top: '80%',
                    left: '61%',
                }}
            >
                <button className='send'>
                    <div class="svg-wrapper-1">
                        <div class="svg-wrapper">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                width="24"
                                height="24"
                            >
                                <path fill="none" d="M0 0h24v24H0z"></path>
                                <path
                                    fill="currentColor"
                                    d="M1.946 9.315c-.522-.174-.527-.455.01-.634l19.087-6.362c.529-.176.832.12.684.638l-5.454 19.086c-.15.529-.455.547-.679.045L12 14l6-8-8 6-8.054-2.685z"
                                ></path>
                            </svg>
                        </div>
                    </div>
                    <span>Send</span>
                </button>

            </div>

            {capturedImage && (
                <div
                    style={{
                        position: 'absolute',
                        top: '50%',
                        left: '50%',
                        transform: 'translate(-50%, -50%)',
                    }}
                >
                    <img src={capturedImage} alt="Captured" />
                </div>
            )}
        </div>
    );
};

export default CameraComponent;
