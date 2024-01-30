import React, { useState, useEffect } from 'react';

const Container = () => {
    const [data, setData] = useState([]);

    const fetchData = () => {
        fetch('http://localhost:5000/data')
            .then(response => response.json())
            .then(data => setData(data))
            .catch(err => console.log(err));
    };


    return (
        <div className="container">
            <div className="card">
                <div className="card-details">
                    <p className="text-title">Tarea 1 - SO1 - 1s2024</p>
                    <p className="text-body">{data.name}</p>
                    <p className="text-body">{data.id}</p>
                    <p className="text-body">{data.date}</p>
                </div>
                <button className="card-button" onClick={fetchData}>
                    Mostrar Datos
                </button>
            </div>
        </div>
    );
};

export default Container;
