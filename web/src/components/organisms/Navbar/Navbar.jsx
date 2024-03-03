import { AppBar, Container, Toolbar, Stack, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";
import LoremLogo from "../../../assets/logos/lorem.svg";
import useIsAuthenticated from "../../../hooks/useIsAuthenticated";
import MailingButton from "../../molecules/MailingButton/MailingButton";
import LocaleToggleButton from "../../molecules/LocaleToggleButton/LocaleToggleButton";
import ThemeToggleButton from "../../molecules/ThemeToggleButton/ThemeToggleButton";
import UserMenu from "../../molecules/UserMenu/UserMenu";

function Navbar() {
  const navigate = useNavigate();
  const isAuthicated = useIsAuthenticated();

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
              onClick={handleLogoClick}
            />
            <Stack direction="row" alignItems="center" spacing={1}>
              {isAuthicated && <MailingButton />}
              <LocaleToggleButton />
              <ThemeToggleButton />
              {isAuthicated && <UserMenu />}
            </Stack>
          </Stack>
        </Toolbar>
      </Container>
    </AppBar>
  );

  function handleLogoClick() {
    navigate("/");
  }
}

export default Navbar;
