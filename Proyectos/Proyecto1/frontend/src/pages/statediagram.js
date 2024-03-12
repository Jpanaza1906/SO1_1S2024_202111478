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
import {ToastContainer, toast} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import State from '../components/state';





function StateDiagram() {
    const [pid, setPid] = useState(null);
    const [previusAction, setPreviusAction] = useState(null);
    const [currentAction, setCurrentAction] = useState(null);

    const handleClick = (action) => {
        //Si es la accion de start se debe hacer un fetch a la api para obtener el pid
        if (previusAction === action) {
            toast.error('You can not do the same action twice')
            return;
        }

        if (action === 'start') {

            //si el pid es null se debe hacer un fetch a la api para obtener el pid
            if (pid !== null) {
                toast.error('You can not start a new process while other process is running')
                return;
            }
            fetch('/api/statediagram?action=start')
                .then(response => response.json())
                .then(data => {
                    setPid(data.pid);
                    setCurrentAction(action);
                });
        } else{
            // se verifica que el pid no sea null
            if (pid !== null) {
                // si la action es un kill se debe limpiar el pid
                if (action === 'kill') {
                    setPid(null);
                    setPreviusAction(null);
                }

                // se manda por parametros la accion y el pid
                fetch(`/api/statediagram?action=${action}&pid=${pid}`)
                    .then(response => response.json())
                    .then(data => {
                        console.log(data);
                        setCurrentAction(action);
                    });
            }
        }

        // se guarda la accion previa
        setPreviusAction(action);


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
                <Col>
                    <CardG title={"State Diagram"}>
                        <State action={currentAction} />
                    </CardG>
                </Col>
            <ToastContainer />
            </Row>
        </div>
    );
}

export default StateDiagram;