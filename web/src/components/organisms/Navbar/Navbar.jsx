import { AppBar, Container, Toolbar, Stack, Box } from "@mui/material";
import LoremLogo from "../../../assets/logos/lorem.svg";

function Navbar() {
  return (
    <AppBar position="static">
      <Container>
        <Toolbar disableGutters>
          <Stack width="100%" direction="row" justifyContent="space-between">
            <Box
              component="img"
              alt=""
              src={LoremLogo}
              sx={{
                cursor: "pointer",
              }}
              onClick={() => {}}
            />
            <Stack direction="row" alignItems="center" spacing={1} />
          </Stack>
        </Toolbar>
      </Container>
    </AppBar>
  );
}

export default Navbar;
