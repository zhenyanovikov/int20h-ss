import { Box, Container, Stack, Button } from "@mui/material";
import { ArrowBack as ArrowBackIcon } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import SendEmailForm from "../../organisms/SendEmailForm/SendEmailForm";
import { ROUTE } from "../../../constants/router";

function SendEmailTemplate() {
  const { t } = useTranslation();
  const navigate = useNavigate();

  return (
    <Box py={8}>
      <Container>
        <Stack spacing={4}>
          <Stack alignSelf="flex-start">
            <Button onClick={handleGoBack} startIcon={<ArrowBackIcon />}>
              {t("templates.sendEmail.button.back")}
            </Button>
          </Stack>

          <SendEmailForm />
        </Stack>
      </Container>
    </Box>
  );

  function handleGoBack() {
    navigate(ROUTE.EMAIL_TEMPLATES);
  }
}

export default SendEmailTemplate;
