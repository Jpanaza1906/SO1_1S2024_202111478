import React, { useEffect, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/processtree.css';
import CardG from '../components/cardg';
import ComboBox from '../components/combobox';
import Tree from '../components/tree'

function ProcessTree() {
    const [selectedOption, setSelectedOption] = useState(null);
    const [pidData, setPidData] = useState([]);
    const [treeData, setTreeData] = useState('');

    // Consultar el endpoint /processtree para obtener el array de PID disponibles
    useEffect(() => {
        fetchPidData();
        const interval = setInterval(fetchPidData, 1000);

        return () => clearInterval(interval);
    }, []);

    const fetchPidData = () => {
        fetch('http://localhost:8000/processtree')
            .then(response => response.json())
            .then(data => setPidData(data))
            .catch(error => console.error('Error fetching PID options:', error));
    }

    // Consultar el endpoint /processtree/{pid} para obtener el arbol de procesos
    useEffect(() => {
        if (selectedOption) {
            fetch(`http://localhost:8000/processtree/?pid=${selectedOption}`)
                .then(response => response.json())
                .then(data => setTreeData(data))
                .catch(error => console.error('Error fetching process tree:', error));
        }
    }, [selectedOption]);

    const handleChange = (event) => {
        setSelectedOption(event.target.value);
    }

    return (
        <div className='Contain'>
            <ComboBox options={pidData} onChange={handleChange} placeholder={"Seleecione el PID padre"}></ComboBox>
            <CardG title={"Tree View"}>
                {treeData && <Tree processData={treeData}/>} {}
            </CardG>
        </div>
    );
}

export default ProcessTree;
