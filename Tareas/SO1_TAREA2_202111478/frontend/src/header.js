import {Container, Nav, Navbar } from 'react-bootstrap';
import { MdCamera } from "react-icons/md";

const Header = () => {
    return (
          <Navbar bg="dark" data-bs-theme="dark">
            <MdCamera size="2em" color="white"/>
            <Container fluid>
              <Navbar.Brand href="#home">CAPTURE</Navbar.Brand>
              <Nav className="me-0.25">
                <Nav.Link href="#home">HOME</Nav.Link>
                <Nav.Link href="#features">IMAGES</Nav.Link>
              </Nav>
            </Container>
          </Navbar>
      );
}

export default Header;