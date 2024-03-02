import { AppBar, Container, Toolbar } from "@mui/material";

function Navbar() {
  return (
    <AppBar position="static">
      <Container>
        <Toolbar disableGutters />
      </Container>
    </AppBar>
  );
}

export default Navbar;
