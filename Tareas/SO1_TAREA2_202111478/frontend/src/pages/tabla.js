import React, { useEffect, useState } from 'react';
import Header from "../components/header";
import ImageItem from '../components/imageitem';
import Table from 'react-bootstrap/Table';
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css'
const TablaImages = () => {
    const [images, setImages] = useState([]);

    useEffect(() => {
        fetch('http://localhost:5000/imagenes')
            .then(response => response.json())
            .then(data => {
                setImages(data);
            })
            .catch(error => {
                console.error('Error fetching images', error);
            });
    }, []);


    return (
        <div className='main'>
            <Header />
            <div className="back">
                <Table striped bordered hover className="custom-table">
                    <thead>
                        <tr>
                            <th>Imagen</th>
                            <th>Fecha</th>
                        </tr>
                    </thead>
                    <tbody>
                        {images.map(image => (
                            <tr key={image._id}>
                                <td>
                                    <ImageItem image={image} />
                                </td>
                                <td>{image.fecha}</td>
                            </tr>
                        ))}
                    </tbody>
                </Table>
            </div>
        </div>
    );

}

export default TablaImages;