import React, { useEffect, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from '../components/card';
import PieChart from '../components/piechart';
import LineChart from '../components/linechart';
import '../css/monitor.css';


function Monitor() {

    //Constantes para las graficas
    const [data_ram, setDataRam] = useState([50,50]);
    const [data_cpu, setDataCpu] = useState([50,50]);
    const [data_history, setDataHistory] = useState('');


    useEffect(() => {
        const fetchData = async () => {
            // Hacer una solicitud a tu backend en el puerto 8000, endpoint /monitor
            try {
                const response = await fetch('http://localhost:8000/monitor');

                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }

                const resp = await response.json();
                console.log(resp);
                setDataRam(resp.data.ram_percentage);
                setDataCpu(resp.data.cpu_percentage);
                setDataHistory(resp.data_historial);

            } catch (error) {
                console.error('Error:', error);
            }
        };

        // Realizar la primera solicitud al montar el componente
        fetchData();

        // Establecer un intervalo para realizar la solicitud cada 500ms
        const intervalId = setInterval(fetchData, 500);

        // Limpiar el intervalo al desmontar el componente
        return () => clearInterval(intervalId);
    }, []);

    //Datos de la grafica de personas por vehiculo
    //const data_personas = [dataGraphs.personal, dataGraphs.mediano, dataGraphs.grande];

    //const data_ram = [50, 50];

    //const data_cpu = [50, 50];

    //const data = [12, 19, 3, 5, 2, 3];
    //const labels = ['January', 'February', 'March', 'April', 'May', 'June'];

    
    // const data = {
    //     labels: ['1', '2', '3', '4', '5'],
    //     datasets: [
    //         {
    //             label: 'RAM',
    //             data: [54, 44, 31, 29, 13],
    //             borderColor: '#94d2bd',
    //             backgroundColor: '#94d2bd',
    //             borderWidth: 1,
    //             tension: 0.5, // Suaviza la línea
    //             fill: false
    //         },
    //         {
    //             label: 'CPU',
    //             data: [10, 22, 35, 41, 59],
    //             borderColor: '#ee9b00',
    //             backgroundColor: '#ee9b00',
    //             borderWidth: 1,
    //             tension: 0.5, // Suaviza la línea
    //             fill: false
    //         }
    //     ]
    // };

    //hacer una consulta a mi backend en el puerto 8000, endpoint /monitor



    return (
        <div className='Contain'>
            <Container>
                <Row>
                    <Col>
                        <Card title={"Monitoreo en Tiempo Real"}>
                            <Row>
                                <Col>                                    
                                    <PieChart data={data_ram} labels={[`Libre: ${data_ram[0]}%`, `En uso: ${data_ram[1]}%`]} colors={['#94d2bd', '#005f73']} title={"RAM"} />
                                </Col>
                                <Col>
                                    <PieChart data={data_cpu} labels={[`Libre: ${data_cpu[0]}%`, `En uso: ${data_cpu[1]}%`]} colors={['#ca6702', '#bb3e03']} title={"CPU"} />
                                </Col>
                            </Row>
                        </Card>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <Card title={"Monitoreo Histórico"}>
                            <Row>
                                <Col>
                                    <LineChart data={data_history} />
                                </Col>
                            </Row>            </Card>
                    </Col>
                </Row>
            </Container>
        </div>
    );
}

export default Monitor;
