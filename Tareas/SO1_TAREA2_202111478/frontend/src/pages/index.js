import Header from "../components/header";
import CameraComponent from "../components/camera";
import { ToastContainer } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';
import './index.css'
const Index = () => {
    return (
        <div className="main">
            <Header />
            <div className="back">
                <CameraComponent />
                <ToastContainer />
            </div>
        </div>
    );
}

export default Index