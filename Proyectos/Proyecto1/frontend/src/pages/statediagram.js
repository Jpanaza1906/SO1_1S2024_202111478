import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/statediagram.css';
import CardG from '../components/cardg';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { MdAddCircleOutline } from "react-icons/md";
import { FaRegStopCircle } from "react-icons/fa";
import { VscRunAll } from "react-icons/vsc";
import { IoSkullOutline } from "react-icons/io5";





function StateDiagram() {
    return (
        <div className='Contain'>
            <Row>
                <Col>
                    <button>
                        <a href="/statediagram" class="btn2"><span class="spn2">PID = 12343</span></a>
                    </button>
                </Col>
                <Col>
                    <div className="button-container">
                        <button className="button">
                            <MdAddCircleOutline size="1.6em" color="white" />
                        </button>
                        <button className="button">
                            <FaRegStopCircle size="1.4em" color='white' />
                        </button>
                        <button className="button">
                            <VscRunAll size="1.5em" color='white' />
                        </button>
                        <button className="button">
                            <IoSkullOutline size="1.5em" color='white' />
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