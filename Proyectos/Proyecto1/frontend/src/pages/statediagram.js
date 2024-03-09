import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/statediagram.css';
import CardG from '../components/cardg';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { SiAddthis } from "react-icons/si";


import { FaStopCircle } from "react-icons/fa";
import { FaCirclePlay } from "react-icons/fa6";
import { FaSkullCrossbones } from "react-icons/fa6";






function StateDiagram() {
    const [pid, setPid] = useState(null);

    const handleClick = (action) => {
        //Si es la accion de start se debe hacer un fetch a la api para obtener el pid

        if (action === 'start') {
            fetch('http://localhost:8000/statediagram?action=start')
                .then(response => response.json())
                .then(data => {
                    setPid(data.pid);
                });
        } else{
            // se verifica que el pid no sea null
            if (pid !== null) {
                // se manda por parametros la accion y el pid
                fetch(`http://localhost:8000/statediagram?action=${action}&pid=${pid}`)
                    .then(response => response.json())
                    .then(data => {
                        console.log(data);
                    });
            }
        }


    }

    return (
        <div className='Contain'>
            <Row>
                <Col>
                    <button>
                        <a href="/statediagram" className="btn2"><span className="spn2">PID = {pid}</span></a>
                    </button>
                </Col>
                <Col>
                    <div className="button-container">
                        <button className="button" onClick={() => handleClick('start')}> {/* start button */}
                            <SiAddthis size="1.7em" color="white" />
                        </button>
                        <button className="button" onClick={() => handleClick('stop')}> {/* stop button */}
                            <FaStopCircle size="2em" color='white' />
                        </button>
                        <button className="button" onClick={() => handleClick('resume')}> {/* resume button */}
                            <FaCirclePlay size="2em" color='white' />
                        </button>
                        <button className="button" onClick={() => handleClick('kill')}> {/* kill button */}
                            <FaSkullCrossbones size="1.9em" color='white' />
                        </button>
                    </div>
                </Col>
            </Row>
            <Row>
                <CardG title={"State Diagram"}>
                </CardG>
            </Row>
        </div>
    );
}

export default StateDiagram;