import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { MdMonitorHeart } from "react-icons/md";
import { useLocation } from 'react-router-dom';
import '../css/header.css';


function Header() {
    //obtener la ruta actual
    const location = useLocation();
    
    //Funcion para determinar si un enlace tiene clase btn o active
    const isActiveLink = (link) => {
        if (link === '/monitor' && location.pathname === '/') return 'btnactive';
        return location.pathname === link ? 'btnactive' : 'btn';
    }

    return (
        <>
            <Navbar className='navbar'>
                <MdMonitorHeart size="2em" color="white" />
                <Container fluid>
                    <Navbar.Brand href="monitor" className='brand'>SYSTEM MONITOR</Navbar.Brand>
                    <Nav className="me-0.25">
                        <Nav.Link href="monitor">
                            <button className={isActiveLink('/monitor')}>MONITOR</button>
                        </Nav.Link>
                        <Nav.Link href="processtree">
                            <button className={isActiveLink('/processtree')}>PROCESS TREE</button>
                        </Nav.Link>
                        <Nav.Link href="statediagram">
                            <button className={isActiveLink('/statediagram')}>STATE DIAGRAM</button>
                        </Nav.Link>
                    </Nav>
                </Container>
            </Navbar>
        </>
    );
}

export default Header;