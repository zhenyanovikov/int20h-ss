import { Stack, Box, Container, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";
import { Navigate } from "react-router-dom";
import { useGetUser } from "../../../api/user";
import { ROUTE } from "../../../constants/router";
import LoremBackground from "../../../assets/backgrounds/lorem.png";
import useIsAuthenticated from "../../../hooks/useIsAuthenticated";
import GoogleLoginButton from "../../molecules/GoogleLoginButton/GoogleLoginButton";
import { ROLE } from "../../../constants/user";

function EntryTemplate() {
  const { t } = useTranslation();

  const isAuthicated = useIsAuthenticated();
  const { data: userData } = useGetUser({ enabled: isAuthicated });

  if (userData) {
    if (userData.role === ROLE.ADMIN) {
      return <Navigate to={ROUTE.FACULTIES} />;
    }

    return <Navigate to={ROUTE.FACULTIES} />;
  }

  return (
    <>
      <Box py={16}>
        <Container>
          <Stack
            spacing={2}
            sx={{
              width: {
                xs: "100%",
                md: "80%",
              },
            }}
          >
            <Typography variant="h3" component="h1">
              {t("templates.entry.title")}
            </Typography>
            <Typography variant="subtitle1">
              {t("templates.entry.subtitle")}
            </Typography>
            <Box>
              <GoogleLoginButton variant="contained" />
            </Box>
          </Stack>
        </Container>
      </Box>
      <Box
        sx={{
          backgroundImage: `url(${LoremBackground})`,
          backgroundRepeat: "no-repeat",
          backgroundSize: "cover",
          bottom: {
            xs: 32,
            md: "unset",
          },
          height: {
            xs: 200,
            md: 400,
          },
          position: "absolute",
          right: 0,
          top: {
            xs: "unset",
            md: 88,
          },
          width: {
            xs: 150,
            md: 300,
          },
          zIndex: -1,
        }}
      />
    </>
  );
}

export default EntryTemplate;
