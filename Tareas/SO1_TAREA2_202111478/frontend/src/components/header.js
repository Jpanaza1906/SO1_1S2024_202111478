import {Container, Nav, Navbar } from 'react-bootstrap';
import { MdCamera } from "react-icons/md";

const Header = () => {
    return (
          <Navbar bg="dark" data-bs-theme="dark">
            <MdCamera size="2em" color="white"/>
            <Container fluid>
              <Navbar.Brand href="index">CAPTURE</Navbar.Brand>
              <Nav className="me-0.25">
                <Nav.Link href="index">HOME</Nav.Link>
                <Nav.Link href="images">IMAGES</Nav.Link>
              </Nav>
            </Container>
          </Navbar>
      );
}

export default Header;