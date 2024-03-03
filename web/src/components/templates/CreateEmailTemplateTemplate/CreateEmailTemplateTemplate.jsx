import { Box, Container, Stack, Button } from "@mui/material";
import { ArrowBack as ArrowBackIcon } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import CreateEmailTemplateForm from "../../organisms/CreateEmailTemplateForm/CreateEmailTemplateForm";
import { ROUTE } from "../../../constants/router";

function CreateEmailTemplateTemplate() {
  const { t } = useTranslation();
  const navigate = useNavigate();

  return (
    <Box py={8}>
      <Container>
        <Stack spacing={4}>
          <Stack alignSelf="flex-start">
            <Button onClick={handleGoBack} startIcon={<ArrowBackIcon />}>
              {t("templates.createEmailTemplate.button.back")}
            </Button>
          </Stack>

          <CreateEmailTemplateForm />
        </Stack>
      </Container>
    </Box>
  );

  function handleGoBack() {
    navigate(ROUTE.EMAIL_TEMPLATES);
  }
}

export default CreateEmailTemplateTemplate;
